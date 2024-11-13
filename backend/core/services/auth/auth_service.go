package auth

import (
	serviceUser "backend/core/services/user"
	storesOwner "backend/core/stores/owner"
	storesUser "backend/core/stores/user"

	"gorm.io/gorm"
)

type AuthServiceType struct {
	Register *AuthRegisterServiceType
	Login    *AuthLoginServiceType
}

// NewAuthService crée une nouvelle instance de AuthService avec ses sous-services
func AuthService(db *gorm.DB) *AuthServiceType {
	userService := serviceUser.UserService(db)
	ownersStore := storesOwner.OwnerStore(db)
	roleStore := storesUser.RoleStore(db)
	ownerRelationship := storesOwner.OwnerRelationshipStore(db)

	return &AuthServiceType{
		Register: AuthRegisterService(userService, ownersStore, roleStore, ownerRelationship),
		Login:    AuthLoginService(userService),
	}
}
