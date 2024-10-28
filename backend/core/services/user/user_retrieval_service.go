package user

import (
	"backend/core/models"
	stores "backend/core/stores/user"
	"backend/core/utils"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRetrievalService struct {
	store *stores.UserStore
}

// NewUserRetrievalService crée une nouvelle instance de UserRetrievalService
func NewUserRetrievalService(store *stores.UserStore) *UserRetrievalService {
	return &UserRetrievalService{
		store: store,
	}
}

// GetUserByID récupère un utilisateur par ID
func (s *UserRetrievalService) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error retrieving user by ID: " + err.Error())
	}
	return user, nil
}

// GetAllUsers récupère tous les utilisateurs
func (s *UserRetrievalService) GetAllUsers() ([]models.User, error) {
	users, err := s.store.GetAll()
	if err != nil {
		return nil, errors.New("unable to retrieve users: " + err.Error())
	}
	return users, nil
}

// GetUserByEmail récupère un utilisateur par email
func (s *UserRetrievalService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found with the provided email")
		}
		return nil, errors.New("error retrieving user by email: " + err.Error())
	}
	return user, nil
}

// GetUserIDFromRequest récupère l'ID d'un utilisateur à partir des claims JWT ou d'un paramètre ID
func (s *UserRetrievalService) GetUserIDFromRequest(c *gin.Context) (uint, error) {
	// Vérifier si un ID est passé dans l'URL
	idParam := c.Param("id")
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil || id <= 0 {
			return 0, errors.New("invalid ID parameter")
		}
		return uint(id), nil
	}

	// Si aucun ID dans l'URL, récupérer l'ID de l'utilisateur connecté via JWT
	userClaims, exists := c.Get("user")
	if !exists {
		return 0, errors.New("missing or invalid token")
	}

	// Convertir les claims JWT en type structuré
	claims, ok := userClaims.(*utils.JWTClaims)
	if !ok {
		return 0, errors.New("invalid token structure")
	}
	return claims.UserID, nil
}
