package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type RoleStore struct {
	db *gorm.DB
}

func NewRoleStore(db *gorm.DB) *RoleStore {
	return &RoleStore{db: db}
}

// Créer un nouveau rôle
func (s *RoleStore) Create(role *models.Role) error {
	return s.db.Create(role).Error
}

// Mettre à jour un rôle existant
func (s *RoleStore) Update(role *models.Role) error {
	return s.db.Save(role).Error
}

// Supprimer un rôle
func (s *RoleStore) Delete(id uint) error {
	return s.db.Delete(&models.Role{}, id).Error
}

// Récupérer un rôle par son ID
func (s *RoleStore) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	err := s.db.First(&role, id).Error
	return &role, err
}

// Récupérer un rôle par son nom
func (s *RoleStore) GetByName(name string) (*models.Role, error) {
	var role models.Role
	err := s.db.Where("name = ?", name).First(&role).Error
	return &role, err
}
