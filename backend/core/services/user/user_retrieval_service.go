package user

import (
	stores "backend/core/stores/user"
	"backend/core/utils"
	"backend/database/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRetrievalServiceType struct {
	store *stores.UserStoreType
}

// NewUserRetrievalService crée une nouvelle instance de UserRetrievalService
func UserRetrievalService(store *stores.UserStoreType) *UserRetrievalServiceType {
	return &UserRetrievalServiceType{
		store: store,
	}
}

// GetUserByID récupère un utilisateur par ID
func (s *UserRetrievalServiceType) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.store.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("error retrieving user by ID: " + err.Error())
	}
	return user, nil
}

// GetUserByEmail récupère un utilisateur par email
func (s *UserRetrievalServiceType) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New("error retrieving user by email: " + err.Error())
	}
	return user, nil
}

// GetUserByPseudo récupère un utilisateur par son pseudo
func (s *UserRetrievalServiceType) GetUserByPseudo(pseudo string) (*models.User, error) {
	user, err := s.store.GetByPseudo(pseudo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found with the provided pseudo")
		}
		return nil, errors.New("error retrieving user by pseudo: " + err.Error())
	}
	return user, nil
}

// GetUserFromClaims récupère un utilisateur à partir des claims JWT
func (s *UserRetrievalServiceType) GetUserFromClaims(c *gin.Context) (*models.User, error) {
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		return nil, errors.New("failed to retrieve claims from token")
	}
	return s.GetUserByID(claims.UserID)
}

// GetAllUsers récupère tous les utilisateurs
func (s *UserRetrievalServiceType) GetAllUsers() ([]models.User, error) {
	users, err := s.store.GetAll()
	if err != nil {
		return nil, errors.New("unable to retrieve users: " + err.Error())
	}
	return users, nil
}

// GetUserIDFromRequest récupère l'ID d'un utilisateur à partir des claims JWT ou d'un paramètre ID
func (s *UserRetrievalServiceType) GetUserIDFromRequest(c *gin.Context) (uint, error) {
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

func (s *UserRetrievalServiceType) GetOrCreateUserByEmail(email, firstName, lastName string) (*models.User, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to retrieve user by email: %w", err)
	}

	if user == nil {
		user = &models.User{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		}
		if err := s.store.Create(user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	return user, nil
}
