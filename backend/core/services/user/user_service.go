package user

import (
	"backend/core/models"
	"backend/core/stores"
	"backend/core/utils"
	"errors"

	"gorm.io/gorm"
)

// UserService représente le service pour gérer les utilisateurs.
type UserService struct {
	userStore *stores.UserStore
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(userStore *stores.UserStore) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

// CreateUser crée un nouvel utilisateur.
func (s *UserService) CreateUser(user *models.User) error {
	// Validation ou logique supplémentaire ici si nécessaire
	return s.userStore.Create(user)
}

// UpdateUser met à jour un utilisateur existant.
func (s *UserService) UpdateUser(user *models.User) error {
	return s.userStore.Update(user)
}

// DeleteUser supprime un utilisateur par ID.
func (s *UserService) DeleteUser(userID uint) error {
	return s.userStore.Delete(userID)
}

// GetUserByID récupère un utilisateur par ID.
func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.userStore.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUserByEmail récupère un utilisateur par email.
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userStore.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

// GetUserRolesByID récupère les rôles d'un utilisateur par son ID.
func (s *UserService) GetUserRolesByID(userID uint) ([]models.Role, error) {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user.Roles, nil
}

// AssignUserRole assigne un rôle à un utilisateur.
func (s *UserService) AssignUserRole(userID uint, roleID uint) error {
	user, err := s.GetUserByID(userID)
	if err != nil {
		return err
	}

	role := models.UserRole{
		UserID: user.ID,
		RoleID: roleID,
	}

	// Ajoutez une méthode dans UserStore pour gérer l'association des rôles
	return s.userStore.AssignRole(&role)
}

// AuthenticateUser vérifie si un email et un mot de passe correspondent à un utilisateur.
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Vérifier le mot de passe (logique de comparaison de hash)
	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
