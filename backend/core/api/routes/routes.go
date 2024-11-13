package routes

import (
	"backend/core/api/middleware"
	auth_routes "backend/core/api/routes/auth"
	event_routes "backend/core/api/routes/event"
	owner_routes "backend/core/api/routes/owner"
	user_routes "backend/core/api/routes/user"
	authService "backend/core/services/auth"
	emailService "backend/core/services/email"
	eventService "backend/core/services/event"
	ownerService "backend/core/services/owner"
	userService "backend/core/services/user"
	"backend/database"

	"github.com/gin-gonic/gin"
)

// SetupRouter configure toutes les routes de l'application
func SetupRouter(router *gin.Engine) *gin.Engine {
	// Utiliser le middleware CORS ici, pas besoin de le répéter ailleurs
	router.Use(middleware.CORSMiddleware())

	// Charger la configuration depuis utils

	// Initialiser le service d'authentification avec la base de données
	authSvc := authService.AuthService(database.DB)

	// Initialiser le service des événements avec la base de données directement
	eventSvc := eventService.EventService(database.DB)

	// Initialiser le service des utilisateurs professionnels
	ownersSvc := ownerService.OwnerService(database.DB)

	// Initialiser le service des utilisateurs
	userSvc := userService.UserService(database.DB)

	// Initialiser le service d'envoi d'email
	emailSvc := emailService.EmailService()

	// Créer un groupe pour l'API
	apiGroup := router.Group("/api")

	// Enregistrer les différentes routes
	auth_routes.AuthRoutes(apiGroup, authSvc, userSvc, ownersSvc, emailSvc)
	owner_routes.OwnerRoutes(apiGroup, ownersSvc, userSvc)
	user_routes.UserRoutes(apiGroup, userSvc)
	event_routes.EventRoutes(apiGroup, eventSvc, ownersSvc, userSvc)

	return router
}
