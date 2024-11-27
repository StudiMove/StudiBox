package models

import (
	"time"

	"gorm.io/gorm"
)

type PasswordReset struct {
	gorm.Model
	UserID     uint      `json:"userId"`
	ResetCode  int       `json:"resetCode"`
	Expiration time.Time `json:"expiration"`
}
