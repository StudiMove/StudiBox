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

		userGroup.GET("/:id", func(c *gin.Context) {
			userHandlers.HandleGetUser(c, userService)
		})

		// Les autres routes restent inchangées
		userGroup.GET("/profile", func(c *gin.Context) {
			userHandlers.HandleGetUser(c, userService)
		})

		// Utilisation de HandleUpdateUser pour les deux cas
		userGroup.PUT("/profile", func(c *gin.Context) {
			userHandlers.HandleUpdateUser(c, userService)
		})

		userGroup.GET("/all", middleware.RoleMiddleware(authService, []string{"Business"}), func(c *gin.Context) {
			userHandlers.HandleGetAllUsers(c, userService)
		})

		userGroup.GET("/export/all", func(c *gin.Context) {
			userHandlers.HandleExportUsersCSV(c, userService)
		})

		userGroup.PUT("/:id", middleware.RoleMiddleware(authService, []string{"Business"}), func(c *gin.Context) {
			userHandlers.HandleUpdateUser(c, userService)
		})
	}
}
