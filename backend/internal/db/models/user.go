package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName      string `gorm:"not null"`
    LastName       string `gorm:"not null"`
    Pseudo         string `gorm:"unique;not null"`
    Email          string `gorm:"unique;not null"`
    Password       string `gorm:"not null"`
    Phone          string
    ProfileImage   string
    BirthDate      string
    City           string
    ProfileType    string // 'étudiant', 'non étudiant'
    AssociationID  uint   // Référence à l'association, peut être nul
    StudiboxCoins  int    `gorm:"default:0"` // Solde total des Studibox Coins
	Roles          []Role  `gorm:"many2many:user_roles;"` // Relation plusieurs à plusieurs avec les rôles
    CreatedAt      string `gorm:"not null"`   // Date de création
    UpdatedAt      string `gorm:"not null"`   // Date de mise à jour
}
