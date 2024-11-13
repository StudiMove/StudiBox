package event

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type EventDescriptionStoreType struct {
	db *gorm.DB
}

// NewEventDescriptionStore crée une nouvelle instance de EventDescriptionStoreType
func EventDescriptionStore(db *gorm.DB) *EventDescriptionStoreType {
	return &EventDescriptionStoreType{db: db}
}

// Create ajoute une nouvelle description pour un événement
func (s *EventDescriptionStoreType) Create(eventDescription *models.EventDescription) error {
	if err := s.db.Create(eventDescription).Error; err != nil {
		return errors.New("erreur lors de la création de la description de l'événement : " + err.Error())
	}
	return nil
}

// Update met à jour une description existante d'un événement
func (s *EventDescriptionStoreType) Update(eventDescription *models.EventDescription) error {
	if err := s.db.Save(eventDescription).Error; err != nil {
		return errors.New("erreur lors de la mise à jour de la description de l'événement : " + err.Error())
	}
	return nil
}

// Delete supprime une description par ID
func (s *EventDescriptionStoreType) Delete(id uint) error {
	if err := s.db.Delete(&models.EventDescription{}, id).Error; err != nil {
		return errors.New("erreur lors de la suppression de la description de l'événement : " + err.Error())
	}
	return nil
}

func (s *EventDescriptionStoreType) DeleteByEventIDWithTx(eventID uint, tx *gorm.DB) error {
	if err := tx.Where("event_id = ?", eventID).Delete(&models.EventDescription{}).Error; err != nil {
		return errors.New("Erreur lors de la suppression des descriptions : " + err.Error())
	}
	return nil
}

// GetByEventID récupère les descriptions associées à un événement spécifique
func (s *EventDescriptionStoreType) GetByEventID(eventID uint) ([]models.EventDescription, error) {
	var descriptions []models.EventDescription
	if err := s.db.Where("event_id = ?", eventID).Find(&descriptions).Error; err != nil {
		return nil, errors.New("erreur lors de la récupération des descriptions de l'événement : " + err.Error())
	}
	return descriptions, nil
}
