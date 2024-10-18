package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

// NewEventStore crée une nouvelle instance de EventStore
func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{db: db}
}

// Créer un événement
func (s *EventStore) Create(event *models.Event) error {
	return s.db.Create(event).Error
}

// Mettre à jour un événement existant
func (s *EventStore) Update(event *models.Event) error {
	return s.db.Save(event).Error
}

// Supprimer un événement
func (s *EventStore) Delete(id uint) error {
	return s.db.Delete(&models.Event{}, id).Error
}

// Récupérer un événement par son ID
func (s *EventStore) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := s.db.First(&event, id).Error
	return &event, err
}
