package business

import (
	"backend/core/models"
	"backend/core/stores"
	"errors"

	"gorm.io/gorm"
)

// UpdateProfileInput représente les champs pour la mise à jour du profil business
type UpdateProfileInput struct {
	CompanyName string
	Address     string
	City        string
	Postcode    string
	Country     string
	Phone       string
}

// BusinessService struct, centralise les opérations liées aux utilisateurs professionnels
type BusinessService struct {
	db *gorm.DB
}

// NewBusinessService crée une nouvelle instance de BusinessService
func NewBusinessService(db *gorm.DB) *BusinessService {
	return &BusinessService{
		db: db,
	}
}

// GetBusinessUserByID récupère un utilisateur business par son ID
func (s *BusinessService) GetBusinessUserByID(userID uint) (*models.BusinessUser, error) {
	businessUser, err := stores.NewBusinessUserStore(s.db).GetByID(userID) // Utiliser le store interne avec l'instance de DB
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("business user not found")
		}
		return nil, err
	}
	return businessUser, nil
}

// GetAllBusinessUsers récupère tous les utilisateurs business
func (s *BusinessService) GetAllBusinessUsers() ([]models.BusinessUser, error) {
	businessUsers, err := stores.NewBusinessUserStore(s.db).GetAll() // Utiliser le store interne
	if err != nil {
		return nil, err
	}
	return businessUsers, nil
}

// UpdateBusinessUserProfile met à jour le profil d'un utilisateur business
func (s *BusinessService) UpdateBusinessUserProfile(userID uint, input UpdateProfileInput) error {
	businessUser, err := stores.NewBusinessUserStore(s.db).GetByID(userID) // Utiliser le store interne
	if err != nil {
		return err
	}

	// Mettre à jour uniquement les champs non vides
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
	if input.Phone != "" {
		businessUser.User.Phone = input.Phone // Mise à jour du champ Phone dans User
	}

	// Enregistre les modifications
	return stores.NewBusinessUserStore(s.db).Update(businessUser)
}
