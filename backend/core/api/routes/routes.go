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

	// Créer un groupe pour l'API
	apiGroup := router.Group("/api")

	// Enregistrer les différentes routes, en passant authSvc, businessSvc et eventSvc
	auth_routes.RegisterAuthRoutes(apiGroup, authSvc)
	business_routes.RegisterBusinessRoutes(apiGroup, businessSvc, authSvc)
	user_routes.RegisterUserRoutes(apiGroup)
	event_routes.RegisterEventRoutes(apiGroup, eventSvc)

	return router
}
