package models

import (
	"time"

	"gorm.io/gorm"
)

// StudiboxTransaction représente une transaction pour les coins Studibox
type StudiboxTransaction struct {
	gorm.Model
	UserID          uint      `json:"userId"`
	Amount          float64   `json:"amount"`     // Montant de la transaction
	ReferrerID      *uint     `json:"referrerId"` // Peut être null
	CoinsToUser     int       `json:"coinsToUser"`
	CoinsToReferrer int       `json:"coinsToReferrer"`
	CoinsUsed       int       `json:"coinsUsed"`
	EventEndDate    time.Time `json:"eventEndDate"`
	PaymentStatus   string    `json:"paymentStatus"` // Exemples : "pending", "succeeded", "refunded"
	GlobalStatus    string    `json:"globalStatus"`  // Exemples : "pending", "done", "refunded"
	PaymentID       string    `json:"paymentId"`     // ID unique du paiement Stripe
	ChargeID        string    `json:"ChargeId"`      // ID unique du paiement Stripe
	KlarnaStatus    string    `json:"klarnaStatus"`
	DisputeID       string    `json:"disputeId"`
}
