package auth

import (
    "backend/internal/db/models"
    "backend/internal/utils"
    "errors"
    "gorm.io/gorm"
)

type AuthService struct {
    DB *gorm.DB
}

// NewAuthService crée une nouvelle instance de AuthService
func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{DB: db}
}

// RegisterUser gère l'inscription d'un utilisateur standard
func (s *AuthService) RegisterUser(user *models.User) error {
    // Vérifier si le pseudo est déjà pris
    if err := s.checkIfPseudoExists(user.Pseudo); err != nil {
        return err
    }

    // Hash le mot de passe
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    // Enregistre l'utilisateur dans la base de données
    return s.DB.Create(user).Error
}

// RegisterBusinessUser gère l'inscription d'un utilisateur entreprise
func (s *AuthService) RegisterBusinessUser(businessUser *models.BusinessUser) error {
    // Vérifier si le pseudo est déjà pris
    if err := s.checkIfPseudoExists(businessUser.User.Pseudo); err != nil {
        return err
    }

    // Hash le mot de passe
    hashedPassword, err := utils.HashPassword(businessUser.User.Password)
    if err != nil {
        return err
    }
    businessUser.User.Password = hashedPassword

    // Enregistre d'abord l'utilisateur
    if err := s.DB.Create(&businessUser.User).Error; err != nil {
        return err
    }

    // Associe l'ID de l'utilisateur à BusinessUser
    businessUser.UserID = businessUser.User.ID

    // Enregistre ensuite le BusinessUser
    return s.DB.Create(businessUser).Error
}

// checkIfPseudoExists vérifie si le pseudo existe déjà
func (s *AuthService) checkIfPseudoExists(pseudo string) error {
    var existingUser models.User
    if err := s.DB.Where("pseudo = ?", pseudo).First(&existingUser).Error; err == nil {
        return errors.New("pseudo already taken")
    }
    return nil
}

// CheckUserRole vérifie si l'utilisateur a un rôle spécifique
func (s *AuthService) CheckUserRole(userID uint, roleName string) (bool, error) {
    var role models.Role // Assurez-vous d'importer le bon modèle
    if err := s.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
        return false, err
    }

    var userRole models.UserRole
    if err := s.DB.Where("user_id = ? AND role_id = ?", userID, role.ID).First(&userRole).Error; err != nil {
        return false, nil // Pas de correspondance, donc l'utilisateur n'a pas ce rôle
    }

    return true, nil // L'utilisateur a ce rôle
}

// Login gère la connexion d'un utilisateur
func (s *AuthService) Login(email, password string) (string, error) {
    var user models.User

    // Vérifie si l'utilisateur existe
    if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return "", err // Retourne l'erreur si l'utilisateur n'est pas trouvé
    }

    // Vérifie le mot de passe
    if err := utils.VerifyPassword(user.Password, password); err != nil {
        return "", err // Retourne l'erreur si le mot de passe ne correspond pas
    }

    // Génère un JWT pour l'utilisateur avec son ID
    token, err := utils.GenerateJWT(user.ID, "your-secret-key") // Remplacez par votre clé secrète
    if err != nil {
        return "", err // Retourne l'erreur si la génération du token échoue
    }

    return token, nil
}
