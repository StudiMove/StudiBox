// package stores

package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

// Créer un utilisateur
func (s *UserStore) Create(user *models.User) error {
	return s.db.Create(user).Error
}

// Mettre à jour un utilisateur existant
func (s *UserStore) Update(user *models.User) error {
	return s.db.Save(user).Error
}

// Supprimer un utilisateur
func (s *UserStore) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

// Récupérer un utilisateur par son ID
func (s *UserStore) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	return &user, err
}

// Récupérer un utilisateur par son adresse e-mail
func (s *UserStore) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Récupérer tous les utilisateurs
func (s *UserStore) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}

// Assigner un rôle à un utilisateur
func (s *UserStore) AssignRole(userRole *models.UserRole) error {
	return s.db.Create(userRole).Error
}
