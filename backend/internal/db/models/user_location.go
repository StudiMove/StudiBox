package models

import (
	"time"

	"gorm.io/gorm"
)

type UserLocation struct {
	gorm.Model
	UserID       uint      `gorm:"not null" json:"userId"`          // Clé étrangère vers la table User
	Street       string    `json:"street"`                          // Rue
	NumberStreet string    `json:"numberStreet"`                    // Numéro de rue
	City         string    `json:"city"`                            // Ville
	Postcode     string    `json:"postcode"`                        // Code postal
	Region       string    `json:"region"`                          // Région
	Country      string    `json:"country"`                         // Pays
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"` // Créé automatiquement
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"` // Met à jour automatiquement
}
