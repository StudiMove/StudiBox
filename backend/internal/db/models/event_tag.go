package models

import "gorm.io/gorm"

// Modèle pour les tags d'événements
type EventTag struct {
	gorm.Model
	Name   string  `gorm:"unique;not null" json:"name"`
	Events []Event `gorm:"many2many:event_tag_events" json:"events"`
}
