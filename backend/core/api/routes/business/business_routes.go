package business_routes

import (
	businessHandlers "backend/core/api/handlers/business"
	"backend/core/api/middleware"
	"backend/core/services/auth"
	"backend/core/services/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBusinessRoutes(routerGroup *gin.RouterGroup, businessService *business.BusinessService, authService *auth.AuthService) {
	businessGroup := routerGroup.Group("/business")
	businessGroup.Use(middleware.AuthMiddleware())

	{
		businessGroup.GET("/profile", func(c *gin.Context) {
			businessHandlers.HandleGetBusinessProfile(c, businessService)
		})

		businessGroup.PUT("/profile", func(c *gin.Context) {
			businessHandlers.HandleUpdateBusinessProfile(c, businessService)
		})

		// Route pour récupérer un business spécifique par ID
		businessGroup.GET("/:id", func(c *gin.Context) {
			businessHandlers.HandleGetBusinessByID(c, businessService)
		})

		// Route pour récupérer tous les businesses
		businessGroup.GET("/all", func(c *gin.Context) {
			businessHandlers.HandleGetAllBusinesses(c, businessService)
		})

		// Route pour mettre à jour un business par ID (seulement pour les Admins)
		businessGroup.PUT("/:id", func(c *gin.Context) {
			requiredRoles := []string{"Admin"}
			middleware.RoleMiddleware(authService, requiredRoles, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				businessHandlers.HandleUpdateBusinessByID(c, businessService)
			}))
		})
	}
}
