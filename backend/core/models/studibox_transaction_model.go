package models

import "gorm.io/gorm"

type StudiboxTransaction struct {
	gorm.Model
	UserID          uint   // Référence à l'utilisateur
	Amount          int    // Montant de la transaction
	Status          string // 'pending', 'confirmed', 'cancelled'
	TransactionDate string
	CreatedAt       string `gorm:"not null"` // Date de création
	UpdatedAt       string `gorm:"not null"` // Date de mise à jour
}
