package auth

import (
	"backend/core/models"
	"backend/core/services/user"
	storesBusiness "backend/core/stores/business"
	storesUser "backend/core/stores/user"
	"backend/core/utils"
	"fmt"
)

type AuthRegisterService struct {
	userService   *user.UserService
	businessStore *storesBusiness.BusinessUserStore
	roleStore     *storesUser.RoleStore
	userRoleStore *storesUser.UserRoleStore
}

func NewAuthRegisterService(userService *user.UserService, businessStore *storesBusiness.BusinessUserStore, roleStore *storesUser.RoleStore, userRoleStore *storesUser.UserRoleStore) *AuthRegisterService {
	return &AuthRegisterService{
		userService:   userService,
		businessStore: businessStore,
		roleStore:     roleStore,
		userRoleStore: userRoleStore,
	}
}

// RegisterUser gère l'inscription d'un utilisateur standard
func (s *AuthRegisterService) RegisterUser(user *models.User) error {
	existingUser, err := s.userService.Retrieval.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("email déjà utilisé")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	user.Password = hashedPassword

	if err := s.userService.Management.CreateUser(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	role, err := s.roleStore.GetByName("User")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle : %w", err)
	}

	if err := s.userService.Management.AssignUserRole(user.ID, role.ID); err != nil {
		return fmt.Errorf("erreur lors de l'attribution du rôle à l'utilisateur : %w", err)
	}

	return nil
}

// RegisterBusinessUser gère l'inscription d'un utilisateur entreprise avec gestion des transactions
func (s *AuthRegisterService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
	// Vérifie si l'email est déjà utilisé
	existingUser, err := s.userService.Retrieval.GetUserByEmail(businessUser.User.Email)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("email déjà utilisé")
	}

	// Hash du mot de passe
	hashedPassword, err := utils.HashPassword(businessUser.User.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	businessUser.User.Password = hashedPassword

	// Création de l'utilisateur
	if err := s.userService.Management.CreateUser(&businessUser.User); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	// Associe l'utilisateur à un BusinessUser
	businessUser.UserID = businessUser.User.ID
	if err := s.businessStore.Create(businessUser); err != nil {
		return fmt.Errorf("erreur lors de la création du BusinessUser : %w", err)
	}

	// Récupération du rôle "Business" et assignation
	role, err := s.roleStore.GetByName("Business")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle 'business' : %w", err)
	}

	if err := s.userService.Management.AssignUserRole(businessUser.User.ID, role.ID); err != nil {
		return fmt.Errorf("erreur lors de l'attribution du rôle à l'utilisateur : %w", err)
	}

	return nil
}
