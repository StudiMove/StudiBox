package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Pseudo         string    `json:"pseudo"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Password       string    `gorm:"not null" json:"password"`
	Phone          string    `json:"phone"`
	ProfileImage   string    `json:"profileImage"`
	BirthDate      string    `json:"birthDate"`
	City           string    `json:"city"`
	ProfileType    string    `json:"profileType"` // 'étudiant', 'non étudiant'
	ParrainageCode string    `json:"parrainageCode"`
	ParrainCode    string    `gorm:"unique" json:"parrainCode"` // Code unique généré pour parrainer d'autres utilisateurs
	AssociationID  uint      `json:"associationId"`
	SchoolID       uint      `json:"schoolId"`
	StudiboxCoins  int       `gorm:"default:0" json:"studiboxCoins"`    // Solde total des Studibox Coins
	Roles          []Role    `gorm:"many2many:user_roles" json:"roles"` // Relation plusieurs à plusieurs avec les rôles
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`   // Crée automatiquement lors de l'insertion
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updatedAt"`   // Met à jour automatiquement à chaque modification
}
