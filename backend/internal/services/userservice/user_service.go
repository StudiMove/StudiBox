// commentairebackend/internal/services/userservice/user_service.go
package userservice

import (
	"backend/internal/api/models/profil/request"
	"backend/internal/db/models"
	"backend/internal/utils"
	"errors"

	"gorm.io/gorm"
)

// UserService représente le service pour gérer les utilisateurs.
type UserService struct {
	db *gorm.DB // Instance de Gorm pour interagir avec la base de données.
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUserRolesByID récupère les rôles d'un utilisateur à partir de son ID.
func (s *UserService) GetUserRolesByID(userID uint) ([]models.Role, error) {
	var user models.User

	// Utilise Preload pour charger les rôles associés à l'utilisateur.
	if err := s.db.Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, err // Retourne l'erreur si la récupération échoue.
	}

	return user.Roles, nil // Retourne les rôles de l'utilisateur.
}

// GetUserStudiboxCoinsByID récupère le solde des Studibox Coins d'un utilisateur par son ID.
func (s *UserService) GetUserStudiboxCoinsByID(userID uint) (int, error) {
	var user models.User

	// Récupère l'utilisateur par ID, mais uniquement le champ StudiboxCoins
	if err := s.db.Select("studibox_coins").First(&user, userID).Error; err != nil {
		return 0, err // Retourne 0 et l'erreur si la récupération échoue
	}

	return user.StudiboxCoins, nil // Retourne le solde des Studibox Coins
}

// UpdateUserCoinsByID met à jour le solde des Studibox Coins pour un utilisateur donné en soustrayant le montant reçu.
func (s *UserService) UpdateUserCoinsByID(userID int, coins int) error {
	// Récupère l'utilisateur pour obtenir le solde actuel
	var user models.User
	if err := s.db.Select("studibox_coins").Where("id = ?", userID).First(&user).Error; err != nil {
		return err // Retourne une erreur si l'utilisateur n'est pas trouvé
	}

	// Calculer le nouveau solde
	newCoins := user.StudiboxCoins - coins
	if newCoins < 0 {
		return errors.New("insufficient coins") // Retourne une erreur si le solde devient négatif
	}

	// Met à jour le solde dans la base de données
	if err := s.db.Model(&models.User{}).Where("id = ?", userID).Update("studibox_coins", newCoins).Error; err != nil {
		return err // Retourne une erreur si la mise à jour échoue
	}

	return nil
}

// AddUserCoinsByID met à jour le solde des Studibox Coins pour un utilisateur donné en ajoutant le montant reçu.
func (s *UserService) AddUserCoinsByID(userID int, coins int) error {
	// Récupère l'utilisateur pour obtenir le solde actuel
	var user models.User
	if err := s.db.Select("studibox_coins").Where("id = ?", userID).First(&user).Error; err != nil {
		return err // Retourne une erreur si l'utilisateur n'est pas trouvé
	}

	// Calculer le nouveau solde
	newCoins := user.StudiboxCoins + coins

	// Met à jour le solde dans la base de données
	if err := s.db.Model(&models.User{}).Where("id = ?", userID).Update("studibox_coins", newCoins).Error; err != nil {
		return err // Retourne une erreur si la mise à jour échoue
	}

	return nil
}

// GetUserIDByEmail récupère l'ID d'un utilisateur à partir de son email.
func (s *UserService) GetUserIDByEmail(email string) (uint, error) {
	var user models.User

	// Recherche de l'utilisateur par email
	if err := s.db.Select("id").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("email inconnu : cet utilisateur n'est pas inscrit")
		}
		return 0, err // Retourne une autre erreur en cas d'échec de la requête
	}

	return user.ID, nil // Retourne l'ID de l'utilisateur
}

func (s *UserService) UpdateUser(userID uint, req request.UpdateUserRequest) error {
	var user models.User

	// Vérifier si l'utilisateur existe.
	if err := s.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Mise à jour des champs facultatifs.
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	// Gestion des mots de passe.
	if req.OldPassword != nil && req.NewPassword != nil {
		if err := utils.VerifyPassword(user.Password, *req.OldPassword); err != nil {
			return errors.New("invalid old password")
		}

		// Hasher le nouveau mot de passe avant de le sauvegarder.
		hashedPassword, err := utils.HashPassword(*req.NewPassword)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	// Mettre à jour les champs dans la table `User`.
	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	// Gestion des champs de localisation dans la table `UserLocation`.
	var location models.UserLocation
	if err := s.db.Where("user_id = ?", userID).First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Si aucune localisation n'existe, en créer une nouvelle.
			location = models.UserLocation{UserID: userID}
		} else {
			return err
		}
	}

	if req.Street != nil {
		location.Street = *req.Street
	}
	if req.NumberStreet != nil {
		location.NumberStreet = *req.NumberStreet
	}
	if req.City != nil {
		location.City = *req.City
	}
	if req.Postcode != nil {
		location.Postcode = *req.Postcode
	}
	if req.Region != nil {
		location.Region = *req.Region
	}
	if req.Country != nil {
		location.Country = *req.Country
	}

	// Sauvegarder les données de localisation.
	if err := s.db.Save(&location).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUserByID supprime un utilisateur ainsi que ses références dans d'autres tables.
func (s *UserService) DeleteUserByID(userID uint) error {
	// Utilisation de la transaction pour garantir la cohérence.
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Vérifier si l'utilisateur existe.
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("user not found")
			}
			return err
		}

		// Supprimer les relations avec les rôles (table intermédiaire user_roles).
		if err := tx.Model(&user).Association("Roles").Clear(); err != nil {
			return err
		}

		// Supprimer les données de localisation (UserLocation).
		if err := tx.Where("user_id = ?", userID).Delete(&models.UserLocation{}).Error; err != nil {
			return err
		}

		// Supprimer l'utilisateur lui-même.
		if err := tx.Delete(&user).Error; err != nil {
			return err
		}

		return nil
	})
}
