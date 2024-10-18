package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	OwnerID       uint   // Référence à l'utilisateur ou à l'association/école
	OwnerType     string // 'business', 'school', 'association'
	ImageURLsJSON string // JSON pour stocker les URLs des images dans la base de données
	VideoURL      string
	Title         string `gorm:"not null"`
	Subtitle      string
	Description   string `gorm:"not null"`
	StartDate     string
	EndDate       string
	StartTime     string
	EndTime       string
	IsOnline      bool
	IsVisible     bool
	Price         int
	Address       string
	City          string
	Postcode      string
	Region        string
	Country       string
	Statistics    string // JSON pour stocker des statistiques
	Category      string
	Tags          string
	TicketsSold   int           `gorm:"default:0"`
	Revenue       int           `gorm:"default:0"`
	Options       []EventOption `gorm:"foreignKey:EventID"`

	// Non persistant, utilisé uniquement pour manipulation dans le code
	ImageURLs []string `gorm:"-"` // Ignorer dans la base de données

	CreatedAt string `gorm:"not null"`
	UpdatedAt string `gorm:"not null"`
}
