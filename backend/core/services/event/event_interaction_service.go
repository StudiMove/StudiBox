package event

import (
	stores "backend/core/stores/event"
	"backend/database/models"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

type EventInteractionServiceType struct {
	eventStore         *stores.EventStoreType
	eventLikeStore     *stores.EventLikeStoreType
	eventViewStore     *stores.EventViewStoreType
	tarifStore         *stores.EventTarifStoreType
	optionStore        *stores.EventOptionStoreType
	descriptionStore   *stores.EventDescriptionStoreType
	eventTagStore      *stores.EventTagStoreType
	eventCategoryStore *stores.EventCategoryStoreType
}

func EventInteractionService(eventStore *stores.EventStoreType, eventLikeStore *stores.EventLikeStoreType, eventViewStore *stores.EventViewStoreType, tarifStore *stores.EventTarifStoreType, optionStore *stores.EventOptionStoreType, descriptionStore *stores.EventDescriptionStoreType, eventTagStore *stores.EventTagStoreType, eventCategoryStore *stores.EventCategoryStoreType) *EventInteractionServiceType {
	return &EventInteractionServiceType{
		eventStore:         eventStore,
		eventLikeStore:     eventLikeStore,
		eventViewStore:     eventViewStore,
		tarifStore:         tarifStore,
		optionStore:        optionStore,
		descriptionStore:   descriptionStore,
		eventTagStore:      eventTagStore,
		eventCategoryStore: eventCategoryStore,
	}
}

func (s *EventInteractionServiceType) LogEventView(userID, eventID uint) error {
	// Recherche si une vue existe déjà pour cet utilisateur et cet événement
	eventView, err := s.eventViewStore.FindEventView(userID, eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Si la vue n'existe pas, on en crée une nouvelle
			return s.eventViewStore.CreateEventView(userID, eventID)
		}
		// Retourne l'erreur si elle est différente de "record not found"
		return err
	}

	// Incrémente le compteur de vues si elle existe
	return s.eventViewStore.IncrementViewCount(eventView)
}

func (s *EventInteractionServiceType) GetLikesCount(eventID uint) (int, error) {
	count, err := s.eventLikeStore.CountLikes(eventID)
	if err != nil {
		return 0, errors.New("erreur lors de la récupération du nombre de likes")
	}
	return count, nil
}

func (s *EventInteractionServiceType) GetViewsCount(eventID uint) (int, error) {
	count, err := s.eventViewStore.CountViews(eventID)
	if err != nil {
		return 0, errors.New("erreur lors de la récupération du nombre de vues")
	}
	return count, nil
}

func (s *EventInteractionServiceType) LikeEvent(userID, eventID uint) error {
	event, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("l'événement spécifié n'existe pas")
		}
		return err
	}

	if !event.IsPublic {
		return errors.New("cet événement n'est plus actif")
	}

	// Vérifie si le like existe déjà en utilisant le store
	exists, err := s.eventLikeStore.IsLikeExists(userID, eventID)
	if err != nil {
		return err
	}
	if exists {
		return nil // Le like existe déjà, donc on n'ajoute pas
	}

	// Ajoute le like en utilisant le store
	return s.eventLikeStore.AddLike(userID, eventID)
}

func (s *EventInteractionServiceType) UnlikeEvent(userID, eventID uint) error {
	_, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("l'événement spécifié n'existe pas")
		}
		return err
	}

	return s.eventLikeStore.UnlikeEvent(userID, eventID)
}

