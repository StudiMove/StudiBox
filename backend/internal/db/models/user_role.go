package models

import "gorm.io/gorm"

type UserRole struct {
    gorm.Model
    UserID uint `gorm:"uniqueIndex:user_role_unique" json:"userId"` // Référence à l'utilisateur
    RoleID uint `gorm:"uniqueIndex:user_role_unique" json:"roleId"` // Référence au rôle
}
