package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EducationalInstitutionStore struct {
	db *gorm.DB
}

func NewEducationalInstitutionStore(db *gorm.DB) *EducationalInstitutionStore {
	return &EducationalInstitutionStore{db: db}
}

// Créer un établissement éducatif
func (s *EducationalInstitutionStore) Create(educationalInstitution *models.EducationalInstitution) error {
	return s.db.Create(educationalInstitution).Error
}

// Mettre à jour un établissement éducatif existant
func (s *EducationalInstitutionStore) Update(educationalInstitution *models.EducationalInstitution) error {
	return s.db.Save(educationalInstitution).Error
}

// Supprimer un établissement éducatif
func (s *EducationalInstitutionStore) Delete(id uint) error {
	return s.db.Delete(&models.EducationalInstitution{}, id).Error
}

// Récupérer un établissement éducatif par son ID
func (s *EducationalInstitutionStore) GetByID(id uint) (*models.EducationalInstitution, error) {
	var educationalInstitution models.EducationalInstitution
	err := s.db.First(&educationalInstitution, id).Error
	return &educationalInstitution, err
}

// Récupérer toutes les institutions éducatives
func (s *EducationalInstitutionStore) GetAll() ([]models.EducationalInstitution, error) {
	var institutions []models.EducationalInstitution
	err := s.db.Find(&institutions).Error
	return institutions, err
}
