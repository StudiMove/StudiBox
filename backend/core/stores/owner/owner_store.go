package owner

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type OwnerStoreType struct {
	db *gorm.DB
}

func OwnerStore(db *gorm.DB) *OwnerStoreType {
	return &OwnerStoreType{db: db}
}

// Créer un nouvel Owner
func (s *OwnerStoreType) Create(owner *models.Owner) error {
	return s.db.Create(owner).Error
}

// Mettre à jour un Owner existant
func (s *OwnerStoreType) Update(owner *models.Owner) error {
	return s.db.Save(owner).Error
}

// Mettre à jour uniquement les champs spécifiés pour un Owner
func (s *OwnerStoreType) UpdateFields(ownerID uint, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}
	return s.db.Model(&models.Owner{}).Where("id = ?", ownerID).Updates(fields).Error
}

// Supprimer un Owner par ID
func (s *OwnerStoreType) Delete(id uint) error {
	return s.db.Delete(&models.Owner{}, id).Error
}

// Récupérer un Owner par son ID
func (s *OwnerStoreType) GetByID(id uint) (*models.Owner, error) {
	var owner models.Owner
	err := s.db.First(&owner, id).Error
	return &owner, err
}

// Récupérer tous les Owners
func (s *OwnerStoreType) GetAll() ([]models.Owner, error) {
	var owners []models.Owner
	err := s.db.Find(&owners).Error
	return owners, err
}

// Récupérer tous les Owners par statut
func (s *OwnerStoreType) GetByStatus(status string) ([]models.Owner, error) {
	var owners []models.Owner
	err := s.db.Where("status = ?", status).Find(&owners).Error
	return owners, err
}

// Récupérer un Owner en utilisant l'email de l'utilisateur associé
func (s *OwnerStoreType) GetByUserEmail(email string) (*models.Owner, error) {
	var owner models.Owner
	err := s.db.Joins("JOIN users ON owners.id = users.owner_id").
		Where("users.email = ?", email).
		Where("owners.deleted_at IS NULL").
		First(&owner).Error
	return &owner, err
}

// Récupère un Owner en utilisant l'ID de l'utilisateur associé
func (s *OwnerStoreType) GetByUserID(userID uint) (*models.Owner, error) {
	var owner models.Owner
	err := s.db.Joins("JOIN users ON users.owner_id = owners.id").
		Where("users.id = ?", userID).
		First(&owner).Error
	return &owner, err
}

// Récupérer un Owner par SIRET
func (s *OwnerStoreType) GetBySIRET(siret string) (*models.Owner, error) {
	var owner models.Owner
	err := s.db.Where("siret = ?", siret).First(&owner).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("SIRET non trouvé")
	}
	return &owner, err
}
