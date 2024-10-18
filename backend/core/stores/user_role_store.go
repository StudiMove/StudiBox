package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type UserRoleStore struct {
	db *gorm.DB
}

func NewUserRoleStore(db *gorm.DB) *UserRoleStore {
	return &UserRoleStore{db: db}
}

// Créer un rôle d'utilisateur
func (s *UserRoleStore) Create(userRole *models.UserRole) error {
	return s.db.Create(userRole).Error
}

// Mettre à jour un rôle d'utilisateur existant
func (s *UserRoleStore) Update(userRole *models.UserRole) error {
	return s.db.Save(userRole).Error
}

// Supprimer un rôle d'utilisateur
func (s *UserRoleStore) Delete(id uint) error {
	return s.db.Delete(&models.UserRole{}, id).Error
}

// Récupérer un rôle d'utilisateur par son ID
func (s *UserRoleStore) GetByID(id uint) (*models.UserRole, error) {
	var userRole models.UserRole
	err := s.db.First(&userRole, id).Error
	return &userRole, err
}
