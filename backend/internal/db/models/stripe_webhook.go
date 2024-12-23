package models

import "gorm.io/gorm"

type StripeWebhook struct {
	gorm.Model
	StripeEventID string `gorm:"not null;unique"`   // ID unique de l'événement Stripe
	EventType     string `gorm:"not null"`          // Type de l'événement (e.g., payment_intent.succeeded)
	Payload       string `gorm:"type:jsonb"`        // Contenu brut de l'événement (JSON)
	Status        string `gorm:"default:'pending'"` // État du traitement (e.g., 'pending', 'processed', 'failed')
}
