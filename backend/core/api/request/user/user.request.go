// package request

package request

// UpdateUserProfileRequest représente les champs pour la mise à jour du profil utilisateur
type UpdateUserProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
