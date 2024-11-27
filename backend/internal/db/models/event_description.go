package models

import "gorm.io/gorm"

type EventDescription struct {
    gorm.Model
    EventID     uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"eventId"` // Référence à l'événement parent
    Title       string `gorm:"not null" json:"title"`           // Titre de la description
    Description string `gorm:"not null" json:"description"`     // Contenu de la description
}
