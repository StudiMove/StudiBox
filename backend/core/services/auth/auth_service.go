package auth

import (
	"backend/core/services/user"
	storesBusiness "backend/core/stores/business"
	storesUser "backend/core/stores/user"

	"gorm.io/gorm"
)

type AuthService struct {
	Register *AuthRegisterService
	Login    *AuthLoginService
	Role     *AuthRoleService
}

// NewAuthService crée une nouvelle instance de AuthService avec ses sous-services
func NewAuthService(db *gorm.DB) *AuthService {
	userService := user.NewUserService(db)
	businessStore := storesBusiness.NewBusinessUserStore(db)
	userStore := storesUser.NewUserStore(db)
	roleStore := storesUser.NewRoleStore(db)
	userRoleStore := storesUser.NewUserRoleStore(db)

	return &AuthService{
		Register: NewAuthRegisterService(userService, businessStore, roleStore, userRoleStore),
		Login:    NewAuthLoginService(userService),
		Role:     NewAuthRoleService(userStore),
	}
}
