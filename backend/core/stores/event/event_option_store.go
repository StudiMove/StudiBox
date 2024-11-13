package event

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type EventOptionStoreType struct {
	db *gorm.DB
}

// NewEventOptionStore instancie un store pour les options d'événements
func EventOptionStore(db *gorm.DB) *EventOptionStoreType {
	return &EventOptionStoreType{db: db}
}

func (s *EventOptionStoreType) Create(eventOption *models.EventOption) error {
	if err := s.db.Create(eventOption).Error; err != nil {
		return errors.New("erreur lors de la création de l'option : " + err.Error())
	}
	return nil
}

func (s *EventOptionStoreType) Update(eventOption *models.EventOption) error {
	if err := s.db.Save(eventOption).Error; err != nil {
		return errors.New("erreur lors de la mise à jour de l'option : " + err.Error())
	}
	return nil
}

func (s *EventOptionStoreType) Delete(id uint) error {
	if err := s.db.Delete(&models.EventOption{}, id).Error; err != nil {
		return errors.New("erreur lors de la suppression de l'option : " + err.Error())
	}
	return nil
}

func (s *EventOptionStoreType) DeleteByEventIDWithTx(eventID uint, tx *gorm.DB) error {
	if err := tx.Where("event_id = ?", eventID).Delete(&models.EventOption{}).Error; err != nil {
		return errors.New("Erreur lors de la suppression des options : " + err.Error())
	}
	return nil
}

func (s *EventOptionStoreType) GetByID(id uint) (*models.EventOption, error) {
	var eventOption models.EventOption
	if err := s.db.First(&eventOption, id).Error; err != nil {
		return nil, errors.New("erreur lors de la récupération de l'option : " + err.Error())
	}
	return &eventOption, nil
}

func (s *EventOptionStoreType) GetByEventID(eventID uint) ([]models.EventOption, error) {
	var options []models.EventOption
	if err := s.db.Where("event_id = ?", eventID).Find(&options).Error; err != nil {
		return nil, errors.New("erreur lors de la récupération des options de l'événement : " + err.Error())
	}
	return options, nil
}
