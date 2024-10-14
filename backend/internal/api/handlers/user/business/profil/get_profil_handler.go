package profil

import (
    "net/http"
    "encoding/json"
    "backend/internal/services/userservice/business/profilservice" // Corriger l'import
    "backend/internal/utils"
)

type GetProfilHandler struct {
    profilService *profilservice.ProfilService // Utilise le bon service
}

func NewGetProfilHandler(profilService *profilservice.ProfilService) *GetProfilHandler {
    return &GetProfilHandler{profilService: profilService}
}

// HandleGetProfil gère la récupération des informations de profil
func (h *GetProfilHandler) HandleGetProfil(w http.ResponseWriter, r *http.Request) {
    // Récupère l'ID de l'utilisateur à partir du token JWT
    userClaims := r.Context().Value("user").(*utils.JWTClaims)
    userID := userClaims.UserID

    // Récupère le profil de l'utilisateur via le service
    userProfile, err := h.profilService.GetBusinessUserByID(userID)
    if err != nil {
        http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(userProfile)
}
