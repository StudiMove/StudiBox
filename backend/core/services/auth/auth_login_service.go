package auth

import (
	"backend/config"
	"backend/core/models"
	"backend/core/services/user"
	"backend/core/utils"
	"fmt"
)

type AuthLoginService struct {
	userService *user.UserService
}

func NewAuthLoginService(userService *user.UserService) *AuthLoginService {
	return &AuthLoginService{
		userService: userService,
	}
}

// Login gère la connexion d'un utilisateur et retourne un token JWT si valide
func (s *AuthLoginService) Login(email, password string) (string, error) {
	user, err := s.userService.Retrieval.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", fmt.Errorf("email ou mot de passe invalide : %w", err)
	}

	token, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey, "StudiMove", "studi_users", 72)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la génération du token : %w", err)
	}

	return token, nil
}

// FirebaseLogin gère la connexion via Firebase ID Token et retourne un JWT pour l'utilisateur
func (s *AuthLoginService) FirebaseLogin(idToken string) (string, *models.User, error) {
	// Vérification du ID Token de Firebase
	token, err := config.VerifyIDToken(idToken)
	if err != nil {
		return "", nil, fmt.Errorf("invalid Firebase ID token : %w", err)
	}

	// Récupération des informations d'utilisateur depuis le token Firebase
	userEmail := token.Claims["email"].(string)
	firstName := token.Claims["given_name"].(string)
	lastName := token.Claims["family_name"].(string)

	// Vérifie ou crée l'utilisateur dans la base de données
	user, err := s.userService.Retrieval.GetOrCreateUserByEmail(userEmail, firstName, lastName)
	if err != nil {
		return "", nil, fmt.Errorf("failed to retrieve or create user : %w", err)
	}

	// Génération du JWT pour l'utilisateur
	jwtToken, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey, "StudiMove", "studi_users", 72)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate JWT : %w", err)
	}

	return jwtToken, user, nil
}
