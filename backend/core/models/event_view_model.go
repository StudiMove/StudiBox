// models/event_view.go
package models

import "gorm.io/gorm"

// EventView représente une vue d'un événement par un utilisateur
type EventView struct {
	gorm.Model
	UserID  uint // Référence à l'utilisateur qui a vu l'événement
	EventID uint // Référence à l'événement vu
	Count   int  // Nombre de fois que l'utilisateur a vu cet événement
}
