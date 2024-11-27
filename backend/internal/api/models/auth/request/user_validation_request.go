package request

// UserValidationRequest représente la requête pour valider l'utilisateur selon son rôle
type UserValidationRequest struct {
	UserID uint   `json:"userId"`
	Role   string `json:"role"`
}
