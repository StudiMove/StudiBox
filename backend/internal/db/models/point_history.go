package models

import "gorm.io/gorm"

type PointHistory struct {
    gorm.Model
    UserID            uint   // Référence à l'utilisateur
    Points            int    // Points gagnés ou dépensés
    TransactionID     uint   // Référence à la transaction associée
    EventID           uint   // Référence à l'événement associé (si applicable)
    ChangeDate        string
    Type              string // 'earned', 'spent'
    CreatedAt         string `gorm:"not null"` // Date de création
    UpdatedAt         string `gorm:"not null"` // Date de mise à jour
}
