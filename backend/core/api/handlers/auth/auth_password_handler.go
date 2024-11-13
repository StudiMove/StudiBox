package auth

import (
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	"backend/core/services/auth"
	"backend/core/services/email"
	"backend/core/services/owner"
	"backend/core/services/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler pour envoyer le code de réinitialisation
func HandleGetResetCode(c *gin.Context, userService *user.UserServiceType, emailService *email.EmailServiceType) {
	var req request.GetResetCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Email invalide", err))
		return
	}

	resetCode, err := userService.UserPassword.GenerateResetCode(req.Email)
	if err != nil {
		if err.Error() == "utilisateur introuvable" {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Utilisateur non trouvé", err))
		} else {
			c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		}
		return
	}

	data := map[string]string{
		"code": fmt.Sprintf("%06d", resetCode),
	}

	if err := emailService.SendEmailWithTemplate(email.PasswordReset, []string{req.Email}, data); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de l'envoi de l'email", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code envoyé par email"})
}

// Handler pour mettre à jour le mot de passe avec le code
func HandleUpdatePasswordWithCode(
	c *gin.Context,
	authService *auth.AuthServiceType,
	userService *user.UserServiceType,
	emailService *email.EmailServiceType,
	ownerService *owner.OwnerServiceType,
) {
	var req request.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Données invalides",
			"data":    err.Error(),
		})
		return
	}

	// Vérifier que les champs ne sont pas vides
	if req.Code == 0 || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Code ou nouveau mot de passe manquant", nil))
		return
	}

	// Réinitialiser le mot de passe
	err := userService.UserPassword.UpdatePasswordWithCode(req.Email, req.Code, req.NewPassword)
	if err != nil {
		if err.Error() == "code de réinitialisation invalide ou expiré" {
			c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Code invalide ou expiré", err))
		} else if err.Error() == "utilisateur introuvable" {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Utilisateur non trouvé", err))
		} else {
			c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		}
		return
	}

	// Connexion automatique après réinitialisation
	token, err := authService.Login.Login(req.Email, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Échec de la connexion après réinitialisation", err))
		return
	}

	if err := buildAuthResponse(c, userService, ownerService, req.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la génération de la réponse", err))
		return
	}
}
