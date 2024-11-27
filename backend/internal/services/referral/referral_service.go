package referral

import (
	"backend/internal/db/models"
	"errors"

	"gorm.io/gorm"
)

// ReferralService gère les opérations liées aux parrainages
type ReferralService struct {
	DB *gorm.DB
}

// NewReferralService crée une nouvelle instance de ReferralService
func NewReferralService(db *gorm.DB) *ReferralService {
	return &ReferralService{DB: db}
}

// GetFilleulsByParrain liste les filleuls associés à un parrain donné
func (s *ReferralService) GetFilleulsByParrain(parrainID uint) ([]models.User, error) {
	var filleuls []models.User
	if err := s.DB.Joins("JOIN referrals ON referrals.filleul_id = users.id").
		Where("referrals.parrain_id = ?", parrainID).
		Find(&filleuls).Error; err != nil {
		return nil, err
	}
	return filleuls, nil
}

// CountFilleuls compte le nombre de filleuls d’un parrain donné
func (s *ReferralService) CountFilleuls(parrainID uint) (int64, error) {
	var count int64
	if err := s.DB.Model(&models.Referral{}).Where("parrain_id = ?", parrainID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// AddReferral ajoute une nouvelle relation de parrainage
func (s *ReferralService) AddReferral(parrainID uint, filleulID uint) error {
	referral := models.Referral{
		ParrainID: parrainID,
		FilleulID: filleulID,
	}
	if err := s.DB.Create(&referral).Error; err != nil {
		return errors.New("failed to create referral")
	}
	return nil
}

// GetFilleulIDsByParrain récupère uniquement les IDs des filleuls d'un parrain
func (s *ReferralService) GetFilleulIDsByParrain(parrainID uint) ([]uint, error) {
	var filleulIDs []uint

	// Utilisez une requête optimisée pour récupérer uniquement les IDs
	if err := s.DB.Model(&models.Referral{}).
		Select("filleul_id").
		Where("parrain_id = ?", parrainID).
		Scan(&filleulIDs).Error; err != nil {
		return nil, err
	}

	return filleulIDs, nil
}
