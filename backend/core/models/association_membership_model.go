package models

import "gorm.io/gorm"

type AssociationMembership struct {
	gorm.Model
	MembershipID  uint `gorm:"primaryKey"`
	UserID        uint // Référence à l'utilisateur
	AssociationID uint // Référence à l'association
	JoinDate      string
	CreatedAt     string `gorm:"not null"` // Date de création
	UpdatedAt     string `gorm:"not null"` // Date de mise à jour
}
