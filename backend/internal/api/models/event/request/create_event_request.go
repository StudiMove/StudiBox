// package request

// import "time"

package request

import (
	"encoding/json"
	"log"
	"time"

	"backend/internal/db/models"
)

type CreateEventRequest struct {
	HostID uint `json:"-"`
	UserID uint `json:"user_id"`

	Title                string                    `json:"title"`
	Subtitle             string                    `json:"subtitle"`
	StartDate            time.Time                 `json:"start_date"`
	EndDate              time.Time                 `json:"end_date"`
	StartTime            time.Time                 `json:"start_time"`
	EndTime              time.Time                 `json:"end_time"`
	IsOnline             bool                      `json:"is_online"`
	IsVisible            bool                      `json:"is_visible"`
	UseStudibox          bool                      `json:"use_studibox"`
	TicketPrice          float64                   `json:"ticket_price"`
	TicketStock          int                       `json:"ticket_stock"`
	Location             EventLocationRequest      `json:"location"`
	Categories           []string                  `json:"categories"`
	Tags                 []string                  `json:"tags"`
	Options              []EventOptionRequest      `json:"options"`
	Tarifs               []EventTarifRequest       `json:"tarifs"`
	Descriptions         []EventDescriptionRequest `json:"descriptions"`
	ExternalTicketingUrl string                    `json:"external_ticketing_url"`
	VideoURL             string                    `json:"video_url"`
	Images               string                    `json:"images"` // Tableau de chaînes
}

// ToEventModel transforme CreateEventRequest en models.Event
func (req *CreateEventRequest) ToEventModel() (*models.Event, error) {
	imageURLsJSON, err := json.Marshal(req.Images)
	if err != nil {
		log.Printf("Erreur lors de la conversion des URLs d'images en JSON : %v", err)
		return nil, err
	}

	event := &models.Event{
		UserID:      req.UserID, // Inclure l'ID utilisateur
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		IsOnline:    req.IsOnline,
		IsVisible:   req.IsVisible,
		UseStudibox: req.UseStudibox,
		TicketPrice: req.TicketPrice,
		TicketStock: req.TicketStock,
		Address:     req.Location.Address,
		City:        req.Location.City,
		Postcode:    req.Location.Postcode,
		Region:      req.Location.Region,
		Country:     req.Location.Country,
		VideoURL:    req.VideoURL,
		ImageURLs:   string(imageURLsJSON),
	}

	return event, nil
}
