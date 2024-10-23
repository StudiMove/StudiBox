package business

import (
	stores "backend/core/stores/business"

	"gorm.io/gorm"
)

type BusinessService struct {
	Management *BusinessManagementService
	Retrieval  *BusinessRetrievalService
	Profile    *BusinessProfileService
}

// NewBusinessService crée une nouvelle instance de BusinessService avec ses sous-services
func NewBusinessService(db *gorm.DB) *BusinessService {
	store := stores.NewBusinessUserStore(db)
	return &BusinessService{
		Management: NewBusinessManagementService(store),
		Retrieval:  NewBusinessRetrievalService(store),
		Profile:    NewBusinessProfileService(store),
	}
}
