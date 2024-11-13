package auth

import (
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/services/auth"
	"backend/core/services/owner"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleLogin gère la connexion des utilisateurs
func HandleLogin(c *gin.Context, authService *auth.AuthServiceType, userService *user.UserServiceType, ownerService *owner.OwnerServiceType) {
	var loginReq request.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données invalides", err))
		return
	}

	token, err := authService.Login.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Email ou mot de passe incorrect", err))
		return
	}

	// Appel de buildAuthResponse sans authService
	if err := buildAuthResponse(c, userService, ownerService, loginReq.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la génération de la réponse", err))
	}
}

// HandleFirebaseLogin gère la connexion des utilisateurs via Firebase
func HandleFirebaseLogin(c *gin.Context, authService *auth.AuthLoginServiceType) {
	var loginReq struct {
		IDToken string `json:"idToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données invalides", err))
		return
	}

	jwtToken, user, err := authService.FirebaseLogin(loginReq.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("ID token Firebase invalide", err))
		return
	}

	resp := response.LoginResponse{
		Token:           jwtToken,
		IsAuthenticated: true,
		User: response.UserAuthResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Connexion réussie", resp))
}
