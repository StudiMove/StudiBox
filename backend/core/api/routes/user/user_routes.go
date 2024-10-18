package user_routes

import (
	"backend/core/api/handlers/user"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes enregistre les routes liées aux utilisateurs
func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/users")
	{
		userGroup.GET("/:id", user.HandleGetUserByID)
		userGroup.PUT("/:id", user.HandleUpdateUser)
	}
}
