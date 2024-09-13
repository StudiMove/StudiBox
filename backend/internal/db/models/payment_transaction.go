package models

import "gorm.io/gorm"

type PaymentTransaction struct {
    gorm.Model
    TransactionID      uint   `gorm:"primaryKey"`
    PaymentID          uint   // Référence à l'identifiant de paiement
    TicketID           uint   // Référence au billet si applicable
    Amount             int    // Montant de la transaction
    Status             string // Statut de la transaction
    TransactionDate    string
    CancellationDate    string // Date d'annulation si applicable
    CreatedAt          string `gorm:"not null"` // Date de création
    UpdatedAt          string `gorm:"not null"` // Date de mise à jour
}
