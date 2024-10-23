package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type SchoolMembershipStore struct {
	db *gorm.DB
}

func NewSchoolMembershipStore(db *gorm.DB) *SchoolMembershipStore {
	return &SchoolMembershipStore{db: db}
}

// Créer une nouvelle adhésion scolaire
func (s *SchoolMembershipStore) Create(schoolMembership *models.SchoolMembership) error {
	return s.db.Create(schoolMembership).Error
}

// Mettre à jour une adhésion scolaire existante
func (s *SchoolMembershipStore) Update(schoolMembership *models.SchoolMembership) error {
	return s.db.Save(schoolMembership).Error
}

// Supprimer une adhésion scolaire
func (s *SchoolMembershipStore) Delete(id uint) error {
	return s.db.Delete(&models.SchoolMembership{}, id).Error
}

// Récupérer une adhésion scolaire par son ID
func (s *SchoolMembershipStore) GetByID(id uint) (*models.SchoolMembership, error) {
	var schoolMembership models.SchoolMembership
	err := s.db.First(&schoolMembership, id).Error
	return &schoolMembership, err
}
