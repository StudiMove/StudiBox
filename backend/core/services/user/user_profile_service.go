package user

import (
	"backend/core/models"
	stores "backend/core/stores/user"
	"errors"

	"gorm.io/gorm"
)

type UserProfileService struct {
	store *stores.UserStore
}

// NewUserProfileService crée une nouvelle instance de UserProfileService
func NewUserProfileService(store *stores.UserStore) *UserProfileService {
	return &UserProfileService{
		store: store,
	}
}

// GetUserByID récupère un utilisateur par son ID
func (s *UserProfileService) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
