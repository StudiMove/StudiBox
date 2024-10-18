package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EventOptionStore struct {
	db *gorm.DB
}

// NewEventOptionStore crée une nouvelle instance de EventOptionStore
func NewEventOptionStore(db *gorm.DB) *EventOptionStore {
	return &EventOptionStore{db: db}
}

// Créer une option d'événement
func (s *EventOptionStore) Create(eventOption *models.EventOption) error {
	return s.db.Create(eventOption).Error
}

// Mettre à jour une option d'événement existante
func (s *EventOptionStore) Update(eventOption *models.EventOption) error {
	return s.db.Save(eventOption).Error
}

// Supprimer une option d'événement
func (s *EventOptionStore) Delete(id uint) error {
	return s.db.Delete(&models.EventOption{}, id).Error
}

// Récupérer une option d'événement par son ID
func (s *EventOptionStore) GetByID(id uint) (*models.EventOption, error) {
	var eventOption models.EventOption
	err := s.db.First(&eventOption, id).Error
	return &eventOption, err
}

// Récupérer toutes les options pour un événement spécifique
func (s *EventOptionStore) GetByEventID(eventID uint) ([]models.EventOption, error) {
	var options []models.EventOption
	err := s.db.Where("event_id = ?", eventID).Find(&options).Error
	return options, err
}
