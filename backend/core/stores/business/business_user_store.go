package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

// BusinessUserStore représente le store pour les utilisateurs professionnels
type BusinessUserStore struct {
	db *gorm.DB
}

// NewBusinessUserStore crée une nouvelle instance de BusinessUserStore
func NewBusinessUserStore(db *gorm.DB) *BusinessUserStore {
	return &BusinessUserStore{db: db}
}

// Créer un utilisateur professionnel
func (s *BusinessUserStore) Create(businessUser *models.BusinessUser) error {
	return s.db.Create(businessUser).Error
}

// Mettre à jour un utilisateur professionnel existant
func (s *BusinessUserStore) Update(businessUser *models.BusinessUser) error {
	return s.db.Save(businessUser).Error
}

// Supprimer un utilisateur professionnel
func (s *BusinessUserStore) Delete(id uint) error {
	return s.db.Delete(&models.BusinessUser{}, id).Error
}

// Récupérer un utilisateur professionnel par son ID avec les informations associées
func (s *BusinessUserStore) GetByID(id uint) (*models.BusinessUser, error) {
	var businessUser models.BusinessUser
	err := s.db.Preload("User").First(&businessUser, id).Error
	return &businessUser, err
}

// Récupérer tous les utilisateurs professionnels
func (s *BusinessUserStore) GetAll() ([]models.BusinessUser, error) {
	var businessUsers []models.BusinessUser
	err := s.db.Preload("User").Find(&businessUsers).Error
	return businessUsers, err
}
