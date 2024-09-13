package models

import "gorm.io/gorm"

type Association struct {
    gorm.Model
    UserID        uint   `gorm:"primaryKey"` // Référence à l'utilisateur
    AssociationName string `gorm:"not null"`
    Address       string
    Phone         string
    ProfileImage  string
    CreationDate  string
    MemberCount   int    `gorm:"default:0"` // Nombre de membres
    CreatedAt     string `gorm:"not null"`  // Date de création
    UpdatedAt     string `gorm:"not null"`  // Date de mise à jour
}
