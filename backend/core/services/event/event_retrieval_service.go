// event_retrieval.go
package event

import (
	"backend/core/models"
	"encoding/json"
	"errors"
)

// Récupère un événement par ID
func (s *EventService) GetEvent(id uint) (*models.Event, error) {
	event, err := s.eventStore.GetByID(id)
	if err != nil {
		return nil, errors.New("erreur lors de la récupération de l'événement: " + err.Error())
	}

	// Convertir les URLs d'images depuis JSON
	if err := json.Unmarshal([]byte(event.ImageURLsJSON), &event.ImageURLs); err != nil {
		event.ImageURLs = []string{}
	}

	return event, nil
}

// Liste paginée des événements avec des filtres
func (s *EventService) ListEvents(page, limit int, category, city string) ([]models.Event, int64, error) {
	events, total, err := s.eventStore.List(page, limit, category, city)
	if err != nil {
		return nil, 0, err
	}
	return events, total, nil
}

// Récupère les événements likés par un utilisateur
func (s *EventService) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.eventLikeStore.GetLikedEventsByUserWithRelations(userID, &likedEvents)
	if err != nil {
		return nil, err
	}
	return likedEvents, nil
}
