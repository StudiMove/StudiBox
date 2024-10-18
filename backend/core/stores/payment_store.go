package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type PaymentStore struct {
	db *gorm.DB
}

func NewPaymentStore(db *gorm.DB) *PaymentStore {
	return &PaymentStore{db: db}
}

// Créer un paiement
func (s *PaymentStore) Create(payment *models.Payment) error {
	return s.db.Create(payment).Error
}

// Mettre à jour un paiement existant
func (s *PaymentStore) Update(payment *models.Payment) error {
	return s.db.Save(payment).Error
}

// Supprimer un paiement
func (s *PaymentStore) Delete(id uint) error {
	return s.db.Delete(&models.Payment{}, id).Error
}

// Récupérer un paiement par son ID
func (s *PaymentStore) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.First(&payment, id).Error
	return &payment, err
}
