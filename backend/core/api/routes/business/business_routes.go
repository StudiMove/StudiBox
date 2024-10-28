package business_routes

import (
	businessHandlers "backend/core/api/handlers/business"
	"backend/core/api/middleware"
	"backend/core/services/auth"
	"backend/core/services/business"
	"backend/core/services/user"

	"github.com/gin-gonic/gin"
)

func BusinessRoutes(routerGroup *gin.RouterGroup, businessService *business.BusinessService, userService *user.UserService, authService *auth.AuthService) {
	businessGroup := routerGroup.Group("/business")
	businessGroup.Use(middleware.AuthMiddleware())

	{
		// Route pour récupérer le profil de l'utilisateur connecté
		businessGroup.GET("/profile", func(c *gin.Context) {
			businessHandlers.HandleGetBusiness(c, businessService, userService)
		})

		// Route pour mettre à jour le profil de l'utilisateur connecté ou un business par ID
		businessGroup.PUT("/profile", func(c *gin.Context) {
			businessHandlers.HandleUpdateBusiness(c, businessService, userService)
		})

		// Route pour récupérer un business spécifique par ID
		businessGroup.GET("/:id", func(c *gin.Context) {
			businessHandlers.HandleGetBusiness(c, businessService, userService)
		})

		// Route pour récupérer tous les businesses
		businessGroup.GET("/all", func(c *gin.Context) {
			businessHandlers.HandleGetAllBusinesses(c, businessService)
		})

		// Route pour mettre à jour un business par ID (Admin seulement)
		businessGroup.PUT("/:id", middleware.RoleMiddleware(authService, []string{"Business"}), func(c *gin.Context) {
			businessHandlers.HandleUpdateBusiness(c, businessService, userService)
		})
	}
}
