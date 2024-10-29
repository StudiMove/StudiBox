package auth

import (
	storesUser "backend/core/stores/user"
	"fmt"
)

type AuthRoleService struct {
	userStore *storesUser.UserStore
}

// NewAuthRoleService crée une nouvelle instance de AuthRoleService
func NewAuthRoleService(userStore *storesUser.UserStore) *AuthRoleService {
	return &AuthRoleService{
		userStore: userStore,
	}
}

// CheckUserRole vérifie si l'utilisateur a le rôle spécifié
func (s *AuthRoleService) CheckUserRole(userID uint, role string) (bool, error) {
	user, err := s.userStore.PreloadRoles(userID)
	if err != nil {
		return false, fmt.Errorf("erreur lors de la récupération de l'utilisateur : %w", err)
	}

	for _, userRole := range user.Roles {
		if userRole.Name == role {
			return true, nil
		}
	}
	return false, nil
}
