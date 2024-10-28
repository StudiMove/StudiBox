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

	// Utiliser un map pour faire correspondre les champs à mettre à jour
	fieldsToUpdate := map[string]interface{}{
		"CompanyName": input.CompanyName,
		"Address":     input.Address,
		"City":        input.City,
		"Postcode":    input.Postcode,
		"Country":     input.Country,
		"Phone":       input.Phone,
	}

	// Appliquer les changements conditionnels
	for field, value := range fieldsToUpdate {
		if v, ok := value.(string); ok && v != "" {
			switch field {
			case "CompanyName":
				businessUser.CompanyName = v
			case "Address":
				businessUser.Address = v
			case "City":
				businessUser.City = v
			case "Postcode":
				businessUser.Postcode = v
			case "Country":
				businessUser.Country = v
			case "Phone":
				businessUser.User.Phone = v
			}
		}
	}

	// Mise à jour dans la base de données
	return s.store.Update(businessUser)
}
