package models

import "gorm.io/gorm"

type SchoolMembership struct {
    gorm.Model
    MembershipID uint   `gorm:"primaryKey" json:"membershipId"`
    UserID       uint   `json:"userId"`          // Référence à l'utilisateur
    SchoolID     uint   `json:"schoolId"`        // Référence à l'établissement éducatif
    JoinDate     string `json:"joinDate"`
    CreatedAt    string `gorm:"not null" json:"createdAt"` // Date de création
    UpdatedAt    string `gorm:"not null" json:"updatedAt"` // Date de mise à jour
}
