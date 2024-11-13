package auth

import (
	request "backend/core/api/request/auth"
	"backend/core/services/user"
	storesOwner "backend/core/stores/owner"
	storesUser "backend/core/stores/user"
	"backend/core/utils"
	"backend/database/models"
	"fmt"
)

type AuthRegisterServiceType struct {
	userService       *user.UserServiceType
	ownerStore        *storesOwner.OwnerStoreType
	roleStore         *storesUser.RoleStoreType
	ownerRelationship *storesOwner.OwnerRelationshipStoreType
}

func AuthRegisterService(
	userService *user.UserServiceType,
	ownerStore *storesOwner.OwnerStoreType,
	roleStore *storesUser.RoleStoreType,
	ownerRelationship *storesOwner.OwnerRelationshipStoreType,
) *AuthRegisterServiceType {
	return &AuthRegisterServiceType{
		userService:       userService,
		ownerStore:        ownerStore,
		roleStore:         roleStore,
		ownerRelationship: ownerRelationship,
	}
}

// RegisterUser gère l'inscription d'un utilisateur standard
func (s *AuthRegisterServiceType) RegisterUser(req *request.RegisterUserRequest) error {
	// Vérification si l'email est déjà utilisé
	existingUser, err := s.userService.Retrieval.GetUserByEmail(req.Email)
	if err != nil && err.Error() != "user not found" {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("l'email est déjà utilisé")
	}

	// Vérification si le pseudo est déjà utilisé
	existingPseudo, err := s.userService.Retrieval.GetUserByPseudo(req.Pseudo)
	if err != nil && err.Error() != "user not found" {
		return fmt.Errorf("erreur lors de la vérification du pseudo : %w", err)
	}
	if existingPseudo != nil {
		return fmt.Errorf("le pseudo est déjà utilisé")
	}

	// Hachage du mot de passe
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}

	// Récupérer le rôle
	role, err := s.roleStore.GetByName("User")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle 'User' : %w", err)
	}

	// Créer un nouvel utilisateur
	user := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Pseudo:      req.Pseudo,
		Email:       req.Email,
		Password:    hashedPassword,
		Phone:       req.Phone,
		ProfileType: req.ProfileType,
		RoleID:      role.ID,
	}

	// Appeler la fonction CreateUser pour insérer dans la base de données
	if err := s.userService.Management.CreateUser(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	return nil
}

// RegisterOwnerUser gère l'inscription d'un utilisateur propriétaire (Owner)
func (s *AuthRegisterServiceType) RegisterOwnerUser(req *request.RegisterOwnerRequest) error {
	// Vérification de l'existence de l'utilisateur par email et pseudo
	existingUser, err := s.userService.Retrieval.GetUserByEmail(req.Email)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("l'email est déjà utilisé")
	}

	existingPseudo, err := s.userService.Retrieval.GetUserByPseudo(req.Pseudo)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification du pseudo : %w", err)
	}
	if existingPseudo != nil {
		return fmt.Errorf("le pseudo est déjà utilisé")
	}

	// Hachage du mot de passe
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}

	// Récupération du rôle "Owner"
	role, err := s.roleStore.GetByName("Owner")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle 'Owner' : %w", err)
	}

	// Création de l'utilisateur avec le role_id et pseudo définis
	user := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Pseudo:      req.Pseudo,
		Email:       req.Email,
		Phone:       req.Phone,
		Password:    hashedPassword,
		ProfileType: models.ProfileTypeNonStudent,
		RoleID:      role.ID,
	}

	// Sauvegarde de l'utilisateur
	if err := s.userService.Management.CreateUser(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	// Création de l'Owner
	owner := &models.Owner{
		User:        *user,
		CompanyName: req.CompanyName,
		Address:     req.Address,
		PostalCode:  req.PostalCode,
		City:        req.City,
		Country:     req.Country,
		Region:      req.Region,
		Description: req.Description,
		Type:        req.Type,
	}

	// Sauvegarde de l'Owner
	if err := s.ownerStore.Create(owner); err != nil {
		return fmt.Errorf("erreur lors de la création de l'Owner : %w", err)
	}

	// Création de la relation avec une école si nécessaire
	if req.Type == models.OwnerTypeAssociation && req.SchoolID > 0 {
		relationship := &models.OwnerRelationship{
			SchoolID:      req.SchoolID,
			AssociationID: owner.ID,
		}
		if err := s.ownerRelationship.Create(relationship); err != nil {
			return fmt.Errorf("erreur lors de la création de la relation avec l'école : %w", err)
		}
	}
	return nil
}
