package profil

import (
	"backend/internal/api/models/profil/request"
	"backend/internal/api/models/profil/response"
	"backend/internal/services/profilservice"
	"backend/internal/services/userservice"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

type UpdateProfilHandler struct {
	profilService *profilservice.ProfilService
	userService   *userservice.UserService
}

func NewUpdateProfilHandler(profilService *profilservice.ProfilService, userService *userservice.UserService) *UpdateProfilHandler {
	return &UpdateProfilHandler{profilService: profilService, userService: userService}
}

// HandleUpdateOwnProfile gère la mise à jour du profil de l'utilisateur connecté sans besoin de spécifier un targetId
func (h *UpdateProfilHandler) HandleUpdateOwnProfile(w http.ResponseWriter, r *http.Request) {
	// Étape 1: Récupérer les revendications utilisateur du contexte
	log.Println("Début de HandleUpdateOwnProfile")

	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		log.Println("Erreur: utilisateur non authentifié ou claims JWT non trouvés")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userClaims.UserID
	log.Printf("UserID from context: %d", userID)

	// Étape 2: Décoder les données de mise à jour du corps de la requête
	var updateData request.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.Printf("Erreur lors du décodage de la requête: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("Données de mise à jour reçues: %+v", updateData)

	// Étape 3: Récupérer les rôles de l'utilisateur authentifié
	log.Printf("Récupération des rôles pour l'utilisateur ID: %d", userID)
	roles, err := h.userService.GetUserRolesByID(userID)
	if err != nil {
		log.Printf("Erreur lors de la récupération des rôles utilisateur: %v", err)
		http.Error(w, "Failed to retrieve user roles", http.StatusInternalServerError)
		return
	}
	log.Printf("Rôles récupérés pour l'utilisateur: %v", roles)

	// Étape 4: Mettre à jour le profil utilisateur
	log.Printf("Mise à jour du profil de l'utilisateur ID: %d avec les données: %+v", userID, updateData)
	if err := h.profilService.UpdateUserProfile(userID, updateData, roles); err != nil {
		log.Printf("Erreur lors de la mise à jour du profil utilisateur: %v", err)
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}
	log.Println("Profil utilisateur mis à jour avec succès")

	// Étape 5: Répondre avec un message de succès
	responseData := response.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated successfully",
	}
	log.Println("Envoi de la réponse de succès pour la mise à jour du profil")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}


// HandleUpdateTargetProfile gère la mise à jour du profil d'un utilisateur spécifique
func (h *UpdateProfilHandler) HandleUpdateTargetProfile(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		TargetId   uint                        	`json:"targetId"` 
		UpdateData   request.UpdateProfileRequest `json:"updateData"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.profilService.UpdateUserProfileByTargetID(requestData.TargetId, requestData.UpdateData); err != nil {
		log.Printf("Error updating profile: %v", err)
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	responseData := response.UpdateProfileResponse{
		Success: true,
		Message: "Profile updated successfully",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
