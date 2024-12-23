// backend/internal/api/handlers/user/get_user_role_handler.go
package user

import (
	"backend/config"
	"backend/internal/api/models/profil/request"
	"backend/internal/services/userservice" // Mise à jour du nom du package
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// UserHandler gère les routes liées aux utilisateurs.
type UserHandler struct {
	userService *userservice.UserService
}

// NewUserHandler initialise un nouveau UserHandler.
func NewUserHandler(userService *userservice.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserRoleHandler gère la récupération du rôle d'un utilisateur.
func GetUserRoleHandler(userService *userservice.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetUserRoleHandler reached")

		// Récupérer le token à partir des en-têtes
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		// Vérifie que le token commence par "Bearer "
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			http.Error(w, "Unauthorized: invalid token format", http.StatusUnauthorized)
			return
		}

		// Extraire le token sans le préfixe "Bearer "
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		// Extraire l'ID de l'utilisateur à partir du token
		userID, err := utils.ExtractUserIDFromToken(tokenStr, config.AppConfig.JwtSecretAccessKey)
		if err != nil {
			log.Printf("Erreur d'extraction du token: %v", err)
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// Utiliser le UserService pour obtenir les rôles de l'utilisateur
		roles, err := userService.GetUserRolesByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Extraire uniquement les noms des rôles
		var roleNames []string
		for _, role := range roles {
			roleNames = append(roleNames, role.Name)
		}

		// Renvoyer seulement le nom du rôle en réponse
		response := map[string]interface{}{
			"role": roleNames[0], // Si vous voulez juste le premier rôle (si un utilisateur a plusieurs rôles)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Erreur lors de l'encodage de la réponse: %v", err)
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		}
	}
}

// GetStudiboxCoins récupère le solde des Studibox Coins d'un utilisateur.
func (h *UserHandler) GetStudiboxCoins(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID utilisateur depuis les paramètres de l'URL
	vars := mux.Vars(r)
	userIDParam, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convertit l'ID en entier
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Appelle le service pour récupérer les Studibox Coins
	coins, err := h.userService.GetUserStudiboxCoinsByID(uint(userID))
	if err != nil {
		http.Error(w, "Error fetching Studibox Coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode et retourne les Studibox Coins en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"studiboxCoins": coins,
	})
}

func (h *UserHandler) UpdateCoinsForAuthenticatedUser(w http.ResponseWriter, r *http.Request) {
	// Récupérer les claims utilisateur depuis le contexte
	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		http.Error(w, "Unauthorized request: invalid or missing JWT claims", http.StatusUnauthorized)
		return
	}

	// Récupérer l'ID utilisateur depuis les claims
	userID := int(userClaims.UserID) // Conversion explicite de uint à int

	// Récupérer le montant des coins depuis le body
	var payload struct {
		Coins int `json:"coins"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Appeler le service pour déduire les coins
	if err := h.userService.UpdateUserCoinsByID(userID, payload.Coins); err != nil {
		if err.Error() == "insufficient coins" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to update coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Coins deducted successfully"}`))
}

func (h *UserHandler) UpdateCoinsByID(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID utilisateur depuis les paramètres de l'URL
	vars := mux.Vars(r)
	userIDParam, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Récupérer les coins depuis le body
	var payload struct {
		Coins int `json:"coins"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Appeler le service pour mettre à jour les coins
	if err := h.userService.UpdateUserCoinsByID(userID, payload.Coins); err != nil {
		http.Error(w, "Failed to update coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Coins updated successfully"}`))
}
func (h *UserHandler) AddCoinsByID(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID utilisateur depuis les paramètres de l'URL
	vars := mux.Vars(r)
	userIDParam, exists := vars["id"]
	if !exists {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convertir l'ID en entier
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Récupérer le montant des coins depuis le body
	var payload struct {
		Coins int `json:"coins"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Appeler le service pour ajouter les coins
	if err := h.userService.AddUserCoinsByID(userID, payload.Coins); err != nil {
		http.Error(w, "Failed to add coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Coins added successfully"}`))
}

func (h *UserHandler) GetUserIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire l'email depuis les paramètres de la requête ou le corps
	var payload struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Appeler le service pour récupérer l'ID de l'utilisateur
	userID, err := h.userService.GetUserIDByEmail(payload.Email) // Notez l'utilisation correcte de h.userService
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // Retourner une erreur si l'utilisateur n'est pas trouvé
		return
	}

	// Répondre avec l'ID de l'utilisateur
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
	})
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Décoder le corps de la requête JSON directement dans `request.UpdateUserRequest`.
	var payload request.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Récupérer l'ID utilisateur depuis les claims (JWT).
	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		http.Error(w, "Unauthorized request: invalid or missing JWT claims", http.StatusUnauthorized)
		return
	}

	userID := userClaims.UserID

	// Appeler le service pour mettre à jour les données utilisateur.
	if err := h.userService.UpdateUser(userID, payload); err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			http.Error(w, "Invalid old password", http.StatusBadRequest)
			return
		}
		log.Printf("Failed to update user: %v", err)
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Réponse en cas de succès.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User updated successfully"}`))
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire l'ID utilisateur depuis les claims JWT
	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		http.Error(w, "Unauthorized request: invalid or missing JWT claims", http.StatusUnauthorized)
		return
	}

	userID := userClaims.UserID

	// Appeler le service pour supprimer l'utilisateur
	if err := h.userService.DeleteUserByID(userID); err != nil {
		if strings.Contains(err.Error(), "user not found") {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Réponse en cas de succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User deleted successfully"}`))
}
