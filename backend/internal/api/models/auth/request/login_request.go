package request

// LoginRequest représente les données envoyées pour une demande de connexion.
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
