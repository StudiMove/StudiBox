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
func AuthRoutes(routerGroup *gin.RouterGroup, authService *auth.AuthServiceType, userService *user.UserServiceType, ownerService *owner.OwnerServiceType, emailService *email.EmailServiceType) {
	authGroup := routerGroup.Group("/auth")
	{
		// Passer authService aux handlers
		authGroup.POST("/login", func(c *gin.Context) {
			authHandlers.HandleLogin(c, authService, userService, ownerService)
		})
		authGroup.POST("/register/user", func(c *gin.Context) {
			authHandlers.HandleRegisterUser(c, authService, userService, ownerService, emailService)
		})
		authGroup.POST("/register/owner", func(c *gin.Context) {
			authHandlers.HandleRegisterOwner(c, authService, userService, ownerService, emailService)
		})
	}
}
