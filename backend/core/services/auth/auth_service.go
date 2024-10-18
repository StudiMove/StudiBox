// package auth

// import (
//     "backend/internal/db/models"
//     "backend/internal/utils" // Assure-toi d'importer le bon package
//     "backend/config" // Importer le package de configuration

//     "errors"
//     "gorm.io/gorm"
// )

// type AuthService struct {
//     DB *gorm.DB
// }

// // NewAuthService crée une nouvelle instance de AuthService
// func NewAuthService(db *gorm.DB) *AuthService {
//     return &AuthService{DB: db}
// }

// // RegisterUser gère l'inscription d'un utilisateur standard
// func (s *AuthService) RegisterUser(user *models.User) error {
//     // Vérifier si l'email est déjà pris
//     if err := s.checkIfEmailExists(user.Email); err != nil {
//         return err
//     }

//     // Hash le mot de passe
//     hashedPassword, err := utils.HashPassword(user.Password)
//     if err != nil {
//         return err
//     }
//     user.Password = hashedPassword

//     // Enregistre l'utilisateur dans la base de données
//     return s.DB.Create(user).Error
// }

// // RegisterBusinessUser gère l'inscription d'un utilisateur entreprise
// func (s *AuthService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
//     // Vérifier si le pseudo est déjà pris
//     if err := s.checkIfPseudoExists(businessUser.User.Pseudo); err != nil {
//         return err
//     }

//     // Hash le mot de passe
//     hashedPassword, err := utils.HashPassword(businessUser.User.Password)
//     if err != nil {
//         return err
//     }
//     businessUser.User.Password = hashedPassword

//     // Enregistre d'abord l'utilisateur
//     if err := s.DB.Create(&businessUser.User).Error; err != nil {
//         return err
//     }

//     // Associe l'ID de l'utilisateur à BusinessUser
//     businessUser.UserID = businessUser.User.ID

//     // Enregistre ensuite le BusinessUser
//     return s.DB.Create(businessUser).Error
// }

// // checkIfEmailExists vérifie si l'email existe déjà
// func (s *AuthService) checkIfEmailExists(email string) error {
//     var existingUser models.User
//     if err := s.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
//         return errors.New("email already taken")
//     }
//     return nil
// }

// // checkIfPseudoExists vérifie si le pseudo existe déjà
// func (s *AuthService) checkIfPseudoExists(pseudo string) error {
//     var existingUser models.User
//     if err := s.DB.Where("pseudo = ?", pseudo).First(&existingUser).Error; err == nil {
//         return errors.New("pseudo already taken")
//     }
//     return nil
// }
// // CheckUserRole vérifie si l'utilisateur a le rôle spécifié
// func (s *AuthService) CheckUserRole(userID uint, role string) (bool, error) {
//     // Logique pour vérifier le rôle de l'utilisateur
//     var user models.User
//     if err := s.DB.First(&user, userID).Error; err != nil {
//         return false, err // Retourne faux si l'utilisateur n'est pas trouvé
//     }

//     // Assure-toi d'avoir un champ dans ton modèle qui stocke les rôles
//     for _, r := range user.Roles {
//         if r.Name == role {
//             return true, nil // Retourne vrai si l'utilisateur a le rôle
//         }
//     }
//     return false, nil // Retourne faux si l'utilisateur n'a pas le rôle
// }

// // Login gère la connexion d'un utilisateur
// func (s *AuthService) Login(email, password string) (string, error) {
//     var user models.User

//     // Vérifie si l'utilisateur existe
//     if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
//         return "", err // Retourne l'erreur si l'utilisateur n'est pas trouvé
//     }

//     // Vérifie le mot de passe
//     if err := utils.VerifyPassword(user.Password, password); err != nil {
//         return "", err // Retourne l'erreur si le mot de passe ne correspond pas
//     }

//     // Génère un JWT pour l'utilisateur avec son ID
//     token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey) // Remplace par ta clé secrète
//     if err != nil {
//         return "", err // Retourne l'erreur si la génération du token échoue
//     }

//     return token, nil
// }
// // GetRoleIDByName récupère l'ID d'un rôle par son nom.
// func (s *AuthService) GetRoleIDByName(roleName string) (uint, error) {
//     var role models.Role
//     if err := s.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
//         return 0, err // Retourne une erreur si le rôle n'est pas trouvé
//     }
//     return role.ID, nil
// }
// // AssignUserRole associe un utilisateur à un rôle donné.
// func (s *AuthService) AssignUserRole(userID uint, roleID uint) error {
//     // Créez une entrée dans la table de liaison user_role
//     userRole := models.UserRole{
//         UserID: userID,
//         RoleID: roleID,
//     }
//     return s.DB.Create(&userRole).Error
// }

