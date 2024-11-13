package response

// OwnerResponse représente la structure de réponse pour le profil propriétaire
type OwnerResponse struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  int32  `json:"postal_code"`
	Region      string `json:"region"`
	Country     string `json:"country"`
	Phone       int    `json:"phone"`
	Description string `json:"description"`
	Status      string `json:"status"`
	SIRET       string `json:"siret"`
}
