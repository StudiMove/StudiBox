package models

import (
	"gorm.io/gorm"
)

type EventView struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	EventID uint `gorm:"not null"`
	Count   int  `gorm:"default:0"`
}
