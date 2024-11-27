package models

import "gorm.io/gorm"

// Modèle pour les catégories d'événements
type EventCategory struct {
	gorm.Model
	Name   string  `gorm:"unique;not null" json:"name"`
	Events []Event `gorm:"many2many:event_category_events" json:"events"`
}
