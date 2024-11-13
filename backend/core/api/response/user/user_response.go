package response

// UserResponse définit les informations utilisateur renvoyées après l'inscription ou la connexion
type UserResponse struct {
	ID            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Pseudo        string `json:"pseudo"`
	Email         string `json:"email"`
	Phone         int    `json:"phone"`
	ProfileImage  string `json:"profile_image"`
	BirthDate     string `json:"birth_date"`
	Country       string `json:"country"`
	Region        string `json:"region"`
	City          string `json:"city"`
	Address       string `json:"address"`
	PostalCode    int32  `json:"postal_code"`
	ProfileType   string `json:"profile_type"`
	StudiboxCoins int    `json:"studibox_coins"`
	Role          string `json:"role"`
}
