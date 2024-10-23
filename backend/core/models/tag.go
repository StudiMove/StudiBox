package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string `gorm:"unique;not null"` // Un nom unique pour chaque tag
	LikeCount int    `gorm:"default:0"`       // Le nombre de likes associés à ce tag
	ViewCount int    `gorm:"default:0"`       // Le nombre de vues associées à ce tag
}