func (s *EventInteractionServiceType) GetRecommendations(userID uint) ([]models.Event, error) {
	// Récupérer les événements likés par l'utilisateur
	likedEvents, err := s.eventLikeStore.GetLikedEventsByUser(userID)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements likés")
	}

	// Récupérer les événements les plus vus par l'utilisateur
	viewedEvents, err := s.eventViewStore.GetMostViewedEventsByUser(userID, 5)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements vus")
	}

	// Extraire les tags et catégories des événements likés et vus
	tags, err := s.extractTagsFromEvents(append(likedEvents, viewedEvents...))
	if err != nil {
		return nil, errors.New("impossible d'extraire les tags")
	}

	categories, err := s.extractCategoriesFromEvents(append(likedEvents, viewedEvents...))
	if err != nil {
		return nil, errors.New("impossible d'extraire les catégories")
	}

	// Obtenir les événements recommandés basés sur les tags et catégories
	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tag.Name
	}

	categoryNames := make([]string, len(categories))
	for i, category := range categories {
		categoryNames[i] = category.Name
	}

	recommendedEvents, err := s.eventStore.GetRecommendedEventsByTags(tagNames, categoryNames, nil, 10)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements recommandés")
	}

	// Calculer les scores des événements recommandés
	eventScores := s.calculateScores(likedEvents, viewedEvents, recommendedEvents)

	// Trier les événements par score
	sortedEvents := s.sortEventsByScore(eventScores)

	// Limiter les résultats à 5 événements maximum
	if len(sortedEvents) > 5 {
		sortedEvents = sortedEvents[:5]
	}

	// Enrichir chaque événement avec des descriptions, options et tarifs
	for i := range sortedEvents {
		if err := s.populateEventDetails(&sortedEvents[i]); err != nil {
			return nil, fmt.Errorf("erreur lors de la récupération des détails pour l'événement ID %d : %w", sortedEvents[i].ID, err)
		}
	}

	return sortedEvents, nil
}

func (s *EventInteractionServiceType) calculateScores(likedEvents, viewedEvents, recommendedEvents []models.Event) map[uint]int {
	eventScores := make(map[uint]int)

	for _, event := range likedEvents {
		eventScores[event.ID] += 3
	}

	for _, event := range viewedEvents {
		eventScores[event.ID] += 1
	}

	for _, event := range recommendedEvents {
		eventScores[event.ID] += 2
	}

	return eventScores
}

func (s *EventInteractionServiceType) sortEventsByScore(eventScores map[uint]int) []models.Event {
	var sortedEvents []models.Event
	for eventID := range eventScores {
		if event, err := s.eventStore.GetByID(eventID); err == nil {
			sortedEvents = append(sortedEvents, *event)
		}
	}

	sort.SliceStable(sortedEvents, func(i, j int) bool {
		return eventScores[sortedEvents[i].ID] > eventScores[sortedEvents[j].ID]
	})

	return sortedEvents
}

func (s *EventInteractionServiceType) extractTagsFromEvents(events []models.Event) ([]models.EventTag, error) {
	tagIDsMap := make(map[int64]struct{})
	for _, event := range events {
		for _, tagID := range event.TagIDs {
			tagIDsMap[tagID] = struct{}{}
		}
	}

	tagIDs := make([]int64, 0, len(tagIDsMap))
	for tagID := range tagIDsMap {
		tagIDs = append(tagIDs, tagID)
	}

	if len(tagIDs) == 0 {
		return []models.EventTag{}, nil
	}
	return s.eventTagStore.GetTagsByIDs(tagIDs)
}

func (s *EventInteractionServiceType) extractCategoriesFromEvents(events []models.Event) ([]models.EventCategory, error) {
	categoryIDsMap := make(map[int64]struct{})
	for _, event := range events {
		for _, categoryID := range event.CategoryIDs {
			categoryIDsMap[categoryID] = struct{}{}
		}
	}

	categoryIDs := make([]int64, 0, len(categoryIDsMap))
	for categoryID := range categoryIDsMap {
		categoryIDs = append(categoryIDs, categoryID)
	}

	if len(categoryIDs) == 0 {
		return []models.EventCategory{}, nil
	}
	return s.eventCategoryStore.GetCategoriesByIDs(categoryIDs)
}

func (s *EventInteractionServiceType) populateEventDetails(event *models.Event) error {
	var err error

	// Récupérer les descriptions
	if event.Descriptions, err = s.descriptionStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des descriptions : %w", err)
	}

	// Récupérer les options
	if event.Options, err = s.optionStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des options : %w", err)
	}

	// Récupérer les tarifs
	if event.Tarifs, err = s.tarifStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des tarifs : %w", err)
	}

	return nil
}
