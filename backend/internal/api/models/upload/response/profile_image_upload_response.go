package response

// ProfileImageUploadResponse représente la réponse après l'upload d'une image de profil.
type ProfileImageUploadResponse struct {
    URL     string `json:"url"`
    Success bool   `json:"success"`
    Message string `json:"message"`
}

