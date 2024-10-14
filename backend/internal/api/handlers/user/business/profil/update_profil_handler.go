package profil

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/userservice/business/profilservice" // Corriger l'import
    "backend/internal/utils"
)

type UpdateProfilHandler struct {
    profilService *profilservice.ProfilService // Utilise le bon service
}

func NewUpdateProfilHandler(profilService *profilservice.ProfilService) *UpdateProfilHandler {
    return &UpdateProfilHandler{profilService: profilService}
}

// HandleUpdateProfil gère la mise à jour des informations de profil
func (h *UpdateProfilHandler) HandleUpdateProfil(w http.ResponseWriter, r *http.Request) {
    // Récupère l'ID de l'utilisateur à partir du token JWT
    userClaims := r.Context().Value("user").(*utils.JWTClaims)
    userID := userClaims.UserID

    var updateData profilservice.UpdateProfileInput
    if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Mise à jour du profil de l'utilisateur via le service
    if err := h.profilService.UpdateBusinessUserProfile(userID, updateData); err != nil {
        http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
