package user

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type RoleStoreType struct {
	db *gorm.DB
}

func RoleStore(db *gorm.DB) *RoleStoreType {
	return &RoleStoreType{db: db}
}

// Créer un nouveau rôle
func (s *RoleStoreType) Create(role *models.Role) error {
	return s.db.Create(role).Error
}

// Mettre à jour un rôle existant
func (s *RoleStoreType) Update(role *models.Role) error {
	return s.db.Save(role).Error
}

// Supprimer un rôle
func (s *RoleStoreType) Delete(id uint) error {
	return s.db.Delete(&models.Role{}, id).Error
}

// Récupérer un rôle par son ID
func (s *RoleStoreType) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	err := s.db.First(&role, id).Error
	return &role, err
}

// Récupérer un rôle par son nom
func (s *RoleStoreType) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := s.db.Where("name = ?", name).First(&role).Error
	return &role, err
}
