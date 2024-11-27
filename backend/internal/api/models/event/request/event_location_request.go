package request

// EventLocationRequest représente la localisation d'un événement
type EventLocationRequest struct {
    Address  string `json:"address"`
    City     string `json:"city"`
    Postcode string `json:"postcode"`
    Region   string `json:"region"`
    Country  string `json:"country"`
}