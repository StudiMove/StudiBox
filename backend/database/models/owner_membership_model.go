package models

import (
	"gorm.io/gorm"
)

type OwnerMembership struct {
	gorm.Model
	MembershipID uint `gorm:"primaryKey"`
	UserID       uint
	OwnerID      uint
	JoinDate     string
}
