package models

import "gorm.io/gorm"

type StudiboxTransaction struct {
    gorm.Model
    UserID          uint   `json:"userId"`           // Référence à l'utilisateur
    Amount          int    `json:"amount"`           // Montant de la transaction
    Status          string `json:"status"`           // 'pending', 'confirmed', 'cancelled'
    TransactionDate string `json:"transactionDate"`

}
