package models

import (
	"gorm.io/gorm"
)

type EventTarif struct {
	gorm.Model
	EventID     uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string  `gorm:"size:100;not null"`
	Price       float64 `gorm:"default:0.0"`
	Stock       int     `gorm:"default:0"`
	Description string  `gorm:"type:text"`
	Revenue     int     `gorm:"default:0"`
}
