package user

import (
	stores "backend/core/stores/user"

	"gorm.io/gorm"
)

type UserServiceType struct {
	Management   *UserManagementServiceType
	Retrieval    *UserRetrievalServiceType
	UserPassword *UserPasswordServiceType
}

func UserService(db *gorm.DB) *UserServiceType {
	userStore := stores.UserStore(db)
	userPasswordStore := stores.UserPasswordStore(db)

	return &UserServiceType{
		Management:   UserManagementService(userStore),
		Retrieval:    UserRetrievalService(userStore),
		UserPassword: UserPasswordService(userPasswordStore),
	}
}
