package response

// UserValidationResponse représente la réponse pour la validation de l'utilisateur
type UserValidationResponse struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}
