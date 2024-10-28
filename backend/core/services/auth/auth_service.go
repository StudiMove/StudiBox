package auth

import (
	"backend/config"
	"backend/core/models"
	"backend/core/services/user"
	storesBusiness "backend/core/stores/business"
	storesUser "backend/core/stores/user"
	"backend/core/utils"
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
	// Vérifier si l'email existe déjà
	existingUser, err := s.userService.Retrieval.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		// Si l'utilisateur existe déjà, retourner une erreur
		return fmt.Errorf("email déjà utilisé")
	}

	// Hachage du mot de passe
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	user.Password = hashedPassword

	// Enregistrement de l'utilisateur
	if err := s.userService.Management.CreateUser(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	// Récupérer l'ID du rôle "User"
	roleID, err := s.GetRoleIDByName("User")
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du rôle : %w", err)
	}

	// Assigner le rôle à l'utilisateur (création d'une entrée dans user_roles)
	if err := s.userService.Management.AssignUserRole(user.ID, roleID); err != nil {
		return fmt.Errorf("erreur lors de l'attribution du rôle à l'utilisateur : %w", err)
	}

	return nil
}

// RegisterBusinessUser gère l'inscription d'un utilisateur entreprise avec gestion des transactions
func (s *AuthService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
	// Vérifier si l'email existe déjà
	existingUser, err := s.userService.Retrieval.GetUserByEmail(businessUser.User.Email)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
	}
	if existingUser != nil {
		return fmt.Errorf("email déjà utilisé")
	}

	// Hachage du mot de passe
	hashedPassword, err := utils.HashPassword(businessUser.User.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	businessUser.User.Password = hashedPassword

	// Enregistrement de l'utilisateur et du BusinessUser
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&businessUser.User).Error; err != nil {
			return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
		}

		businessUser.UserID = businessUser.User.ID

		if err := tx.Create(businessUser).Error; err != nil {
			return fmt.Errorf("erreur lors de la création du BusinessUser : %w", err)
		}

		// Récupérer l'ID du rôle "Business"
		businessRoleID, err := s.GetRoleIDByName("Business")
		if err != nil {
			return fmt.Errorf("erreur lors de la récupération du rôle 'business' : %w", err)
		}

		// Assigner le rôle Business à l'utilisateur (création d'une entrée dans user_roles)
		if err := s.userService.Management.AssignUserRole(businessUser.User.ID, businessRoleID); err != nil {
			return fmt.Errorf("erreur lors de l'attribution du rôle business : %w", err)
		}

		return nil
	})
}

// Login gère la connexion d'un utilisateur et retourne un token JWT si valide
func (s *AuthService) Login(email, password string) (string, error) {
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

func (s *AuthService) GetOrCreateUserByEmail(email, firstName, lastName string) (*models.User, error) {
	user, err := s.userService.Retrieval.GetUserByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to retrieve user by email: %w", err)
	}

	// Créer un nouvel utilisateur si non trouvé
	if user == nil {
		user = &models.User{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		}
		if err := s.userService.Management.CreateUser(user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	return user, nil
}
