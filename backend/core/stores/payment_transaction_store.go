package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type PaymentTransactionStore struct {
	db *gorm.DB
}

func NewPaymentTransactionStore(db *gorm.DB) *PaymentTransactionStore {
	return &PaymentTransactionStore{db: db}
}

// Créer une nouvelle transaction de paiement
func (s *PaymentTransactionStore) Create(paymentTransaction *models.PaymentTransaction) error {
	return s.db.Create(paymentTransaction).Error
}

// Mettre à jour une transaction de paiement existante
func (s *PaymentTransactionStore) Update(paymentTransaction *models.PaymentTransaction) error {
	return s.db.Save(paymentTransaction).Error
}

// Supprimer une transaction de paiement
func (s *PaymentTransactionStore) Delete(id uint) error {
	return s.db.Delete(&models.PaymentTransaction{}, id).Error
}

// Récupérer une transaction de paiement par son ID
func (s *PaymentTransactionStore) GetByID(id uint) (*models.PaymentTransaction, error) {
	var paymentTransaction models.PaymentTransaction
	err := s.db.First(&paymentTransaction, id).Error
	return &paymentTransaction, err
}
