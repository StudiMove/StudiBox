package models

import "gorm.io/gorm"

// Modèle pour représenter un like sur un événement
type EventLike struct {
	gorm.Model
	UserID  uint // Référence à l'utilisateur qui a liké l'événement
	EventID uint // Référence à l'événement liké
}
