package models

import (
	"gorm.io/gorm"
)

type OwnerRelationship struct {
	gorm.Model
	SchoolID      uint `gorm:"not null"`
	AssociationID uint `gorm:"not null"`
}
