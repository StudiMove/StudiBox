package profil

import (
	"backend/internal/db/models"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type GetUserProfileImageHandler struct {
    db *gorm.DB
}

func NewGetUserProfileImageHandler(db *gorm.DB) *GetUserProfileImageHandler {
    return &GetUserProfileImageHandler{
        db: db,
    }
}

func (h *GetUserProfileImageHandler) HandleGetProfileImage(w http.ResponseWriter, r *http.Request) {
    // Définir le type de contenu comme application/json
    w.Header().Set("Content-Type", "application/json")

    // Définir une structure pour récupérer le corps de la requête
    var requestData struct {
        UserID int `json:"user_id"`
    }

    // Décoder le JSON reçu dans requestData
    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        log.Printf("Failed to decode JSON: %v", err)
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Vérifier si l'ID utilisateur est présent
    if requestData.UserID == 0 {
        http.Error(w, "Missing or invalid user ID", http.StatusBadRequest)
        return
    }

    // Chercher l'utilisateur dans la base de données
    var user models.User
    if err := h.db.Where("id = ?", requestData.UserID).First(&user).Error; err != nil {
        log.Printf("User not found: %v", err)
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Vérifier que l'utilisateur a une image de profil enregistrée
    if user.ProfileImage == "" {
        log.Printf("User %d has no profile image", requestData.UserID)
        http.Error(w, "No profile image found", http.StatusNotFound)
        return
    }

    // Renvoyer l'URL de l'image de profil dans la réponse
    json.NewEncoder(w).Encode(map[string]string{"url": user.ProfileImage})
}
