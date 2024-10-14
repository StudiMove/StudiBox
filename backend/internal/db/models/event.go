package models

import "gorm.io/gorm"

type Event struct {
    gorm.Model
    OwnerID       uint   // Référence à l'utilisateur ou à l'association/école
    OwnerType     string // 'business', 'school', 'association'
    ImageURLs     string // JSON pour stocker les URLs des images
    VideoURL      string
    Title         string `gorm:"not null"`
    Subtitle      string
    Description   string `gorm:"not null"`
    StartDate     string // Date de début
    EndDate       string // Date de fin
    StartTime     string // Heure de début
    EndTime       string // Heure de fin
    isOnline      bool
    isVisible     bool
    Price         int
    Address       string
    City          string
    Postcode      string
    Region        string
    Country       string
    Statistics    string // JSON pour stocker des statistiques
    Category      string
    Tags          string
    TicketsSold   int    `gorm:"default:0"`
    Revenue       int    `gorm:"default:0"`
    
    // Nouvelle relation avec les options
    Options       []EventOption `gorm:"foreignKey:EventID"`
    
    CreatedAt     string `gorm:"not null"` // Date de création
    UpdatedAt     string `gorm:"not null"` // Date de mise à jour
}
