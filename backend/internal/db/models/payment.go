package models

import "gorm.io/gorm"

type Payment struct {
    gorm.Model
    PaymentID          uint   `gorm:"primaryKey"`
    UserID             uint   // Référence à l'utilisateur
    Amount             int    // Montant total du paiement
    Status             string // Statut du paiement
    PaymentDate        string
    InstallmentNumber  int    // Numéro du paiement
    TotalInstallments   int    // Nombre total de paiements
    CancellationDate    string // Date d'annulation si applicable
    CreatedAt          string `gorm:"not null"` // Date de création
    UpdatedAt          string `gorm:"not null"` // Date de mise à jour
}
