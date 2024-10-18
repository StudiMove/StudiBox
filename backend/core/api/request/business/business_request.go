package request

// UpdateBusinessProfileRequest représente les champs pour la mise à jour du profil business
type UpdateBusinessProfileRequest struct {
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}
