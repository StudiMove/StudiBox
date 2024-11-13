package payment

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type PaymentStoreType struct {
	db *gorm.DB
}

func PaymentStore(db *gorm.DB) *PaymentStoreType {
	return &PaymentStoreType{db: db}
}

// Créer un paiement
func (s *PaymentStoreType) Create(payment *models.Payment) error {
	return s.db.Create(payment).Error
}

// Mettre à jour un paiement existant
func (s *PaymentStoreType) Update(payment *models.Payment) error {
	return s.db.Save(payment).Error
}

// Supprimer un paiement
func (s *PaymentStoreType) Delete(id uint) error {
	return s.db.Delete(&models.Payment{}, id).Error
}

// Récupérer un paiement par son ID
func (s *PaymentStoreType) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.First(&payment, id).Error
	return &payment, err
}
