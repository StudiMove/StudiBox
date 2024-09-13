package models

import "gorm.io/gorm"

type Event struct {
    gorm.Model
    OwnerID       uint   // Référence à l'utilisateur ou à l'association/école
    OwnerType     string // 'business', 'school', 'association'
    ImageURLs     string // JSON pour stocker les URLs des images
    Title         string `gorm:"not null"`
    Subtitle      string
    Description   string `gorm:"not null"`
    StartDate     string
    EndDate       string
    Online        bool
    Price         int
    Address       string
    Statistics    string // JSON pour stocker des statistiques
    Category      string
    Tags          string
    TicketsSold   int    `gorm:"default:0"`
    Revenue       int    `gorm:"default:0"`
    CreatedAt     string `gorm:"not null"` // Date de création
    UpdatedAt     string `gorm:"not null"` // Date de mise à jour
}
