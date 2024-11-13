package user

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type UserStoreType struct {
	db *gorm.DB
}

func UserStore(db *gorm.DB) *UserStoreType {
	return &UserStoreType{db: db}
}

// Créer un utilisateur
func (s *UserStoreType) Create(user *models.User) error {
	return s.db.Create(user).Error
}

// Mettre à jour un utilisateur existant
func (s *UserStoreType) Update(user *models.User) error {
	return s.db.Save(user).Error
}

// Supprimer un utilisateur
func (s *UserStoreType) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

// Mettre à jour uniquement les champs spécifiés
func (s *UserStoreType) UpdateFields(userID uint, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}
	return s.db.Model(&models.User{}).Where("id = ?", userID).Updates(fields).Error
}

// Récupérer un utilisateur par son ID avec préchargement du rôle
func (s *UserStoreType) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.Select("*").Preload("Role").First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// Récupérer un utilisateur par son adresse e-mail avec préchargement du rôle
func (s *UserStoreType) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Récupérer un utilisateur par son pseudo avec préchargement du rôle
func (s *UserStoreType) GetByPseudo(pseudo string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Role").Where("pseudo = ?", pseudo).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Récupérer tous les utilisateurs avec préchargement des rôles
func (s *UserStoreType) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Select("*").Preload("Role").Find(&users).Error
	return users, err
}

// Assigner un rôle à un utilisateur en mettant à jour le RoleID
func (s *UserStoreType) AssignRole(userID uint, roleID uint) error {
	user, err := s.GetByID(userID)
	if err != nil {
		return errors.New("erreur lors de la récupération de l'utilisateur : " + err.Error())
	}

	user.RoleID = roleID
	return s.Update(user)
}

// Précharger le rôle d'un utilisateur
func (s *UserStoreType) PreloadRole(userID uint) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Role").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
