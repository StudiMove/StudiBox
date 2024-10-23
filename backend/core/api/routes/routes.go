package routes

import (
	"backend/core/api/middleware"
	auth_routes "backend/core/api/routes/auth"
	business_routes "backend/core/api/routes/business"
	event_routes "backend/core/api/routes/event"
	user_routes "backend/core/api/routes/user"
	authService "backend/core/services/auth"
	businessService "backend/core/services/business"
	eventService "backend/core/services/event"
	userService "backend/core/services/user"
	"backend/database"

	"github.com/gin-gonic/gin"
)

// SetupRouter configure toutes les routes de l'application
func SetupRouter(router *gin.Engine) *gin.Engine {
	// Utiliser le middleware CORS ici, pas besoin de le répéter ailleurs
	router.Use(middleware.CORSMiddleware())

	// Initialiser le service d'authentification avec la base de données
	authSvc := authService.NewAuthService(database.DB)

	// Initialiser le service des événements avec la base de données directement
	eventSvc := eventService.NewEventService(database.DB)

	// Initialiser le service des utilisateurs professionnels
	businessSvc := businessService.NewBusinessService(database.DB)

	// Initialiser le service des utilisateurs
	userSvc := userService.NewUserService(database.DB)

	// Créer un groupe pour l'API
	apiGroup := router.Group("/api")

	// Enregistrer les différentes routes, en passant authSvc, businessSvc, userSvc et eventSvc
	auth_routes.AuthRoutes(apiGroup, authSvc)
	business_routes.BusinessRoutes(apiGroup, businessSvc, userSvc, authSvc)
	user_routes.UserRoutes(apiGroup, userSvc, authSvc)
	event_routes.EventRoutes(apiGroup, eventSvc)

	return router
}
