package profil

import (
	"backend/internal/api/models/profil/request"
	"backend/internal/api/models/profil/response"
	"backend/internal/services/profilservice"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ProfilHandler gère les requêtes pour le profil utilisateur
type ProfilHandler struct {
	profilService *profilservice.ProfilService
}

// NewProfilHandler crée une nouvelle instance de ProfilHandler
func NewProfilHandler(profilService *profilservice.ProfilService) *ProfilHandler {
	return &ProfilHandler{profilService: profilService}
}

// GetUserProfile gère la récupération des informations de profil via le token JWT
func (h *ProfilHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetUserProfile handler")

	// Récupère les claims utilisateur à partir du contexte avec la clé "user"
	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		log.Println("Unauthorized: User claims not found in context")
		http.Error(w, "Unauthorized: No valid user claims found in context", http.StatusUnauthorized)
		return
	}
	userID := userClaims.UserID
	log.Printf("UserID from context: %d", userID)

	// Récupère le profil de l'utilisateur via le service
	userProfile, err := h.profilService.GetUserProfileByTargetID(userID)
	if err != nil {
		log.Println("Failed to retrieve user profile:", err)
		http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
		return
	}

	// Prépare la réponse
	responseData := response.GetProfileResponse{
		UserID:       userProfile.UserID,
		Email:        userProfile.Email,
		Phone:        userProfile.Phone,
		ProfileImage: userProfile.ProfileImage,
		RoleNames:    userProfile.RoleNames,
		Organisation: userProfile.Organisation,
	}

	log.Println("User profile retrieved successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}

// GetUserProfileByTargetID gère la récupération des informations de profil via un targetId passé dans le corps de la requête
func (h *ProfilHandler) GetUserProfileByTargetID(w http.ResponseWriter, r *http.Request) {
	var requestData request.GetProfileByTargetIDRequest

	// Décoder le corps de la requête pour obtenir targetId
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Récupérer le profil de l'utilisateur via targetId
	userProfile, err := h.profilService.GetUserProfileByTargetID(requestData.TargetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Prépare la réponse en utilisant la structure GetProfileResponse
	responseData := response.GetProfileResponse{
		UserID:       userProfile.UserID,
		Email:        userProfile.Email,
		Phone:        userProfile.Phone,
		ProfileImage: userProfile.ProfileImage,
		RoleNames:    userProfile.RoleNames,
		Organisation: userProfile.Organisation,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
