package event

import (
	"backend/core/models"
	stores "backend/core/stores/event"
	"errors"
	"sort"

	"gorm.io/gorm"
)

type EventInteractionService struct {
	eventStore     *stores.EventStore
	eventLikeStore *stores.EventLikeStore
	eventViewStore *stores.EventViewStore
}

func NewEventInteractionService(eventStore *stores.EventStore, eventLikeStore *stores.EventLikeStore, eventViewStore *stores.EventViewStore) *EventInteractionService {
	return &EventInteractionService{
		eventStore:     eventStore,
		eventLikeStore: eventLikeStore,
		eventViewStore: eventViewStore,
	}
}

func (s *EventInteractionService) LogEventView(userID, eventID uint) error {
	return s.eventViewStore.AddEventView(userID, eventID)
}

func (s *EventInteractionService) GetLikesCount(eventID uint) (int, error) {
	return s.eventLikeStore.CountLikes(eventID)
}

func (s *EventInteractionService) GetViewsCount(eventID uint) (int, error) {
	return s.eventViewStore.CountViews(eventID)
}

func (s *EventInteractionService) LikeEvent(userID, eventID uint) error {
	event, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("l'événement spécifié n'existe pas")
		}
		return err
	}

	if !event.IsVisible {
		return errors.New("cet événement n'est plus actif")
	}

	return s.eventLikeStore.LikeEvent(userID, eventID)
}

func (s *EventInteractionService) UnlikeEvent(userID, eventID uint) error {
	_, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("l'événement spécifié n'existe pas")
		}
		return err
	}

	return s.eventLikeStore.UnlikeEvent(userID, eventID)
}

func (s *EventInteractionService) GetRecommendations(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	if err := s.eventLikeStore.GetLikedEventsByUserWithRelations(userID, &likedEvents); err != nil {
		return nil, errors.New("impossible de récupérer les événements likés")
	}
	viewedEvents, err := s.eventViewStore.GetMostViewedEventsByUser(userID, 5)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements vus")
	}

	tags := s.extractTagsFromEvents(likedEvents, viewedEvents)
	categories := s.extractCategoriesFromEvents(likedEvents, viewedEvents)

	recommendedEvents, err := s.eventStore.GetRecommendedEventsByTags(tags, categories, nil, 10)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements recommandés")
	}

	eventScores := s.calculateScores(likedEvents, viewedEvents, recommendedEvents)

	sortedEvents := s.sortEventsByScore(eventScores)

	if len(sortedEvents) > 5 {
		sortedEvents = sortedEvents[:5]
	}

	return sortedEvents, nil
}

func (s *EventInteractionService) calculateScores(likedEvents, viewedEvents, recommendedEvents []models.Event) map[uint]int {
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

func (s *EventInteractionService) sortEventsByScore(eventScores map[uint]int) []models.Event {
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

func (s *EventInteractionService) extractTagsFromEvents(likedEvents, viewedEvents []models.Event) []string {
	tagSet := make(map[string]struct{})
	for _, event := range append(likedEvents, viewedEvents...) {
		for _, tag := range event.Tags {
			tagSet[tag.Name] = struct{}{}
		}
	}
	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	return tags
}

func (s *EventInteractionService) extractCategoriesFromEvents(likedEvents, viewedEvents []models.Event) []string {
	categorySet := make(map[string]struct{})
	for _, event := range append(likedEvents, viewedEvents...) {
		for _, category := range event.Categories {
			categorySet[category.Name] = struct{}{}
		}
	}
	categories := make([]string, 0, len(categorySet))
	for category := range categorySet {
		categories = append(categories, category)
	}
	return categories
}
