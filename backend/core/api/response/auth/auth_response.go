package response

// LoginResponse définit les données renvoyées après une connexion réussie
type LoginResponse struct {
	Token           string `json:"token"`
	IsAuthenticated bool   `json:"isAuthenticated"`
}

// RegisterResponse définit la réponse après une inscription réussie
type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
