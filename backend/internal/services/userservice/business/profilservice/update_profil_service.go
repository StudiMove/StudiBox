// package profilservice

// import (
//     "backend/internal/db/models"
//     "gorm.io/gorm"
// )

// type UpdateProfileInput struct {
//     CompanyName string
//     Address     string
//     City        string
//     Postcode    string
//     Phone       string
//     Country     string
// }

// type ProfilService struct {
//     DB *gorm.DB
// }

// // UpdateBusinessUserProfile permet de mettre à jour les champs non vides.
// func (s *ProfilService) UpdateBusinessUserProfile(userID uint, input UpdateProfileInput) error {
//     var businessUser models.BusinessUser
//     var user models.User

//     // Récupérer l'utilisateur business par son ID
//     if err := s.DB.First(&businessUser, "user_id = ?", userID).Error; err != nil {
//         return err
//     }

//     // Récupérer les informations de l'utilisateur associé dans la table User
//     if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
//         return err
//     }

//     // Mettre à jour uniquement les champs non vides pour BusinessUser
//     if input.CompanyName != "" {
//         businessUser.CompanyName = input.CompanyName
//     }
//     if input.Address != "" {
//         businessUser.Address = input.Address
//     }
//     if input.City != "" {
//         businessUser.City = input.City
//     }
//     if input.Postcode != "" {
//         businessUser.Postcode = input.Postcode
//     }
//     if input.Country != "" {
//         businessUser.Country = input.Country
//     }

//     // Mettre à jour uniquement le champ Phone pour User s'il est fourni
//     if input.Phone != "" {
//         user.Phone = input.Phone
//     }

//     // Enregistrer les changements dans BusinessUser
//     if err := s.DB.Save(&businessUser).Error; err != nil {
//         return err
//     }

//     // Enregistrer les changements dans User
//     return s.DB.Save(&user).Error
// }
package profilservice

import "backend/internal/db/models"

type UpdateProfileInput struct {
    CompanyName string
    Address     string
    City        string
    Postcode    string
    Phone       string
    Country     string
}

// UpdateBusinessUserProfile permet de mettre à jour les champs non vides.
func (s *ProfilService) UpdateBusinessUserProfile(userID uint, input UpdateProfileInput) error {
    var businessUser models.BusinessUser
    var user models.User

    // Récupérer l'utilisateur business par son ID
    if err := s.DB.First(&businessUser, "user_id = ?", userID).Error; err != nil {
        return err
    }

    // Récupérer les informations de l'utilisateur associé dans la table User
    if err := s.DB.First(&user, "id = ?", userID).Error; err != nil {
        return err
    }

    // Mettre à jour uniquement les champs non vides pour BusinessUser
    if input.CompanyName != "" {
        businessUser.CompanyName = input.CompanyName
    }
    if input.Address != "" {
        businessUser.Address = input.Address
    }
    if input.City != "" {
        businessUser.City = input.City
    }
    if input.Postcode != "" {
        businessUser.Postcode = input.Postcode
    }
    if input.Country != "" {
        businessUser.Country = input.Country
    }

    // Mettre à jour uniquement le champ Phone pour User s'il est fourni
    if input.Phone != "" {
        user.Phone = input.Phone
    }

    // Enregistrer les changements dans BusinessUser
    if err := s.DB.Save(&businessUser).Error; err != nil {
        return err
    }

    // Enregistrer les changements dans User
    return s.DB.Save(&user).Error
}
