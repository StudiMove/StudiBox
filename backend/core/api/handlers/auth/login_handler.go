package auth

import (
	"backend/config"
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/services/auth"
	"backend/core/services/user"
	"backend/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleLogin gère la connexion des utilisateurs
func HandleLogin(c *gin.Context, authService *auth.AuthService, userService *user.UserService) {
	var loginReq request.LoginRequest

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

	// Récupérer les informations de l'utilisateur
	user, err := userService.Retrieval.GetUserByEmail(loginReq.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user data", err))
		return
	}

	// Extraire les noms des rôles via le service
	roleNames, err := userService.Management.ExtractRoleNames(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user roles", err))
		return
	}

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

func HandleFirebaseLogin(c *gin.Context, authService *auth.AuthService) {
	var loginReq struct {
		IDToken string `json:"idToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	token, err := config.VerifyIDToken(loginReq.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Invalid Firebase ID token", err))
		return
	}

	userEmail := token.Claims["email"].(string)
	firstName := token.Claims["given_name"].(string)
	lastName := token.Claims["family_name"].(string)

	user, err := authService.GetOrCreateUserByEmail(userEmail, firstName, lastName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve or create user", err))
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, config.AppConfig.JwtSecretAccessKey, "StudiMove", "studi_users", 72)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to generate token", err))
		return
	}

	resp := response.LoginResponse{
		Token:           jwtToken,
		IsAuthenticated: true,
		User: response.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Login successful", resp))
}
