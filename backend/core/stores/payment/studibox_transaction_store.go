package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type StudiboxTransactionStore struct {
	db *gorm.DB
}

func NewStudiboxTransactionStore(db *gorm.DB) *StudiboxTransactionStore {
	return &StudiboxTransactionStore{db: db}
}

// Créer une transaction Studibox
func (s *StudiboxTransactionStore) Create(studiboxTransaction *models.StudiboxTransaction) error {
	return s.db.Create(studiboxTransaction).Error
}

// Mettre à jour une transaction Studibox existante
func (s *StudiboxTransactionStore) Update(studiboxTransaction *models.StudiboxTransaction) error {
	return s.db.Save(studiboxTransaction).Error
}

// Supprimer une transaction Studibox
func (s *StudiboxTransactionStore) Delete(id uint) error {
	return s.db.Delete(&models.StudiboxTransaction{}, id).Error
}

// Récupérer une transaction Studibox par son ID
func (s *StudiboxTransactionStore) GetByID(id uint) (*models.StudiboxTransaction, error) {
	var studiboxTransaction models.StudiboxTransaction
	err := s.db.First(&studiboxTransaction, id).Error
	return &studiboxTransaction, err
}
