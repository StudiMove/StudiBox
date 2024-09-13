package models

import "gorm.io/gorm"

type SchoolMembership struct {
    gorm.Model
    MembershipID  uint   `gorm:"primaryKey"`
    UserID        uint   // Référence à l'utilisateur
    SchoolID      uint   // Référence à l'établissement éducatif
    JoinDate      string
    CreatedAt     string `gorm:"not null"` // Date de création
    UpdatedAt     string `gorm:"not null"` // Date de mise à jour
}
