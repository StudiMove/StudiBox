package models

import "gorm.io/gorm"

type BusinessUser struct {
    gorm.Model
    UserID        uint   `gorm:"primaryKey"`  // Référence à l'utilisateur
    User          User   `gorm:"foreignKey:UserID"` // Relation avec le modèle User
    CompanyName   string `gorm:"not null"`
    SIRET         string
    Address       string
    City          string
    Postcode      string
    Region        string
    Country       string
    CreationDate  string
    IsValidated   bool   `gorm:"default:false"` // Validé par un admin
    IsActivated   bool   `gorm:"default:false"` // Activé par un admin
    CreatedAt     string `gorm:"not null"`      // Date de création
    UpdatedAt     string `gorm:"not null"`      // Date de mise à jour
}
