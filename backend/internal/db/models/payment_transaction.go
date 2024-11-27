package models

import "gorm.io/gorm"

type PaymentTransaction struct {
    gorm.Model
    TransactionID    uint   `gorm:"primaryKey" json:"transactionId"`
    PaymentID        uint   `json:"paymentId"`           // Référence à l'identifiant de paiement
    TicketID         uint   `json:"ticketId"`            // Référence au billet si applicable
    Amount           int    `json:"amount"`              // Montant de la transaction
    Status           string `json:"status"`              // Statut de la transaction
    TransactionDate  string `json:"transactionDate"`
    CancellationDate string `json:"cancellationDate"`    // Date d'annulation si applicable
    CreatedAt        string `gorm:"not null" json:"createdAt"` // Date de création
    UpdatedAt        string `gorm:"not null" json:"updatedAt"` // Date de mise à jour
}
