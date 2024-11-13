package models

import (
	"gorm.io/gorm"
)

type EventCategory struct {
	gorm.Model
	Name      string  `gorm:"size:50;unique;not null"`
	LikeCount int     `gorm:"default:0"`
	ViewCount int     `gorm:"default:0"`
	Events    []Event `gorm:"many2many:event_event_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
