package models

import (
	"time"

	"gorm.io/gorm"
)

// Constantes pour le statut de la transaction de paiement
const (
	PaymentTransactionStatusPending   = "pending"
	PaymentTransactionStatusCompleted = "completed"
	PaymentTransactionStatusCancelled = "cancelled"
)

type PaymentTransaction struct {
	gorm.Model
	PaymentID        uint `gorm:"not null"`
	TicketID         uint
	Amount           int    `gorm:"default:0"`
	Status           string `gorm:"size:20;default:'pending';check:status IN ('pending', 'completed', 'cancelled')"`
	TransactionDate  time.Time
	CancellationDate time.Time
}
