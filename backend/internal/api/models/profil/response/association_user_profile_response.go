package response

type AssociationUserProfileResponse struct {
    Name string `json:"name"`
    Address         string `json:"address"`
    City            string `json:"city"`
    Region          string `json:"region"`
    Postcode        string `json:"postcode"`
    Country         string `json:"country"`
    Description     string `json:"description"`
    Status      string `json:"status"` 

}
