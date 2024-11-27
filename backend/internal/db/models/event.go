package models

import (
	"time"

	"gorm.io/gorm"
)

// Modèle d'événement
type Event struct {
	gorm.Model
	HostID             uint      `json:"hostId"`   // Référence à l'hôte
	HostType           string    `json:"hostType"` // Type de l'hôte
	UserID             uint      `json:"user_id"`
	ImageURLs          string    `json:"imageUrls"` // Stocke les URLs des images en format JSON
	VideoURL           string    `json:"videoUrl"`
	Title              string    `gorm:"not null" json:"title"`
	Subtitle           string    `json:"subtitle"`
	StartDate          time.Time `json:"startDate"`                    // Date de début
	EndDate            time.Time `json:"endDate"`                      // Date de fin
	StartTime          time.Time `json:"startTime"`                    // Heure de début
	EndTime            time.Time `json:"endTime"`                      // Heure de fin
	IsOnline           bool      `json:"isOnline"`                     // Événement en ligne ?
	IsVisible          bool      `json:"isVisible"`                    // Événement public ?
	UseStudibox        bool      `json:"useStudibox"`                  // Indique si l'événement utilise la billetterie Studibox
	TicketPrice        float64   `json:"ticketPrice"`                  // Prix des billets
	TicketStock        int       `json:"ticketStock"`                  // Stock des billets
	Address            string    `json:"address"`                      // Adresse
	City               string    `json:"city"`                         // Ville
	Postcode           string    `json:"postcode"`                     // Code postal
	Region             string    `json:"region"`                       // Région
	Country            string    `json:"country"`                      // Pays
	Statistics         string    `json:"statistics"`                   // Stocke des statistiques
	TicketsSold        int       `gorm:"default:0" json:"ticketsSold"` // Nombre de billets vendus
	Revenue            int       `gorm:"default:0" json:"revenue"`     // Chiffre d'affaires
	IsValidatedByAdmin bool      `json:"isValidatedByAdmin"`           // Validation par un admin

	// Relations
	Descriptions []EventDescription `gorm:"foreignKey:EventID" json:"descriptions"` // Les descriptions de l'événement
	Options      []EventOption      `gorm:"foreignKey:EventID" json:"options"`      // Les options de l'événement
	Tarifs       []EventTarif       `gorm:"foreignKey:EventID" json:"tarifs"`       // Les tarifs de l'événement
	Ticket       []Ticket           `gorm:"foreignKey:EventID" json:"tickets"`      // Les tickets de l'événement
	Categories   []EventCategory    `gorm:"many2many:event_category_events" json:"categories"`
	Tags         []EventTag         `gorm:"many2many:event_tag_events" json:"tags"`
}
