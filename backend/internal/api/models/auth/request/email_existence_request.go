package request

// EmailExistenceRequest représente la requête pour vérifier si un email existe
type EmailExistenceRequest struct {
	Email string `json:"email"`
}
