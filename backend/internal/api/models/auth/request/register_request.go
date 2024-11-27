package request

// RegisterRequest représente les données envoyées pour un enregistrement utilisateur.
type RegisterRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Phone        string `json:"phone"`
	ProfileType  string `json:"profileType" validate:"required"`
	ProfileImage string `json:"profileImage,omitempty"` // URL après l'upload
}

// RegisterUserRequest représente les données pour enregistrer un utilisateur standard.
type RegisterUserRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profileImage,omitempty"`
}

// RegisterOrganisationUserRequest représente les données pour enregistrer un utilisateur organisationnel.
type RegisterOrganisationUserRequest struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	OrganisationName string `json:"organisationName" validate:"required"`
	Address          string `json:"address"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Phone            string `json:"phone"`
	Description      string `json:"description"`
	OrganisationType string `json:"organisationType" validate:"required"`
}

// RegisterBusinessUserRequest représente les données pour enregistrer un utilisateur business.
type RegisterBusinessUserRequest struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	OrganisationName string `json:"organisationName" validate:"required"`
	Address          string `json:"address"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Phone            string `json:"phone"`
	Description      string `json:"description"`
	OrganisationType string `json:"organisationType" validate:"required"`
}

// RegisterSchoolUserRequest représente les données pour enregistrer un utilisateur de type school.
type RegisterSchoolUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	SchoolName  string `json:"schoolName" validate:"required"`
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}

// RegisterAssociationUserRequest représente les données pour enregistrer un utilisateur de type association.
type RegisterAssociationUserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	AssociationName string `json:"associationName" validate:"required"`
	Address         string `json:"address"`
	PostalCode      string `json:"postalCode"`
	City            string `json:"city"`
	Country         string `json:"country"`
	Phone           string `json:"phone"`
	Description     string `json:"description"`
}

type RegisterNormalUserRequest struct {
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	ParrainageCode string `json:"parrainageCode,omitempty"` // Optionnel
}