// /*
// Gestion des images :
// - Ici, tu pourrais ajouter la logique pour gérer le téléchargement et l'association des images aux utilisateurs.
// - Cela pourrait impliquer :
//   - La création d'un modèle pour stocker les informations d'image (chemin, type, etc.).
//   - Une méthode pour gérer le téléchargement d'image.
//   - L'association de l'image à l'utilisateur ou à l'entreprise lors de l'inscription.
//   - Le code ci-dessous est un exemple fictif.

// func (s *AuthService) UploadUserImage(userID uint, imagePath string) error {
//     // Logique pour associer l'image à l'utilisateur
//     // Par exemple, en mettant à jour le champ ImagePath de l'utilisateur dans la base de données
//     var user models.User
//     if err := s.DB.First(&user, userID).Error; err != nil {
//         return err // Retourne l'erreur si l'utilisateur n'est pas trouvé
//     }

//     user.ImagePath = imagePath // Mise à jour du chemin de l'image
//     return s.DB.Save(&user).Error
// }
// */

package auth

import (
	"backend/config"
	"backend/core/models"
	"backend/core/stores"
	"backend/core/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthService struct {
	userStore     *stores.UserStore
	businessStore *stores.BusinessUserStore
	roleStore     *stores.RoleStore
	userRoleStore *stores.UserRoleStore
	db            *gorm.DB
}

// NewAuthService crée une nouvelle instance de AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userStore:     stores.NewUserStore(db),
		businessStore: stores.NewBusinessUserStore(db),
		roleStore:     stores.NewRoleStore(db),
		userRoleStore: stores.NewUserRoleStore(db),
		db:            db,
	}
}

// RegisterUser gère l'inscription d'un utilisateur standard
func (s *AuthService) RegisterUser(user *models.User) error {
	// Vérifier si l'email est déjà pris
	if err := s.checkIfEmailExists(user.Email); err != nil {
		return fmt.Errorf("email déjà utilisé : %w", err)
	}

	// Hasher le mot de passe
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
	}
	user.Password = hashedPassword

	// Enregistre l'utilisateur dans la base de données
	if err := s.userStore.Create(user); err != nil {
		return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
	}

	return nil
}

// RegisterBusinessUser gère l'inscription d'un utilisateur entreprise avec gestion des transactions
func (s *AuthService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// Vérifier si l'email est déjà pris
		if err := s.checkIfEmailExists(businessUser.User.Email); err != nil {
			return fmt.Errorf("email déjà utilisé : %w", err)
		}

		// Hasher le mot de passe
		hashedPassword, err := utils.HashPassword(businessUser.User.Password)
		if err != nil {
			return fmt.Errorf("erreur lors du hachage du mot de passe : %w", err)
		}
		businessUser.User.Password = hashedPassword

		// Enregistrer l'utilisateur
		if err := tx.Create(&businessUser.User).Error; err != nil {
			return fmt.Errorf("erreur lors de la création de l'utilisateur : %w", err)
		}

		// Associer l'utilisateur à BusinessUser
		businessUser.UserID = businessUser.User.ID

		// Enregistrer BusinessUser
		if err := tx.Create(businessUser).Error; err != nil {
			return fmt.Errorf("erreur lors de la création du BusinessUser : %w", err)
		}

		// Attribuer le rôle Business à l'utilisateur
		businessRoleID, err := s.GetRoleIDByName("Business")
		if err != nil {
			return fmt.Errorf("erreur lors de la récupération du rôle 'business' : %w", err)
		}
		if err := s.AssignUserRole(businessUser.User.ID, businessRoleID); err != nil {
			return fmt.Errorf("erreur lors de l'attribution du rôle à l'utilisateur : %w", err)
		}

		return nil
	})
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

// Login gère la connexion d'un utilisateur et retourne un token JWT si valide
func (s *AuthService) Login(email, password string) (string, error) {
	// Vérifie si l'utilisateur existe
	user, err := s.userStore.GetByEmail(email)
	if err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	// Vérifie le mot de passe
	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	// Génère un JWT pour l'utilisateur avec une expiration de 72 heures
	token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey, 72)
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

// AssignUserRole associe un utilisateur à un rôle donné.
func (s *AuthService) AssignUserRole(userID uint, roleID uint) error {
	userRole := models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	if err := s.userRoleStore.Create(&userRole); err != nil {
		return fmt.Errorf("erreur lors de l'attribution du rôle à l'utilisateur : %w", err)
	}
	return nil
}

// checkIfEmailExists vérifie si l'email existe déjà
func (s *AuthService) checkIfEmailExists(email string) error {
	_, err := s.userStore.GetByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Si aucune correspondance n'est trouvée, l'email n'est pas encore utilisé
		return nil
	}
	// Si une autre erreur survient ou si l'email est trouvé
	if err == nil {
		return errors.New("email déjà utilisé")
	}
	return fmt.Errorf("erreur lors de la vérification de l'email : %w", err)
}
