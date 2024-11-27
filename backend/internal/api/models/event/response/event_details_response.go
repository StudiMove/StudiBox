package response

import "time"

// EventDetailResponse représente les détails complets de l'événement pour une réponse JSON.
type EventDetailResponse struct {
	ID                 uint                       `json:"id"`
	Title              string                     `json:"title"`
	Subtitle           string                     `json:"subtitle"`
	StartDate          time.Time                  `json:"start_date"`
	EndDate            time.Time                  `json:"end_date"`
	StartTime          time.Time                  `json:"start_time"`
	EndTime            time.Time                  `json:"end_time"`
	IsOnline           bool                       `json:"is_online"`
	IsActivated        bool                       `json:"is_activated"`
	TicketPrice        float64                    `json:"ticket_price"`
	TicketStock        int                        `json:"ticket_stock"`
	Address            string                     `json:"address"`
	City               string                     `json:"city"`
	Postcode           string                     `json:"postcode"`
	Region             string                     `json:"region"`
	Country            string                     `json:"country"`
	Categories         []string                   `json:"categories"`
	Tags               []string                   `json:"tags"`
	Options            []EventOptionResponse      `json:"options"`
	Tarifs             []EventTarifResponse       `json:"tarifs"`
	Descriptions       []EventDescriptionResponse `json:"descriptions"`
	ImageURLs          string                     `json:"image_urls"`
	IsValidatedByAdmin bool                       `json:"isValidatedByAdmin"`
	UseStudibox        bool                       `json:"use_studibox"` // Ajouter ce champ
	VideoURL           string                     `json:"video_url"`
}
