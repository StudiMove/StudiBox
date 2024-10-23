package user

import (
	request "backend/core/api/request/user"
	"backend/core/models"
	stores "backend/core/stores/user"
	"errors"

	"gorm.io/gorm"
)

// UserManagementService gère les opérations de gestion des utilisateurs
type UserManagementService struct {
	store *stores.UserStore
}

// NewUserManagementService crée une nouvelle instance de UserManagementService
func NewUserManagementService(store *stores.UserStore) *UserManagementService {
	return &UserManagementService{
		store: store,
	}
}

// CreateUser crée un nouvel utilisateur
func (s *UserManagementService) CreateUser(user *models.User) error {
	if err := s.store.Create(user); err != nil {
		return errors.New("failed to create user: " + err.Error())
	}
	return nil
}

// UpdateUser met à jour un utilisateur existant
func (s *UserManagementService) UpdateUser(user *models.User) error {
	if err := s.store.Update(user); err != nil {
		return errors.New("failed to update user: " + err.Error())
	}
	return nil
}

// DeleteUser supprime un utilisateur par ID
func (s *UserManagementService) DeleteUser(userID uint) error {
	if err := s.store.Delete(userID); err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}
	return nil
}

// UpdateUserProfile met à jour les champs spécifiques du profil utilisateur
func (s *UserManagementService) UpdateUserProfile(userID uint, input request.UpdateUserProfileRequest) error {
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return errors.New("failed to retrieve user: " + err.Error())
	}

	// Mise à jour dynamique des champs non vides
	updated := false
	if input.FirstName != "" {
		user.FirstName = input.FirstName
		updated = true
	}
	if input.LastName != "" {
		user.LastName = input.LastName
		updated = true
	}
	if input.Email != "" {
		user.Email = input.Email
		updated = true
	}
	if input.Phone != "" {
		user.Phone = input.Phone
		updated = true
	}

	if !updated {
		return errors.New("no fields to update")
	}

	if err := s.store.Update(user); err != nil {
		return errors.New("failed to update user profile: " + err.Error())
	}
	return nil
}

// AssignUserRole assigne un rôle à un utilisateur
func (s *UserManagementService) AssignUserRole(userID uint, roleID uint) error {
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return errors.New("failed to retrieve user: " + err.Error())
	}

	role := models.UserRole{
		UserID: user.ID,
		RoleID: roleID,
	}

	if err := s.store.AssignRole(&role); err != nil {
		return errors.New("failed to assign role: " + err.Error())
	}
	return nil
}
