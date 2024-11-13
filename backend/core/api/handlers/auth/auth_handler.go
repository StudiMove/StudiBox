package auth

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/services/owner"
	"backend/core/services/user"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// buildAuthResponse construit la réponse d'authentification pour un utilisateur
func buildAuthResponse(c *gin.Context, userService *user.UserServiceType, ownerService *owner.OwnerServiceType, email string, token string) error {
	user, err := userService.Retrieval.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la récupération de l'utilisateur", err))
		return err
	}

	owner, _ := ownerService.Retrieval.GetOwnerByUserEmail(email)
	ownerType := "standard"
	if owner != nil {
		ownerType = owner.Type
	}

	roleName, err := userService.Management.ExtractRoleName(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la récupération du rôle utilisateur", err))
		return err
	}

	resp := response.LoginResponse{
		Token:           token,
		IsAuthenticated: true,
		User: response.UserAuthResponse{
			ID:          user.ID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Pseudo:      user.Pseudo,
			ProfileType: user.ProfileType,
			Type:        ownerType,
			Role:        roleName,
		},
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Opération d'authentification réussie", resp))
	return nil
}

// validateUniqueUser vérifie l'unicité de l'email et du pseudo pour éviter les doublons
func validateUniqueUser(c *gin.Context, userService *user.UserServiceType, email, pseudo string) error {
	if existingUser, _ := userService.Retrieval.GetUserByEmail(email); existingUser != nil {
		c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("L'email est déjà utilisé", nil))
		return errors.New("email already used")
	}

	if existingUser, _ := userService.Retrieval.GetUserByPseudo(pseudo); existingUser != nil {
		c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("Le pseudo est déjà utilisé", nil))
		return errors.New("pseudo already used")
	}

	return nil
}

// handleRegistrationError gère les erreurs d'inscription de manière plus détaillée
func handleRegistrationError(c *gin.Context, err error) {
	// Analyse des messages d'erreur et génération de réponses spécifiques
	switch {
	case strings.Contains(err.Error(), "email déjà utilisé"):
		c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("L'email est déjà utilisé", err))
	case strings.Contains(err.Error(), "le pseudo est déjà utilisé"):
		c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("Le pseudo est déjà utilisé", err))
	case strings.Contains(err.Error(), "erreur lors de la vérification"):
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Erreur de vérification des données", err))
	case strings.Contains(err.Error(), "constraint") && strings.Contains(err.Error(), "unique"):
		c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("Contrainte d'unicité violée", err))
	case strings.Contains(err.Error(), "validation failed"):
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Échec de la validation des données", err))
	default:
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de l'enregistrement", err))
	}
}
