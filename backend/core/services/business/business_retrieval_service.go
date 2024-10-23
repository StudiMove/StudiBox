package business

import (
	"backend/core/models"
	stores "backend/core/stores/business"
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
	return s.store.GetByID(userID)
}

// GetAllBusinessUsers récupère tous les utilisateurs business
func (s *BusinessRetrievalService) GetAllBusinessUsers() ([]models.BusinessUser, error) {
	return s.store.GetAll()
}
