package authentification

import (
	"backend/config"
	"backend/internal/api/models/auth/request"
	"backend/internal/api/models/auth/response"
	"backend/internal/services/auth"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	authService *auth.AuthService
}

// NewAuthHandler crée une nouvelle instance de AuthHandler
func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Initialisation de la requête de connexion
	var loginReq request.LoginRequest

	// Décode les données JSON reçues dans LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Appelle le service d'authentification pour se connecter
	loginResp, err := h.authService.Login(&loginReq)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}

	// Récupérer l'ID utilisateur à partir du token JWT pour vérification des rôles
	claims, err := utils.ValidateJWT(loginResp.Token, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID // Récupérer l'ID utilisateur à partir des claims

	// Log pour le débogage de l’authentification
	log.Printf("User ID %d attempting to login with token: %s", userID, loginResp.Token)

	// Préparation de la requête CheckUserRoleRequest
	roleCheckReq := &request.CheckUserRoleRequest{
		UserID: userID,
		Roles:  []string{"admin", "business", "association", "school"},
	}

	// Vérifier si l'utilisateur a le rôle requis
	roleCheckResp, err := h.authService.CheckUserRole(roleCheckReq)
	if err != nil {
		http.Error(w, "Error checking user roles", http.StatusInternalServerError)
		return
	}

	// Log pour le débogage de la vérification de rôle
	log.Printf("Checking roles for user ID %d: %v. Has role: %v", userID, roleCheckReq.Roles, roleCheckResp.HasRole)

	// Si l'utilisateur n'a pas le rôle requis, retournez un message d'attente de validation
	if !roleCheckResp.HasRole {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Votre compte est en attente de validation de notre part",
		})
		return
	}

	// Réponse complète avec les détails d'authentification
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResp)
}

// HandleGetUserIDByEmail gère la récupération de l'ID utilisateur par email
func (h *AuthHandler) HandleGetUserIDByEmail(w http.ResponseWriter, r *http.Request) {
	var req request.GetUserIDByEmailRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userIDResponse, err := h.authService.GetUserIDByEmail(req.Email)
	if err != nil {
		res := response.GetUserIDByEmailResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Si l'utilisateur n'a pas été trouvé
	if !userIDResponse.Success {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(userIDResponse)
		return
	}

	// Si l'utilisateur a été trouvé
	res := response.GetUserIDByEmailResponse{
		UserID:  userIDResponse.UserID,
		Success: true,
		Message: "User ID retrieved successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *AuthHandler) HandleLoginNormalUser(w http.ResponseWriter, r *http.Request) {
	// Initialisation de la requête de connexion
	var loginReq request.LoginRequest

	// Décode les données JSON reçues dans LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Appelle le service d'authentification pour se connecter
	loginResp, err := h.authService.Login(&loginReq)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}

	// Récupérer l'ID utilisateur à partir du token JWT pour confirmer l'authentification
	claims, err := utils.ValidateJWT(loginResp.Token, config.AppConfig.JwtSecretAccessKey)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID // Récupérer l'ID utilisateur à partir des claims

	// Log pour le débogage de l'authentification
	log.Printf("Normal User ID %d successfully logged in with token: %s", userID, loginResp.Token)

	// Réponse complète avec les détails d'authentification
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResp)
}
