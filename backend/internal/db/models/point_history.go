package models

import "gorm.io/gorm"

type PointHistory struct {
    gorm.Model
    UserID        uint   `json:"userId"`           // Référence à l'utilisateur
    Points        int    `json:"points"`           // Points gagnés ou dépensés
    TransactionID uint   `json:"transactionId"`    // Référence à la transaction associée
    EventID       uint   `json:"eventId"`          // Référence à l'événement associé (si applicable)
    ChangeDate    string `json:"changeDate"`
    Type          string `json:"type"`             // 'earned', 'spent'
    CreatedAt     string `gorm:"not null" json:"createdAt"` // Date de création
    UpdatedAt     string `gorm:"not null" json:"updatedAt"` // Date de mise à jour
}
