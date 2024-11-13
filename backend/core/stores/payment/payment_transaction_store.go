package payment

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type PaymentTransactionStoreType struct {
	db *gorm.DB
}

func PaymentTransactionStore(db *gorm.DB) *PaymentTransactionStoreType {
	return &PaymentTransactionStoreType{db: db}
}

// Créer une nouvelle transaction de paiement
func (s *PaymentTransactionStoreType) Create(paymentTransaction *models.PaymentTransaction) error {
	return s.db.Create(paymentTransaction).Error
}

// Mettre à jour une transaction de paiement existante
func (s *PaymentTransactionStoreType) Update(paymentTransaction *models.PaymentTransaction) error {
	return s.db.Save(paymentTransaction).Error
}

// Supprimer une transaction de paiement
func (s *PaymentTransactionStoreType) Delete(id uint) error {
	return s.db.Delete(&models.PaymentTransaction{}, id).Error
}

// Récupérer une transaction de paiement par son ID
func (s *PaymentTransactionStoreType) GetByID(id uint) (*models.PaymentTransaction, error) {
	var paymentTransaction models.PaymentTransaction
	err := s.db.First(&paymentTransaction, id).Error
	return &paymentTransaction, err
}
