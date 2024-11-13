package request

import "time"

// UserRequest représente les champs pour la mise à jour du profil utilisateur
type UserRequest struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        int       `json:"phone"`
	Pseudo       string    `json:"pseudo"`
	ProfileImage string    `json:"profile_image"`
	BirthDate    time.Time `json:"birth_date"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	PostalCode   int32     `json:"postal_code"`
	ProfileType  string    `json:"profile_type" binding:"omitempty,oneof=etudiant non_etudiant"`
}
