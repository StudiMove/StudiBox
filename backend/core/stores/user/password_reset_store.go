package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type PasswordResetStore struct {
	db *gorm.DB
}

func NewPasswordResetStore(db *gorm.DB) *PasswordResetStore {
	return &PasswordResetStore{db: db}
}

// Créer un nouveau password reset
func (s *PasswordResetStore) Create(passwordReset *models.PasswordReset) error {
	return s.db.Create(passwordReset).Error
}

// Mettre à jour un password reset existant
func (s *PasswordResetStore) Update(passwordReset *models.PasswordReset) error {
	return s.db.Save(passwordReset).Error
}

// Supprimer un password reset
func (s *PasswordResetStore) Delete(id uint) error {
	return s.db.Delete(&models.PasswordReset{}, id).Error
}

// Récupérer un password reset par son ID
func (s *PasswordResetStore) GetByID(id uint) (*models.PasswordReset, error) {
	var passwordReset models.PasswordReset
	err := s.db.First(&passwordReset, id).Error
	return &passwordReset, err
}
