package event

import (
	"backend/core/models"
	stores "backend/core/stores/event"
	"encoding/json"
	"errors"
)

type EventRetrievalService struct {
	eventStore     *stores.EventStore
	eventLikeStore *stores.EventLikeStore
}

func NewEventRetrievalService(eventStore *stores.EventStore, eventLikeStore *stores.EventLikeStore) *EventRetrievalService {
	return &EventRetrievalService{
		eventStore:     eventStore,
		eventLikeStore: eventLikeStore,
	}
}

func (s *EventRetrievalService) GetEvent(id uint) (*models.Event, error) {
	event, err := s.eventStore.GetByID(id)
	if err != nil {
		return nil, errors.New("erreur lors de la récupération de l'événement: " + err.Error())
	}

	if err := json.Unmarshal([]byte(event.ImageURLsJSON), &event.ImageURLs); err != nil {
		event.ImageURLs = []string{}
	}

	return event, nil
}

func (s *EventRetrievalService) ListEvents(page, limit int, category, city string) ([]models.Event, int64, error) {
	events, total, err := s.eventStore.List(page, limit, category, city)
	if err != nil {
		return nil, 0, err
	}
	return events, total, nil
}

func (s *EventRetrievalService) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.eventLikeStore.GetLikedEventsByUserWithRelations(userID, &likedEvents)
	if err != nil {
		return nil, err
	}
	return likedEvents, nil
}
