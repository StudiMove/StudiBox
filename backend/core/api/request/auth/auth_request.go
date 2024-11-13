package request

// LoginRequest définit les données envoyées pour la connexion
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterUserRequest représente la requête pour inscrire un utilisateur standard
type RegisterUserRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Pseudo      string `json:"pseudo" binding:"required"`
	Phone       int    `json:"phone" binding:"required"`
	ProfileType string `json:"profile_type" binding:"required"`
}

// RegisterOwnerRequest représente la requête pour inscrire un utilisateur propriétaire (association, entreprise, école)
type RegisterOwnerRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Pseudo      string `json:"pseudo" validate:"required,unique"`
	Phone       int    `json:"phone" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	CompanyName string `json:"company_name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=association owner school"`
	Address     string `json:"address" validate:"required"`
	PostalCode  int32  `json:"postalcode" validate:"required"`
	City        string `json:"city" validate:"required"`
	Country     string `json:"country" validate:"required"`
	Region      string `json:"region,omitempty"`
	Description string `json:"description,omitempty"`
	SchoolID    uint   `json:"school_id,omitempty"`
}
