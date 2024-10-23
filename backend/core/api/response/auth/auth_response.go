// package response

package response

// UserResponse définit les informations utilisateur renvoyées après l'inscription ou la connexion
type UserResponse struct {
	ID        uint     `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Roles     []string `json:"roles"` // Liste des noms de rôles sous forme de chaînes de caractères
}

// LoginResponse définit les données renvoyées après une connexion réussie
type LoginResponse struct {
	Token           string       `json:"token"`
	IsAuthenticated bool         `json:"isAuthenticated"`
	User            UserResponse `json:"user"`
}

// RegisterResponse définit la réponse après une inscription réussie
type RegisterResponse struct {
	ID    uint         `json:"id"`
	Email string       `json:"email"`
	Role  string       `json:"role"`
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
