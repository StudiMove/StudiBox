package auth

import (
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	"backend/core/services/auth"
	"backend/core/services/email"
	"backend/core/services/owner"
	"backend/core/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleRegisterUser gère la création et la connexion d'un utilisateur standard
func HandleRegisterUser(
	c *gin.Context,
	authService *auth.AuthServiceType,
	userService *user.UserServiceType,
	ownerService *owner.OwnerServiceType,
	emailService *email.EmailServiceType,
) {
	var registerReq request.RegisterUserRequest

	// Validation des données entrantes
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données invalides", err))
		return
	}

	// Vérifier l'unicité de l'utilisateur
	if err := validateUniqueUser(c, userService, registerReq.Email, registerReq.Pseudo); err != nil {
		return
	}

	// Enregistrer l'utilisateur
	if err := authService.Register.RegisterUser(&registerReq); err != nil {
		handleRegistrationError(c, err)
		return
	}

	// Envoi de l'email de bienvenue après l'inscription
	emailData := map[string]string{
		"subject": "Bienvenue sur StudiMove !",
		"name":    registerReq.FirstName,
	}
	err := emailService.SendEmailWithTemplate(email.EventRegistration, []string{registerReq.Email}, emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de l'envoi de l'email", err))
		return
	}

	// Générer le token pour l'utilisateur
	token, err := authService.Login.Login(registerReq.Email, registerReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Échec de la connexion", err))
		return
	}

	// Générer la réponse d'authentification
	if err := buildAuthResponse(c, userService, ownerService, registerReq.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la génération de la réponse", err))
	}
}

// HandleRegisterOwner gère la création et la connexion d'un utilisateur propriétaire (Owner)
func HandleRegisterOwner(
	c *gin.Context,
	authService *auth.AuthServiceType,
	userService *user.UserServiceType,
	ownerService *owner.OwnerServiceType,
	emailService *email.EmailServiceType,
) {
	var registerReq request.RegisterOwnerRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Données invalides", err))
		return
	}

	if err := validateUniqueUser(c, userService, registerReq.Email, registerReq.Pseudo); err != nil {
		return
	}

	if err := authService.Register.RegisterOwnerUser(&registerReq); err != nil {
		handleRegistrationError(c, err)
		return
	}

	// Envoi de l'email de bienvenue après l'inscription
	emailData := map[string]string{
		"subject": "Bienvenue sur StudiMove !",
		"name":    registerReq.FirstName,
	}
	err := emailService.SendEmailWithTemplate(email.EventRegistration, []string{registerReq.Email}, emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de l'envoi de l'email", err))
		return
	}

	token, err := authService.Login.Login(registerReq.Email, registerReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Échec de la connexion", err))
		return
	}

	// Appel de buildAuthResponse sans authService
	if err := buildAuthResponse(c, userService, ownerService, registerReq.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Échec de la génération de la réponse après l'inscription", err))
	}
}
