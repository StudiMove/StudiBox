package event

import (
	"backend/core/models"
	"backend/core/stores"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type EventService struct {
	eventStore *stores.EventStore
}

// NewEventService crée une nouvelle instance de EventService avec une base de données
func NewEventService(db *gorm.DB) *EventService {
	eventStore := stores.NewEventStore(db)
	return &EventService{eventStore: eventStore}
}

// Crée un événement
func (s *EventService) CreateEvent(event *models.Event) error {
	// Convertir les URLs d'image en JSON pour stockage
	imageURLsJSON, err := json.Marshal(event.ImageURLs)
	if err != nil {
		return errors.New("Erreur de conversion des URLs d'images en JSON: " + err.Error())
	}
	event.ImageURLsJSON = string(imageURLsJSON)

	if err := s.eventStore.Create(event); err != nil {
		return errors.New("Erreur lors de la création de l'événement: " + err.Error())
	}

	return nil
}

// Met à jour un événement existant
func (s *EventService) UpdateEvent(event *models.Event) error {
	imageURLsJSON, err := json.Marshal(event.ImageURLs)
	if err != nil {
		return errors.New("Erreur de conversion des URLs d'images en JSON: " + err.Error())
	}
	event.ImageURLsJSON = string(imageURLsJSON)

	if err := s.eventStore.Update(event); err != nil {
		return errors.New("Erreur lors de la mise à jour de l'événement: " + err.Error())
	}

	return nil
}

// Supprime un événement par ID
func (s *EventService) DeleteEvent(id uint) error {
	if err := s.eventStore.Delete(id); err != nil {
		return errors.New("Erreur lors de la suppression de l'événement: " + err.Error())
	}
	return nil
}

// Récupère un événement par ID
func (s *EventService) GetEvent(id uint) (*models.Event, error) {
	event, err := s.eventStore.GetByID(id)
	if err != nil {
		return nil, errors.New("Erreur lors de la récupération de l'événement: " + err.Error())
	}

	// Convertir le JSON des URLs d'images en tableau
	if err := json.Unmarshal([]byte(event.ImageURLsJSON), &event.ImageURLs); err != nil {
		event.ImageURLs = []string{}
	}

	return event, nil
}
