package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type AssociationStore struct {
	db *gorm.DB
}

func NewAssociationStore(db *gorm.DB) *AssociationStore {
	return &AssociationStore{db: db}
}

// Créer une nouvelle association
func (s *AssociationStore) Create(association *models.Association) error {
	return s.db.Create(association).Error
}

// Mettre à jour une association existante
func (s *AssociationStore) Update(association *models.Association) error {
	return s.db.Save(association).Error
}

// Supprimer une association
func (s *AssociationStore) Delete(id uint) error {
	return s.db.Delete(&models.Association{}, id).Error
}

// Récupérer une association par son ID
func (s *AssociationStore) GetByID(id uint) (*models.Association, error) {
	var association models.Association
	err := s.db.First(&association, id).Error
	return &association, err
}
