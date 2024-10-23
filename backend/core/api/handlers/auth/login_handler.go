package auth

import (
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/services/auth"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleLogin gère la connexion des utilisateurs
func HandleLogin(c *gin.Context, authService *auth.AuthService, userService *user.UserService) {
	var loginReq request.LoginRequest

	// Validation des données d'entrée
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	// Tenter la connexion
	token, err := authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Invalid email or password", err))
		return
	}

	// Récupérer les informations de l'utilisateur pour les inclure dans la réponse
	user, err := userService.Retrieval.GetUserByEmail(loginReq.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user data", err))
		return
	}

	// Extraire les noms des rôles via le service
	roleNames := authService.ExtractRoleNames(user.Roles)

	// Créer la réponse
	resp := response.LoginResponse{
		Token:           token,
		IsAuthenticated: true,
		User: response.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Roles:     roleNames,
		},
	}
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Login successful", resp))
}
