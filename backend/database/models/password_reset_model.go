package models

import (
	"time"

	"gorm.io/gorm"
)

type PasswordReset struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	ResetCode  int  `gorm:"size:100;not null"`
	Expiration time.Time
}
