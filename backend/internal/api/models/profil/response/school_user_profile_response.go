package response

type SchoolUserProfileResponse struct {
    Name		string `json:"name"`
    Address     string `json:"address"`
    City        string `json:"city"`
    Postcode    string `json:"postcode"`
    Region          string `json:"region"`
    Country     string `json:"country"`
    Description string `json:"description"`
    Status      string `json:"status"` // Ajout du champ Status

    
}
