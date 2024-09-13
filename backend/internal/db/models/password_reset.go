package models

import "gorm.io/gorm"

type PasswordReset struct {
    gorm.Model // Inclut les champs ID, CreatedAt, UpdatedAt, DeletedAt
    UserID    uint   // Référence à l'utilisateur
    ResetCode string // Code de réinitialisation
    Expiration string // Date d'expiration
}
