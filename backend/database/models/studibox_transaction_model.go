package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes pour le statut de la transaction
const (
	TransactionStatusPending   = "pending"
	TransactionStatusConfirmed = "confirmed"
	TransactionStatusCancelled = "cancelled"
)

type StudiboxTransaction struct {
	gorm.Model
	UserID          uint   `gorm:"not null"`
	Amount          int    `gorm:"default:0"`
	Status          string `gorm:"size:20;default:'pending';check:status IN ('pending', 'confirmed', 'cancelled')"`
	TransactionDate time.Time
}
