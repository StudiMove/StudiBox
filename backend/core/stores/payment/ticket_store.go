package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type TicketStore struct {
	db *gorm.DB
}

func NewTicketStore(db *gorm.DB) *TicketStore {
	return &TicketStore{db: db}
}

// Créer un ticket
func (s *TicketStore) Create(ticket *models.Ticket) error {
	return s.db.Create(ticket).Error
}

// Mettre à jour un ticket existant
func (s *TicketStore) Update(ticket *models.Ticket) error {
	return s.db.Save(ticket).Error
}

// Supprimer un ticket
func (s *TicketStore) Delete(id uint) error {
	return s.db.Delete(&models.Ticket{}, id).Error
}

// Récupérer un ticket par son ID
func (s *TicketStore) GetByID(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	err := s.db.First(&ticket, id).Error
	return &ticket, err
}
