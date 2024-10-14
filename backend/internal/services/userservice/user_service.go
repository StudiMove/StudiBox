// backend/internal/services/userservice/user_service.go
package userservice

import (
    "backend/internal/db/models"
    "gorm.io/gorm"
)

// UserService représente le service pour gérer les utilisateurs.
type UserService struct {
    db *gorm.DB // Instance de Gorm pour interagir avec la base de données.
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

// GetUserRolesByID récupère les rôles d'un utilisateur à partir de son ID.
func (s *UserService) GetUserRolesByID(userID uint) ([]models.Role, error) {
    var user models.User

    // Utilise Preload pour charger les rôles associés à l'utilisateur.
    if err := s.db.Preload("Roles").First(&user, userID).Error; err != nil {
        return nil, err // Retourne l'erreur si la récupération échoue.
    }

    return user.Roles, nil // Retourne les rôles de l'utilisateur.
}
