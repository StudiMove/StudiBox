package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type PointHistoryStore struct {
	db *gorm.DB
}

func NewPointHistoryStore(db *gorm.DB) *PointHistoryStore {
	return &PointHistoryStore{db: db}
}

// Créer un historique de points
func (s *PointHistoryStore) Create(pointHistory *models.PointHistory) error {
	return s.db.Create(pointHistory).Error
}

// Mettre à jour un historique de points existant
func (s *PointHistoryStore) Update(pointHistory *models.PointHistory) error {
	return s.db.Save(pointHistory).Error
}

// Supprimer un historique de points
func (s *PointHistoryStore) Delete(id uint) error {
	return s.db.Delete(&models.PointHistory{}, id).Error
}

// Récupérer un historique de points par son ID
func (s *PointHistoryStore) GetByID(id uint) (*models.PointHistory, error) {
	var pointHistory models.PointHistory
	err := s.db.First(&pointHistory, id).Error
	return &pointHistory, err
}
