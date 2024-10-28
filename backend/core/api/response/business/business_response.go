package response

// GetBusinessProfileResponse représente la structure de réponse pour le profil business
type BusinessProfileResponse struct {
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
}
