package models

import "gorm.io/gorm"

// Modèle pour les options d'un événement
type EventOption struct {
	gorm.Model
	EventID     uint   // Référence à l'événement parent
	Title       string `gorm:"not null"`  // Titre de l'option
	Description string `gorm:"type:text"` // Description de l'option
	Price       int    // Prix de l'option
}
