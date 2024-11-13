package user

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type UserPasswordStoreType struct {
	db *gorm.DB
}

func UserPasswordStore(db *gorm.DB) *UserPasswordStoreType {
	return &UserPasswordStoreType{db: db}
}

// Créer un nouveau password reset
func (s *UserPasswordStoreType) Create(userPassword *models.PasswordReset) error {
	return s.db.Create(userPassword).Error
}

// Mettre à jour un password reset existant
func (s *UserPasswordStoreType) Update(userPassword *models.PasswordReset) error {
	return s.db.Save(userPassword).Error
}

// Supprimer un password reset
func (s *UserPasswordStoreType) Delete(id uint) error {
	return s.db.Delete(&models.PasswordReset{}, id).Error
}

// Récupérer un password reset par son ID
func (s *UserPasswordStoreType) GetByID(id uint) (*models.PasswordReset, error) {
	var userPassword models.PasswordReset
	err := s.db.First(&userPassword, id).Error
	return &userPassword, err
}

// Récupérer un password reset par UserID
func (s *UserPasswordStoreType) GetByUserID(userID uint) (*models.PasswordReset, error) {
	if userID == 0 {
		return nil, errors.New("userID invalide")
	}

	var userPassword models.PasswordReset
	err := s.db.Where("user_id = ?", userID).First(&userPassword).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Renvoie nil si aucun enregistrement n'est trouvé
		}
		return nil, err
	}
	return &userPassword, nil
}

// Mettre à jour le mot de passe d'un utilisateur par son ID
func (s *UserPasswordStoreType) UpdateUserPassword(userID uint, hashedPassword string) error {
	return s.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}
