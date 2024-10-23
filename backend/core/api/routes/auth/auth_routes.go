package auth_routes

import (
	authHandlers "backend/core/api/handlers/auth"
	"backend/core/services/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes enregistre les routes d'authentification avec authService
func AuthRoutes(routerGroup *gin.RouterGroup, authService *auth.AuthService) {
	authGroup := routerGroup.Group("/auth")
	{
		// Passer authService aux handlers
		authGroup.POST("/login", func(c *gin.Context) {
			authHandlers.HandleLogin(c, authService)
		})
		authGroup.POST("/register/user", func(c *gin.Context) {
			authHandlers.HandleRegisterUser(c, authService)
		})
		authGroup.POST("/register/business", func(c *gin.Context) {
			authHandlers.HandleRegisterBusinessUser(c, authService)
		})
	}
}
