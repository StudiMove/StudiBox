package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name      string `gorm:"unique;not null"` // Un nom unique pour chaque catégorie
	LikeCount int    `gorm:"default:0"`       // Le nombre de likes associés à cette catégorie
	ViewCount int    `gorm:"default:0"`       // Le nombre de vues associées à cette catégorie
}
