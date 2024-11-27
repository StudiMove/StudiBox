package request

type UpdateProfileRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	Region      string `json:"region"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	Email       string `json:"email"`
	SIRET       string `json:"siret"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	IsValidated bool   `json:"is_validated"`
	IsPending   bool   `json:"is_pending"`
	IsActivated bool   `json:"is_activated"`
}
