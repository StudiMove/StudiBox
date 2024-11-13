package payment

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type PointHistoryStoreType struct {
	db *gorm.DB
}

func PointHistoryStore(db *gorm.DB) *PointHistoryStoreType {
	return &PointHistoryStoreType{db: db}
}

// Créer un historique de points
func (s *PointHistoryStoreType) Create(pointHistory *models.PointHistory) error {
	return s.db.Create(pointHistory).Error
}

// Mettre à jour un historique de points existant
func (s *PointHistoryStoreType) Update(pointHistory *models.PointHistory) error {
	return s.db.Save(pointHistory).Error
}

// Supprimer un historique de points
func (s *PointHistoryStoreType) Delete(id uint) error {
	return s.db.Delete(&models.PointHistory{}, id).Error
}

// Récupérer un historique de points par son ID
func (s *PointHistoryStoreType) GetByID(id uint) (*models.PointHistory, error) {
	var pointHistory models.PointHistory
	err := s.db.First(&pointHistory, id).Error
	return &pointHistory, err
}
