package request

// GetUserIDByEmailRequest représente la requête pour obtenir l'ID de l'utilisateur par email
type GetUserIDByEmailRequest struct {
    Email string `json:"email"`
}
