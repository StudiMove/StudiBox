package models

import (
	"gorm.io/gorm"
)

// EventLike représente une relation de "like" entre un utilisateur et un événement
type EventLike struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	EventID uint `gorm:"not null"`
}
