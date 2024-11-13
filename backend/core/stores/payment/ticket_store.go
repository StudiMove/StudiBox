package payment

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type TicketStoreType struct {
	db *gorm.DB
}

func TicketStore(db *gorm.DB) *TicketStoreType {
	return &TicketStoreType{db: db}
}

// Créer un ticket
func (s *TicketStoreType) Create(ticket *models.Ticket) error {
	return s.db.Create(ticket).Error
}

// Mettre à jour un ticket existant
func (s *TicketStoreType) Update(ticket *models.Ticket) error {
	return s.db.Save(ticket).Error
}

// Supprimer un ticket
func (s *TicketStoreType) Delete(id uint) error {
	return s.db.Delete(&models.Ticket{}, id).Error
}

// Récupérer un ticket par son ID
func (s *TicketStoreType) GetByID(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	err := s.db.First(&ticket, id).Error
	return &ticket, err
}
