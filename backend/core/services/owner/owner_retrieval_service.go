package owner

import (
	stores "backend/core/stores/owner"
	"backend/database/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OwnerRetrievalServiceType struct {
	store *stores.OwnerStoreType
}

// NewOwnerRetrievalService crée une nouvelle instance de OwnerRetrievalService
func OwnerRetrievalService(store *stores.OwnerStoreType) *OwnerRetrievalServiceType {
	return &OwnerRetrievalServiceType{
		store: store,
	}
}

// GetOwnerUserByID récupère un utilisateur owner par son ID
func (s *OwnerRetrievalServiceType) GetOwnerByID(userID uint) (*models.Owner, error) {
	ownerUser, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("owner user with ID %d not found: %w", userID, err)
		}
		return nil, fmt.Errorf("error retrieving owner user with ID %d: %w", userID, err)
	}
	return ownerUser, nil
}

// GetOwnerByUserID récupère un owner par l'ID de l'utilisateur
func (s *OwnerRetrievalServiceType) GetOwnerByUserID(userID uint) (*models.Owner, error) {
	owner, err := s.store.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("erreur lors de la récupération du propriétaire par ID utilisateur : %w", err)
		}
		return nil, fmt.Errorf("erreur lors de la récupération du propriétaire par ID utilisateur : %w", err)
	}
	return owner, nil
}

// GetAllOwnerUsers récupère tous les utilisateurs owner
func (s *OwnerRetrievalServiceType) GetAllOwnerUsers() ([]models.Owner, error) {
	ownerUsers, err := s.store.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving all owner users: %w", err)
	}
	return ownerUsers, nil
}

// GetActiveOrganisations récupère les Owners actifs (Validés) avec gestion des erreurs.
func (s *OwnerRetrievalServiceType) GetActiveOrganisations() ([]models.Owner, error) {
	owners, err := s.store.GetByStatus(models.StatusValidated)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des Owners actifs: %w", err)
	}
	return owners, nil
}

// GetInactiveOrganisations récupère les Owners inactifs avec gestion des erreurs.
func (s *OwnerRetrievalServiceType) GetInactiveOrganisations() ([]models.Owner, error) {
	owners, err := s.store.GetByStatus(models.StatusInactive)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des Owners inactifs: %w", err)
	}
	return owners, nil
}

// GetPendingOrganisations récupère les Owners en attente avec gestion des erreurs.
func (s *OwnerRetrievalServiceType) GetPendingOrganisations() ([]models.Owner, error) {
	owners, err := s.store.GetByStatus(models.StatusPending)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des Owners en attente: %w", err)
	}
	return owners, nil
}

// GetSuspendedOrganisations récupère les Owners suspendus avec gestion des erreurs.
func (s *OwnerRetrievalServiceType) GetSuspendedOrganisations() ([]models.Owner, error) {
	owners, err := s.store.GetByStatus(models.StatusRejected)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des Owners suspendus: %w", err)
	}
	return owners, nil
}

// GetOwnerByUserEmail récupère un owner par l'email de l'utilisateur
func (s *OwnerRetrievalServiceType) GetOwnerByUserEmail(email string) (*models.Owner, error) {
	owner, err := s.store.GetByUserEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Retourne nil si aucun owner n'est trouvé pour cet email
		}
		return nil, fmt.Errorf("erreur lors de la récupération du propriétaire par email : %w", err)
	}
	return owner, nil
}
