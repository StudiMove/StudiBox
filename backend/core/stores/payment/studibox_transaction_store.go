package payment

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type StudiboxTransactionStoreType struct {
	db *gorm.DB
}

func StudiboxTransactionStore(db *gorm.DB) *StudiboxTransactionStoreType {
	return &StudiboxTransactionStoreType{db: db}
}

// Créer une transaction Studibox
func (s *StudiboxTransactionStoreType) Create(studiboxTransaction *models.StudiboxTransaction) error {
	return s.db.Create(studiboxTransaction).Error
}

// Mettre à jour une transaction Studibox existante
func (s *StudiboxTransactionStoreType) Update(studiboxTransaction *models.StudiboxTransaction) error {
	return s.db.Save(studiboxTransaction).Error
}

// Supprimer une transaction Studibox
func (s *StudiboxTransactionStoreType) Delete(id uint) error {
	return s.db.Delete(&models.StudiboxTransaction{}, id).Error
}

// Récupérer une transaction Studibox par son ID
func (s *StudiboxTransactionStoreType) GetByID(id uint) (*models.StudiboxTransaction, error) {
	var studiboxTransaction models.StudiboxTransaction
	err := s.db.First(&studiboxTransaction, id).Error
	return &studiboxTransaction, err
}
