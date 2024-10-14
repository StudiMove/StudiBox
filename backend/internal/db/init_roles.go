package db

import (
    "backend/internal/db/models" // Assurez-vous que ce chemin est correct
    "gorm.io/gorm"
)

func InitRoles(db *gorm.DB) {
    // Créer les rôles
    roles := []models.Role{
        {Name: "admin"},
        {Name: "business"},
        {Name: "user"},
    }

    // Vérifier si les rôles existent déjà
    for _, role := range roles {
        var existingRole models.Role
        // Vérifiez si le rôle existe déjà
        if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
            // Si le rôle n'existe pas, l'ajouter
            db.Create(&role)
        }
    }

    // Ajouter l'entrée dans user_roles pour lier user_id 8 à role_id 1
    var userRole models.UserRole
    userRole.UserID = 8
    userRole.RoleID = 1

    // Vérifier si l'association existe déjà
    if err := db.Where("user_id = ? AND role_id = ?", userRole.UserID, userRole.RoleID).First(&userRole).Error; err != nil {
        // Si l'association n'existe pas, l'ajouter
        db.Create(&userRole)
    }
}
