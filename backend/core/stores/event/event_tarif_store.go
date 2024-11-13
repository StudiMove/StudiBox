package event

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type EventTarifStoreType struct {
	db *gorm.DB
}

// NewEventTarifStore instancie un store pour les tarifs d'événements
func EventTarifStore(db *gorm.DB) *EventTarifStoreType {
	return &EventTarifStoreType{db: db}
}

func (s *EventTarifStoreType) Create(tarif *models.EventTarif) error {
	if err := s.db.Create(tarif).Error; err != nil {
		return errors.New("erreur lors de la création du tarif : " + err.Error())
	}
	return nil
}

func (s *EventTarifStoreType) Update(tarif *models.EventTarif) error {
	if err := s.db.Save(tarif).Error; err != nil {
		return errors.New("erreur lors de la mise à jour du tarif : " + err.Error())
	}
	return nil
}

func (s *EventTarifStoreType) Delete(id uint) error {
	if err := s.db.Delete(&models.EventTarif{}, id).Error; err != nil {
		return errors.New("erreur lors de la suppression du tarif : " + err.Error())
	}
	return nil
}

func (s *EventTarifStoreType) DeleteByEventIDWithTx(eventID uint, tx *gorm.DB) error {
	if err := tx.Where("event_id = ?", eventID).Delete(&models.EventTarif{}).Error; err != nil {
		return errors.New("Erreur lors de la suppression des tarifs : " + err.Error())
	}
	return nil
}

func (s *EventTarifStoreType) GetByEventID(eventID uint) ([]models.EventTarif, error) {
	var tarifs []models.EventTarif
	if err := s.db.Where("event_id = ?", eventID).Find(&tarifs).Error; err != nil {
		return nil, errors.New("erreur lors de la récupération des tarifs de l'événement : " + err.Error())
	}
	return tarifs, nil
}

func (s *EventTarifStoreType) GetByID(id uint) (*models.EventTarif, error) {
	var tarif models.EventTarif
	if err := s.db.First(&tarif, id).Error; err != nil {
		return nil, errors.New("erreur lors de la récupération du tarif : " + err.Error())
	}
	return &tarif, nil
}
