package auth

import (
	"backend/config"
	"backend/core/models"
	"backend/core/services/user"
	storesBusiness "backend/core/stores/business"
	storesUser "backend/core/stores/user"
	"backend/core/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthService struct {
	userService   *user.UserService
	businessStore *storesBusiness.BusinessUserStore
	roleStore     *storesUser.RoleStore
	db            *gorm.DB
}

// NewAuthService crée une nouvelle instance de AuthService avec toutes les dépendances
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userService:   user.NewUserService(db),
		businessStore: storesBusiness.NewBusinessUserStore(db),
		roleStore:     storesUser.NewRoleStore(db),
		db:            db,
	}
}

// RegisterUser gère l'inscription d'un utilisateur standard
func (s *AuthService) RegisterUser(user *models.User) error {
	if err := s.checkIfEmailExists(user.Email); err != nil {
		return fmt.Errorf("email déjà utilisé : %w", err)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	user.Password = hashedPassword

	// Enregistrement de l'utilisateur
	if err := s.userService.Management.CreateUser(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	// Optionnel : attribuer un rôle par défaut
	roleID, err := s.GetRoleIDByName("User")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle : %w", err)
	}
	return s.userService.Management.AssignUserRole(user.ID, roleID) // Utilisation du service pour assigner le rôle
}

// RegisterBusinessUser gère l'inscription d'un utilisateur entreprise avec gestion des transactions
func (s *AuthService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.checkIfEmailExists(businessUser.User.Email); err != nil {
			return fmt.Errorf("email déjà utilisé : %w", err)
		}

		hashedPassword, err := utils.HashPassword(businessUser.User.Password)
		if err != nil {
			return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
		}
		businessUser.User.Password = hashedPassword

		// Enregistrement de l'utilisateur
		if err := tx.Create(&businessUser.User).Error; err != nil {
			return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
		}

		businessUser.UserID = businessUser.User.ID

		// Enregistrer le BusinessUser
		if err := tx.Create(businessUser).Error; err != nil {
			return fmt.Errorf("erreur lors de la création du BusinessUser : %w", err)
		}

		// Attribuer le rôle Business à l'utilisateur
		businessRoleID, err := s.GetRoleIDByName("Business")
		if err != nil {
			return fmt.Errorf("erreur lors de la récupération du rôle 'business' : %w", err)
		}
		return s.userService.Management.AssignUserRole(businessUser.User.ID, businessRoleID)
	})
}

// Login gère la connexion d'un utilisateur et retourne un token JWT si valide
func (s *AuthService) Login(email, password string) (string, error) {
	// Vérification de l'utilisateur
	user, err := s.userService.Retrieval.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	// Vérification du mot de passe
	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	// Génération du token JWT avec expiration de 72 heures
	token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey, "StudiMove", "studi_users", 72)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la génération du token : %w", err)
	}

	return token, nil
}

// checkIfEmailExists vérifie si l'email existe déjà
func (s *AuthService) checkIfEmailExists(email string) error {
	_, err := s.userService.Retrieval.GetUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Si l'email n'existe pas
		return nil
	}
	// Si l'email existe déjà
	if err == nil {
		return errors.New("email déjà utilisé")
	}
	return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
}

// GetRoleIDByName récupère l'ID d'un rôle par son nom.
func (s *AuthService) GetRoleIDByName(roleName string) (uint, error) {
	role, err := s.roleStore.GetByName(roleName)
	if err != nil {
		return 0, fmt.Errorf("erreur lors de la récupération du rôle %s : %w", roleName, err)
	}
	return role.ID, nil
}

// CheckUserRole vérifie si l'utilisateur a le rôle spécifié
func (s *AuthService) CheckUserRole(userID uint, role string) (bool, error) {
	var user models.User
	if err := s.db.Preload("Roles").First(&user, userID).Error; err != nil {
		return false, fmt.Errorf("erreur lors de la récupération de l'utilisateur : %w", err)
	}

	for _, userRole := range user.Roles {
		if userRole.Name == role {
			return true, nil
		}
	}
	return false, nil
}

// ExtractRoleNames retourne une liste des noms de rôle d'un utilisateur
func (s *AuthService) ExtractRoleNames(roles []models.Role) []string {
	roleNames := make([]string, len(roles))
	for i, role := range roles {
		roleNames[i] = role.Name
	}
	return roleNames
}
