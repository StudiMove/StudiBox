package models

import (
	"gorm.io/gorm"
)

type EventDescription struct {
	gorm.Model
	EventID     uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`
}
