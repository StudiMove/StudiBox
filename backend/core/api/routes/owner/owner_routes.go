package owner_routes

import (
	ownerHandlers "backend/core/api/handlers/owner"
	"backend/core/api/middleware"
	"backend/core/services/owner"
	"backend/core/services/user"

	"github.com/gin-gonic/gin"
)

func OwnerRoutes(routerGroup *gin.RouterGroup, ownerService *owner.OwnerServiceType, userService *user.UserServiceType) {
	ownerGroup := routerGroup.Group("/owner")
	ownerGroup.Use(middleware.AuthMiddleware())

	{
		// Route pour récupérer le profil de l'utilisateur connecté
		ownerGroup.GET("/profile", func(c *gin.Context) {
			ownerHandlers.HandleGetOwner(c, ownerService, userService)
		})

		// Route pour mettre à jour le profil de l'utilisateur connecté
		ownerGroup.PUT("/profile", func(c *gin.Context) {
			ownerHandlers.HandleUpdateOwner(c, ownerService, userService)
		})

		// Route pour récupérer un owner spécifique par ID
		ownerGroup.GET("/:id", func(c *gin.Context) {
			ownerHandlers.HandleGetOwner(c, ownerService, userService)
		})

		// Route pour récupérer tous les owners
		ownerGroup.GET("/all", func(c *gin.Context) {
			ownerHandlers.HandleGetAllOwners(c, ownerService)
		})

		// Route pour mettre à jour un owner par ID (Admin ou rôle spécifique seulement)
		ownerGroup.PUT("/:id", middleware.RoleMiddleware(userService, []string{"Owner"}), func(c *gin.Context) {
			ownerHandlers.HandleUpdateOwner(c, ownerService, userService)
		})
	}
}
