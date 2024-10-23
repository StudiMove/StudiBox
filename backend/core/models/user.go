package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string
	LastName      string
	Pseudo        string
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	Phone         string
	ProfileImage  string
	BirthDate     string
	City          string
	ProfileType   string    // 'étudiant', 'non étudiant'
	AssociationID uint      // Référence à l'association, peut être nul
	StudiboxCoins int       `gorm:"default:0"`             // Solde total des Studibox Coins
	Roles         []Role    `gorm:"many2many:user_roles;"` // Relation plusieurs à plusieurs avec les rôles
	CreatedAt     time.Time `gorm:"autoCreateTime"`        // Crée automatiquement lors de l'insertion
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`        // Met à jour automatiquement à chaque modification
}
