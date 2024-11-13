// package request
package request

// UpdateOwnerRequest représente les champs pour la mise à jour du profil owner
type UpdateOwnerRequest struct {
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  int32  `json:"postal_code"`
	Country     string `json:"country"`
	Phone       int    `json:"phone"`
	Region      string `json:"region"`
	SIRET       string `json:"siret"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
}
