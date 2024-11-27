package models

import "gorm.io/gorm"

type Payment struct {
    gorm.Model
    PaymentID         uint   `gorm:"primaryKey" json:"paymentId"`
    UserID            uint   `json:"userId"`             // Référence à l'utilisateur
    Amount            int    `json:"amount"`             // Montant total du paiement
    Status            string `json:"status"`             // Statut du paiement
    PaymentDate       string `json:"paymentDate"`
    InstallmentNumber int    `json:"installmentNumber"`  // Numéro du paiement
    TotalInstallments int    `json:"totalInstallments"`  // Nombre total de paiements
    CancellationDate  string `json:"cancellationDate"`   // Date d'annulation si applicable
    CreatedAt         string `gorm:"not null" json:"createdAt"` // Date de création
    UpdatedAt         string `gorm:"not null" json:"updatedAt"` // Date de mise à jour
}
