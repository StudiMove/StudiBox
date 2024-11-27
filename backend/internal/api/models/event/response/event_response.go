package response

import "time"

// EventResponse représente les détails de l'événement dans la réponse
type EventResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	IsOnline    bool      `json:"is_online"`
	TicketPrice float64   `json:"ticket_price"`
	TicketStock int       `json:"ticket_stock"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Postcode    string    `json:"postcode"`
	Region      string    `json:"region"`
	Country     string    `json:"country"`
	Categories  []string  `json:"categories"`
	Tags        []string  `json:"tags"`
	ImageURLs   string    `json:"image_urls"`
	VideoURL    string    `json:"video_url"`
}

type EventOptionResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type EventTarifResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type UploadEventImageResponse struct {
	URLs []string `json:"urls"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// EventDescriptionResponse représente une description pour l'événement.
type EventDescriptionResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type EventCategoryResponse struct {
	Name string `json:"name"`
}

// CategoryResponse représente les informations d'une catégorie
type CategoryResponse struct {
	Name string `json:"name"`
}

// TagResponse représente les informations d'un tag
type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
