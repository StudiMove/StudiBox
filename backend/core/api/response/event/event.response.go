package response

import (
	"time"
)

type DescriptionResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type OptionResponse struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type TarifResponse struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type EventResponse struct {
	ID           uint                  `json:"id"`
	OwnerID      uint                  `json:"owner_id"`
	OwnerType    string                `json:"owner_type"`
	ImageURL     string                `json:"image_urls"`
	VideoURL     string                `json:"video_url"`
	Title        string                `json:"title"`
	Subtitle     string                `json:"subtitle"`
	Description  string                `json:"description"`
	StartDate    time.Time             `json:"start_date"`
	EndDate      time.Time             `json:"end_date"`
	StartTime    time.Time             `json:"start_time"`
	EndTime      time.Time             `json:"end_time"`
	IsOnline     bool                  `json:"is_online"`
	IsPublic     bool                  `json:"is_visible"`
	TicketPrice  float64               `json:"ticket_price"`
	Address      string                `json:"address"`
	City         string                `json:"city"`
	PostalCode   int32                 `json:"postalcode"`
	Region       string                `json:"region"`
	Country      string                `json:"country"`
	Categories   []string              `json:"categories"`
	Tags         []string              `json:"tags"`
	Descriptions []DescriptionResponse `json:"descriptions"`
	Options      []OptionResponse      `json:"options"`
	Tarifs       []TarifResponse       `json:"tarifs"`
	Likes        int                   `json:"likes"`
	Views        int                   `json:"views"`
}

type ListEventsResponse struct {
	Events []EventResponse `json:"events"`
	Page   int             `json:"page"`
	Total  int64           `json:"total"`
}
