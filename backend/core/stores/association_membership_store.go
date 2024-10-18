package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type AssociationMembershipStore struct {
	db *gorm.DB
}

// NewAssociationMembershipStore crée un nouveau store pour gérer les associations membres
func NewAssociationMembershipStore(db *gorm.DB) *AssociationMembershipStore {
	return &AssociationMembershipStore{db: db}
}

// Create ajoute une nouvelle association membre dans la base de données
func (s *AssociationMembershipStore) Create(membership *models.AssociationMembership) error {
	return s.db.Create(membership).Error
}

// FindByID cherche une association membre par ID
func (s *AssociationMembershipStore) FindByID(id uint) (*models.AssociationMembership, error) {
	var membership models.AssociationMembership
	err := s.db.First(&membership, id).Error
	if err != nil {
		return nil, err
	}
	return &membership, nil
}

// Update modifie une association membre existante
func (s *AssociationMembershipStore) Update(membership *models.AssociationMembership) error {
	return s.db.Save(membership).Error
}

// Delete supprime une association membre par ID
func (s *AssociationMembershipStore) Delete(id uint) error {
	return s.db.Delete(&models.AssociationMembership{}, id).Error
}

// FindAll retourne toutes les associations membres
func (s *AssociationMembershipStore) FindAll() ([]models.AssociationMembership, error) {
	var memberships []models.AssociationMembership
	err := s.db.Find(&memberships).Error
	if err != nil {
		return nil, err
	}
	return memberships, nil
}
