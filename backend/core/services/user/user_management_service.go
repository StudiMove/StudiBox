package user

import (
	request "backend/core/api/request/user"
	stores "backend/core/stores/user"
	"backend/database/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// UserManagementServiceType gère les opérations de gestion des utilisateurs
type UserManagementServiceType struct {
	store *stores.UserStoreType
}

// NewUserManagementServiceType crée une nouvelle instance de UserManagementServiceType
func UserManagementService(store *stores.UserStoreType) *UserManagementServiceType {
	return &UserManagementServiceType{
		store: store,
	}
}

// CreateUser crée un nouvel utilisateur
func (s *UserManagementServiceType) CreateUser(user *models.User) error {
	if err := s.store.Create(user); err != nil {
		return errors.New("failed to create user: " + err.Error())
	}
	return nil
}

// UpdateUser met à jour un utilisateur existant
func (s *UserManagementServiceType) UpdateUser(user *models.User) error {
	if err := s.store.Update(user); err != nil {
		return errors.New("failed to update user: " + err.Error())
	}
	return nil
}

// DeleteUser supprime un utilisateur par ID
func (s *UserManagementServiceType) DeleteUser(userID uint) error {
	if err := s.store.Delete(userID); err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}
	return nil
}

// UpdateUserProfile met à jour uniquement les champs spécifiés dans la requête
func (s *UserManagementServiceType) UpdateUserProfile(userID uint, input request.UserRequest) error {
	// Récupérer l'utilisateur existant par son ID
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("utilisateur introuvable")
		}
		return fmt.Errorf("erreur lors de la récupération de l'utilisateur : %w", err)
	}

	// Vérifier si l'utilisateur est nil
	if user == nil {
		return errors.New("utilisateur introuvable")
	}

	// Vérifier si l'email est modifié et s'il est déjà utilisé par un autre utilisateur
	if input.Email != "" && input.Email != user.Email {
		existingUser, err := s.store.GetByEmail(input.Email)
		if err == nil && existingUser != nil && existingUser.ID != userID {
			return errors.New("cet email est déjà utilisé par un autre utilisateur")
		}
	}

	// Créer un map pour stocker les champs à mettre à jour
	updates := map[string]interface{}{
		"first_name":    input.FirstName,
		"last_name":     input.LastName,
		"email":         input.Email,
		"pseudo":        input.Pseudo,
		"profile_image": input.ProfileImage,
		"birth_date":    input.BirthDate,
		"country":       input.Country,
		"city":          input.City,
		"address":       input.Address,
		"postal_code":   input.PostalCode,
		"profile_type":  input.ProfileType,
	}

	// Vérifier si le numéro de téléphone est valide avant de l'ajouter
	if input.Phone != 0 {
		updates["phone"] = input.Phone
	}

	// Supprimer les champs vides pour ne pas les mettre à jour
	for key, value := range updates {
		if value == "" || value == nil {
			delete(updates, key)
		}
	}

	// Si aucun champ à mettre à jour, retourner sans rien faire
	if len(updates) == 0 {
		return nil
	}

	// Mettre à jour les champs dans la base de données
	return s.store.UpdateFields(userID, updates)
}

// AssignUserRole assigne un rôle à un utilisateur en mettant à jour le RoleID dans User
func (s *UserManagementServiceType) AssignUserRole(userID uint, roleID uint) error {
	// Récupère l'utilisateur par ID
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return errors.New("failed to retrieve user: " + err.Error())
	}

	// Mise à jour du RoleID de l'utilisateur
	user.RoleID = roleID

	// Enregistrer la mise à jour de l'utilisateur
	if err := s.store.Update(user); err != nil {
		return errors.New("failed to assign role: " + err.Error())
	}
	return nil
}

// ExtractRoleName récupère le nom du rôle pour un utilisateur en utilisant le store
func (s *UserManagementServiceType) ExtractRoleName(userID uint) (string, error) {
	// Précharge le rôle de l'utilisateur
	user, err := s.store.PreloadRole(userID)
	if err != nil {
		return "", errors.New("erreur lors de la récupération du rôle de l'utilisateur: " + err.Error())
	}

	// Retourne le nom du rôle
	return user.Role.Name, nil
}

// CheckUserRole vérifie si l'utilisateur a le rôle spécifié
func (s *UserManagementServiceType) CheckUserRole(userID uint, roleName string) (bool, error) {
	user, err := s.store.PreloadRole(userID)
	if err != nil {
		return false, fmt.Errorf("erreur lors de la récupération de l'utilisateur : %w", err)
	}

	// Vérifier si le nom du rôle de l'utilisateur correspond à celui attendu
	if user.Role.Name == roleName {
		return true, nil
	}
	return false, nil
}
