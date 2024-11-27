package response

// UploadProfileImageWithTargetIDResponse représente la réponse après l'upload d'une image de profil avec un ID cible.
type UploadProfileImageWithTargetIDResponse struct {
    URL     string `json:"url"`
    Success bool   `json:"success"`
    Message string `json:"message"`
}
