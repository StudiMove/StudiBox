package owner

import (
	request "backend/core/api/request/owner"
	stores "backend/core/stores/owner"
	"errors"
	"fmt"
)

type OwnerManagementServiceType struct {
	store *stores.OwnerStoreType
}

func OwnerManagementService(store *stores.OwnerStoreType) *OwnerManagementServiceType {
	return &OwnerManagementServiceType{
		store: store,
	}
}

// UpdateOwnerProfile met à jour un owner par ID (Admin seulement)
func (s *OwnerManagementServiceType) UpdateOwnerProfile(userID uint, input request.UpdateOwnerRequest) error {
	owner, err := s.store.GetByUserID(userID)
	if err != nil {
		return fmt.Errorf("owner not found: %w", err)
	}

	// Vérifier les doublons pour le champ SIRET
	if input.SIRET != "" && input.SIRET != owner.SIRET {
		existingOwner, err := s.store.GetBySIRET(input.SIRET)
		if err == nil && existingOwner.ID != owner.ID {
			return errors.New("le SIRET est déjà utilisé par un autre propriétaire")
		}
	}

	// Préparer la map des champs à mettre à jour
	updates := map[string]interface{}{
		"company_name": input.CompanyName,
		"address":      input.Address,
		"city":         input.City,
		"postal_code":  input.PostalCode,
		"country":      input.Country,
		"region":       input.Region,
		"siret":        input.SIRET,
		"description":  input.Description,
		"status":       input.Status,
		"type":         input.Type,
	}

	// Vérifier si le numéro de téléphone est modifié
	if input.Phone != 0 && input.Phone != owner.Phone {
		updates["phone"] = input.Phone
	}

	// Supprimer les champs vides ou non définis
	for key, value := range updates {
		if value == "" || value == nil || value == 0 {
			delete(updates, key)
		}
	}

	if len(updates) == 0 {
		return nil
	}

	return s.store.UpdateFields(owner.ID, updates)
}
