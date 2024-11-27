package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model        // Inclut les champs ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"unique" json:"name"` // Nom du rôle (ex: 'admin', 'user')
}
