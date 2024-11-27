package models

import "gorm.io/gorm"

type AssociationMembership struct {
	gorm.Model
	MembershipID  uint   `gorm:"primaryKey" json:"membershipId"`
	UserID        uint   `json:"userId"`        // Référence à l'utilisateur
	AssociationID uint   `json:"associationId"` // Référence à l'association
	JoinDate      string `json:"joinDate"`
	CreatedAt     string `gorm:"not null" json:"createdAt"`
	UpdatedAt     string `gorm:"not null" json:"updatedAt"`
}
