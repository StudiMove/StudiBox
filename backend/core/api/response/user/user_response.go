package response

// UserResponse définit les informations utilisateur renvoyées après l'inscription ou la connexion
type GetUserProfileResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}
