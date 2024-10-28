package user

import (
	stores "backend/core/stores/user"

	"gorm.io/gorm"
)

type UserService struct {
	Management *UserManagementService
	Retrieval  *UserRetrievalService
}

// NewUserService crée une nouvelle instance de UserService avec ses sous-services
func NewUserService(db *gorm.DB) *UserService {
	store := stores.NewUserStore(db)
	return &UserService{
		Management: NewUserManagementService(store),
		Retrieval:  NewUserRetrievalService(store),
	}
}
