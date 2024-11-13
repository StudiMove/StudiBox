package owner

import (
	stores "backend/core/stores/owner"

	"gorm.io/gorm"
)

type OwnerServiceType struct {
	Management *OwnerManagementServiceType
	Retrieval  *OwnerRetrievalServiceType
}

// NewOwnerService crée une nouvelle instance de OwnerService avec ses sous-services
func OwnerService(db *gorm.DB) *OwnerServiceType {
	store := stores.OwnerStore(db)
	return &OwnerServiceType{
		Management: OwnerManagementService(store),
		Retrieval:  OwnerRetrievalService(store),
	}
}
