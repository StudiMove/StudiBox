package models

import "gorm.io/gorm"

type BusinessUser struct {
    gorm.Model
    UserID        uint   `gorm:"primaryKey"` // Référence à l'utilisateur
	User          User   // Relation avec le modèle User
    CompanyName   string `gorm:"not null"`
    SIRET         string
    Address       string
    CreationDate  string
    IsValidated   bool   `gorm:"default:false"` // Validé par un admin
    CreatedAt     string `gorm:"not null"`      // Date de création
    UpdatedAt     string `gorm:"not null"`      // Date de mise à jour
}
