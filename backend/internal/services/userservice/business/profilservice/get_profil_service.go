package profilservice

import (
    "backend/internal/db/models"
    "errors"
    "gorm.io/gorm"
)

// GetBusinessUserByID récupère les informations de profil de l'utilisateur business via son ID
func (s *ProfilService) GetBusinessUserByID(userID uint) (*models.BusinessUser, error) {
    var businessUser models.BusinessUser
    // Utilisez Preload pour charger les rôles de l'utilisateur associé
    if err := s.DB.Preload("User.Roles").First(&businessUser, "user_id = ?", userID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("business user not found")
        }
        return nil, err
    }
    return &businessUser, nil
}
