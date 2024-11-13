package owner

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type OwnerMembershipStoreType struct {
	db *gorm.DB
}

func OwnerMembershipStore(db *gorm.DB) *OwnerMembershipStoreType {
	return &OwnerMembershipStoreType{db: db}
}

// Créer un nouveau OwnerMembership
func (s *OwnerMembershipStoreType) Create(membership *models.OwnerMembership) error {
	return s.db.Create(membership).Error
}

// Mettre à jour un OwnerMembership existant
func (s *OwnerMembershipStoreType) Update(membership *models.OwnerMembership) error {
	return s.db.Save(membership).Error
}

// Supprimer un OwnerMembership par ID
func (s *OwnerMembershipStoreType) Delete(id uint) error {
	return s.db.Delete(&models.OwnerMembership{}, id).Error
}

// Récupérer un OwnerMembership par ID
func (s *OwnerMembershipStoreType) GetByID(id uint) (*models.OwnerMembership, error) {
	var membership models.OwnerMembership
	err := s.db.First(&membership, id).Error
	return &membership, err
}

// Récupérer toutes les adhésions pour un Owner spécifique
func (s *OwnerMembershipStoreType) GetByOwnerID(ownerID uint) ([]models.OwnerMembership, error) {
	var memberships []models.OwnerMembership
	err := s.db.Where("owner_id = ?", ownerID).Find(&memberships).Error
	return memberships, err
}
