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

	// Utiliser un map pour faire correspondre les champs à mettre à jour
	fieldsToUpdate := map[string]interface{}{
		"FirstName": input.FirstName,
		"LastName":  input.LastName,
		"Email":     input.Email,
		"Phone":     input.Phone,
	}

	// Appliquer les changements conditionnels
	for field, value := range fieldsToUpdate {
		if v, ok := value.(string); ok && v != "" {
			switch field {
			case "FirstName":
				user.FirstName = v
			case "LastName":
				user.LastName = v
			case "Email":
				user.Email = v
			case "Phone":
				user.Phone = v
			}
		}
	}

	// Mise à jour dans la base de données
	return s.store.Update(user)
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

// ExtractRoleNames récupère et retourne une liste des noms de rôle pour un utilisateur en fonction de son userID
func (s *UserManagementService) ExtractRoleNames(userID uint) ([]string, error) {
	user, err := s.store.PreloadRoles(userID)
	if err != nil {
		return nil, errors.New("erreur lors de la récupération des rôles de l'utilisateur: " + err.Error())
	}

	// Extraire les noms des rôles
	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Name
	}

	return roleNames, nil
}
