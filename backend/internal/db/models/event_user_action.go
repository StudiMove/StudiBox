package models

import (
	"time"

	"gorm.io/gorm"
)

type EventUserAction struct {
	gorm.Model
	UserID       uint      `json:"userId"`                            // Référence à la table User
	EventID      uint      `json:"eventId"`                           // Référence à la table Event
	IsInterested bool      `gorm:"default:false" json:"isInterested"` // L'utilisateur est intéressé par l'événement
	IsFavorite   bool      `gorm:"default:false" json:"isFavorite"`   // L'utilisateur a marqué l'événement comme favori
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`   // Créé automatiquement lors de l'insertion
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`   // Mise à jour automatiquement lors des modifications
}
