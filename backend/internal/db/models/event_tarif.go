package models

import "gorm.io/gorm"

type EventTarif struct {
    gorm.Model
    EventID     uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"eventId"` // Référence à l'événement parent
    Title       string `gorm:"not null" json:"title"`              // Titre du tarif
    Price       float64    `json:"price"`                              // Prix du tarif
    Stock       int    `json:"stock"`                              // Stock du tarif
    Description string `gorm:"type:text" json:"description"`       // Description du tarif
}
