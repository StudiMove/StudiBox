package auth_routes

import (
	authHandlers "backend/core/api/handlers/auth"
	"backend/core/services/auth"
	"backend/core/services/email"
	"backend/core/services/owner"
	"backend/core/services/user"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes enregistre les routes d'authentification avec authService
func AuthRoutes(
	routerGroup *gin.RouterGroup,
	authService *auth.AuthServiceType,
	userService *user.UserServiceType,
	ownerService *owner.OwnerServiceType,
	emailService *email.EmailServiceType,
) {
	authGroup := routerGroup.Group("/auth")
	{
		// Authentification classique
		authGroup.POST("/login", func(c *gin.Context) {
			authHandlers.HandleLogin(c, authService, userService, ownerService)
		})
		authGroup.POST("/register/user", func(c *gin.Context) {
			authHandlers.HandleRegisterUser(c, authService, userService, ownerService, emailService)
		})
		authGroup.POST("/register/owner", func(c *gin.Context) {
			authHandlers.HandleRegisterOwner(c, authService, userService, ownerService, emailService)
		})

		// Routes pour la réinitialisation de mot de passe
		authGroup.POST("/password/reset-code", func(c *gin.Context) {
			authHandlers.HandleGetResetCode(c, userService, emailService)
		})
		authGroup.POST("/password/update", func(c *gin.Context) {
			authHandlers.HandleUpdatePasswordWithCode(c, authService, userService, emailService, ownerService)
		})
	}
}
