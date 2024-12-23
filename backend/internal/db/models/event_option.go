package models

import "gorm.io/gorm"

type EventOption struct {
	gorm.Model
	EventID     uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"eventId"` // Référence à l'événement parent
	Title       string  `gorm:"not null" json:"title"`                                                 // Titre de l'option
	Description string  `gorm:"type:text" json:"description"`                                          // Description de l'option
	Price       float64 `json:"price"`                                                                 // Prix de l'option
	Stock       int     `json:"stock"`                                                                 // Prix de l'option
	PriceID     string  `gorm:"unique;" json:"price_id"`                                               // Ajouter une contrainte unique et obligatoire
}
