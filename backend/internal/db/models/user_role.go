package models

import "gorm.io/gorm"

type UserRole struct {
    gorm.Model // Inclut les champs ID, CreatedAt, UpdatedAt, DeletedAt
    UserID uint // Référence à l'utilisateur
    RoleID uint // Référence au rôle
}
