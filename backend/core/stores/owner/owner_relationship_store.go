package owner

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type OwnerRelationshipStoreType struct {
	db *gorm.DB
}

func OwnerRelationshipStore(db *gorm.DB) *OwnerRelationshipStoreType {
	return &OwnerRelationshipStoreType{db: db}
}

// Créer une nouvelle relation Owner
func (s *OwnerRelationshipStoreType) Create(relationship *models.OwnerRelationship) error {
	return s.db.Create(relationship).Error
}

// Supprimer une relation Owner par ID
func (s *OwnerRelationshipStoreType) Delete(id uint) error {
	return s.db.Delete(&models.OwnerRelationship{}, id).Error
}

// Récupérer toutes les relations pour une SchoolID spécifique
func (s *OwnerRelationshipStoreType) GetBySchoolID(schoolID uint) ([]models.OwnerRelationship, error) {
	var relationships []models.OwnerRelationship
	err := s.db.Where("school_id = ?", schoolID).Find(&relationships).Error
	return relationships, err
}

// Récupérer toutes les relations pour une AssociationID spécifique
func (s *OwnerRelationshipStoreType) GetByAssociationID(associationID uint) ([]models.OwnerRelationship, error) {
	var relationships []models.OwnerRelationship
	err := s.db.Where("association_id = ?", associationID).Find(&relationships).Error
	return relationships, err
}
