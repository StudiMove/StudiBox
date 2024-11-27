package request

// EventDescriptionRequest représente une description pour un événement
type EventDescriptionRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}