package business

import (
	request "backend/core/api/request/business"
	stores "backend/core/stores/business"
)

type BusinessManagementService struct {
	store *stores.BusinessUserStore
}

// NewBusinessManagementService crée une nouvelle instance de BusinessManagementService
func NewBusinessManagementService(store *stores.BusinessUserStore) *BusinessManagementService {
	return &BusinessManagementService{
		store: store,
	}
}

// UpdateBusinessUserProfile met à jour un business par ID (Admin seulement)
func (s *BusinessManagementService) UpdateBusinessUserProfile(userID uint, input request.UpdateBusinessProfileRequest) error {
	businessUser, err := s.store.GetByID(userID)
	if err != nil {
		return err
	}

	// Mise à jour des champs de l'objet businessUser
	if input.CompanyName != "" {
		businessUser.CompanyName = input.CompanyName
	}
	if input.Address != "" {
		businessUser.Address = input.Address
	}
	if input.City != "" {
		businessUser.City = input.City
	}
	if input.Postcode != "" {
		businessUser.Postcode = input.Postcode
	}
	if input.Country != "" {
		businessUser.Country = input.Country
	}
	if input.Phone != "" {
		businessUser.User.Phone = input.Phone
	}

	return s.store.Update(businessUser)
}
