package response

// UserResponse définit les informations utilisateur renvoyées après l'inscription ou la connexion
type UserAuthResponse struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Pseudo      string `json:"pseudo" validate:"required,unique"`
	Email       string `json:"email" validate:"required,email"`
	ProfileType string `json:"profile_type" validate:"required,oneof=etudiant non_etudiant"`
	Type        string `json:"type" validate:"required,oneof=association owner school"`
	Role        string `json:"role"`
}

// LoginResponse définit les données renvoyées après une connexion réussie
type LoginResponse struct {
	Token           string           `json:"token"`
	IsAuthenticated bool             `json:"isAuthenticated"`
	User            UserAuthResponse `json:"user"`
}

// RegisterResponse définit la réponse après une inscription réussie
type RegisterResponse struct {
	ID    uint             `json:"id"`
	Email string           `json:"email"`
	Role  string           `json:"role"`
	Token string           `json:"token"`
	User  UserAuthResponse `json:"user"`
}
