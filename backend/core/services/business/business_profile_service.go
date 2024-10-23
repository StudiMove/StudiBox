package business

import (
	"backend/core/models"
	stores "backend/core/stores/business"
	"errors"

	"gorm.io/gorm"
)

type BusinessProfileService struct {
	store *stores.BusinessUserStore
}

// NewBusinessProfileService crée une nouvelle instance de BusinessProfileService
func NewBusinessProfileService(store *stores.BusinessUserStore) *BusinessProfileService {
	return &BusinessProfileService{
		store: store,
	}
}

// GetBusinessUserByID récupère un utilisateur business par son ID
func (s *BusinessProfileService) GetBusinessUserByID(userID uint) (*models.BusinessUser, error) {
	businessUser, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("business user not found")
		}
		return nil, err
	}
	return businessUser, nil
}
