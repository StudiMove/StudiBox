package response

// GetUserIDByEmailResponse représente la réponse contenant l'ID de l'utilisateur
type GetUserIDByEmailResponse struct {
    UserID  uint   `json:"user_id"`
    Success bool   `json:"success"`
    Message string `json:"message,omitempty"`
}
