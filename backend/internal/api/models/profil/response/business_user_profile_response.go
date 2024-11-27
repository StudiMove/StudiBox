package response

type BusinessUserProfileResponse struct {
    Name string `json:"name"`
    SIRET       string `json:"siret"`
    Address     string `json:"address"`
    City        string `json:"city"`
    Region          string `json:"region"`
    Postcode    string `json:"postcode"`
    Country     string `json:"country"`
    Description string `json:"description"`
    Status      string `json:"status"` // Ajout du champ Status

}
