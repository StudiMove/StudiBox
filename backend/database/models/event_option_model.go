package models

import (
	"gorm.io/gorm"
)

type EventOption struct {
	gorm.Model
	EventID     uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string  `gorm:"size:100;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"default:0.0"`
	Stock       int     `gorm:"default:0"`
}
