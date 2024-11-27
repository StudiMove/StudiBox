package models

import "time"

type Referral struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ParrainID uint      `gorm:"not null" json:"parrainId"` // ID du parrain
	FilleulID uint      `gorm:"not null" json:"filleulId"` // ID du filleul
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
