package user_routes

import (
	userHandlers "backend/core/api/handlers/user"
	"backend/core/api/middleware"
	"backend/core/services/auth"
	"backend/core/services/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routerGroup *gin.RouterGroup, userService *user.UserService, authService *auth.AuthService) {
	userGroup := routerGroup.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())

	{
		userGroup.GET("/profile", func(c *gin.Context) {
			userHandlers.HandleGetUserProfile(c, userService)
		})

		userGroup.PUT("/profile", func(c *gin.Context) {
			userHandlers.HandleUpdateUserProfile(c, userService)
		})

		// Route pour récupérer un utilisateur spécifique par ID
		userGroup.GET("/:id", func(c *gin.Context) {
			userHandlers.HandleGetUserByID(c, userService)
		})

		// Route pour récupérer tous les utilisateurs (Admin seulement)
		userGroup.GET("/all", middleware.RoleMiddleware(authService, []string{"Admin"}), func(c *gin.Context) {
			userHandlers.HandleGetAllUsers(c, userService)
		})

		// Route pour mettre à jour un utilisateur par ID (Admin seulement)
		userGroup.PUT("/:id", middleware.RoleMiddleware(authService, []string{"Admin"}), func(c *gin.Context) {
			userHandlers.HandleUpdateUserByID(c, userService)
		})
	}
}
