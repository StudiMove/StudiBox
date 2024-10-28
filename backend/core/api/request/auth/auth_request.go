package request

// LoginRequest définit les données envoyées pour la connexion
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterUserRequest représente la requête pour inscrire un utilisateur standard
type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	IDToken   string `json:"id_token,omitempty"`
}

// RegisterBusinessRequest représente la requête pour inscrire un utilisateur entreprise
type RegisterBusinessRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Address     string `json:"address" validate:"required"`
	Postcode    string `json:"postcode" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	City        string `json:"city" validate:"required"`
	Country     string `json:"country" validate:"required"`
}
