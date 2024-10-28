package business

import (
	"backend/core/models"
	stores "backend/core/stores/business"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BusinessRetrievalService struct {
	store *stores.BusinessUserStore
}

// NewBusinessRetrievalService crée une nouvelle instance de BusinessRetrievalService
func NewBusinessRetrievalService(store *stores.BusinessUserStore) *BusinessRetrievalService {
	return &BusinessRetrievalService{
		store: store,
	}
}

// GetBusinessUserByID récupère un utilisateur business par son ID
func (s *BusinessRetrievalService) GetBusinessUserByID(userID uint) (*models.BusinessUser, error) {
	businessUser, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("business user with ID %d not found: %w", userID, err)
		}
		return nil, fmt.Errorf("error retrieving business user with ID %d: %w", userID, err)
	}
	return businessUser, nil
}

// GetAllBusinessUsers récupère tous les utilisateurs business
func (s *BusinessRetrievalService) GetAllBusinessUsers() ([]models.BusinessUser, error) {
	businessUsers, err := s.store.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving all business users: %w", err)
	}
	return businessUsers, nil
}
