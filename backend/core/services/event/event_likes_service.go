// event_likes.go
package event

import (
	"backend/core/models"
	"errors"

	"gorm.io/gorm"
)

// LogEventView permet de loguer une vue pour un événement
func (s *EventService) LogEventView(userID, eventID uint) error {
	return s.eventViewStore.AddEventView(userID, eventID)
}

// LikeEvent permet à un utilisateur de liker un événement
func (s *EventService) LikeEvent(userID, eventID uint) error {
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

// UnlikeEvent permet à un utilisateur de retirer son like d'un événement
func (s *EventService) UnlikeEvent(userID, eventID uint) error {
	_, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("l'événement spécifié n'existe pas")
		}
		return err
	}

	return s.eventLikeStore.UnlikeEvent(userID, eventID)
}

// GetRecommendations récupère des événements recommandés pour un utilisateur
func (s *EventService) GetRecommendations(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.eventLikeStore.GetLikedEventsByUserWithRelations(userID, &likedEvents)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements likés")
	}

	viewedEvents, err := s.eventViewStore.GetMostViewedEventsByUser(userID, 5)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements vus")
	}

	tags := s.extractTagsFromEvents(likedEvents, viewedEvents)
	categories := s.extractCategoriesFromEvents(likedEvents, viewedEvents)

	recommendedEvents, err := s.eventStore.GetRecommendedEventsByTags(tags, categories, nil, 5)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements recommandés")
	}

	popularEvents, err := s.eventStore.GetPopularEvents(5)
	if err != nil {
		return nil, errors.New("impossible de récupérer les événements populaires")
	}

	allEvents := append(likedEvents, viewedEvents...)
	allEvents = append(allEvents, recommendedEvents...)
	allEvents = append(allEvents, popularEvents...)

	return s.removeDuplicateEvents(allEvents), nil
}

// Fonction pour extraire les tags des événements likés et vus
func (s *EventService) extractTagsFromEvents(likedEvents, viewedEvents []models.Event) []string {
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

// Fonction pour extraire les catégories des événements likés et vus
func (s *EventService) extractCategoriesFromEvents(likedEvents, viewedEvents []models.Event) []string {
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

// RemoveDuplicateEvents supprime les événements dupliqués
func (s *EventService) removeDuplicateEvents(events []models.Event) []models.Event {
	eventMap := make(map[uint]bool)
	uniqueEvents := []models.Event{}

	for _, event := range events {
		if !eventMap[event.ID] {
			eventMap[event.ID] = true
			uniqueEvents = append(uniqueEvents, event)
		}
	}
	return uniqueEvents
}
