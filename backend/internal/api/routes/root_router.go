package routes

import (
	"backend/internal/api/middleware"
	"backend/internal/services/auth"
	"backend/internal/services/event"
	"backend/internal/services/organisation"
	"backend/internal/services/password"
	"backend/internal/services/profilservice"
	"backend/internal/services/referral"
	"backend/internal/services/storage"
	"backend/internal/services/userservice"
	"backend/pkg/httpclient"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitRouter initialise le routeur principal avec les versions V1 et V2 de l'API
func InitRouter(authService *auth.AuthService, eventService *event.EventService, apiClient *httpclient.APIClient, passwordService *password.PasswordResetService, profilService *profilservice.ProfilService, userService *userservice.UserService, storageService storage.StorageService, db *gorm.DB, jwtSecret string) *mux.Router {
	router := mux.NewRouter()

	// Middleware global pour le logging de chaque requête
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware) // Pour voir chaque requête reçue

	// V1 Routes
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	registerV1Routes(v1Router, authService, eventService, apiClient, passwordService, profilService, userService, storageService, db, jwtSecret)

	// V2 Routes (configuration vide pour le moment)
	v2Router := router.PathPrefix("/api/v2").Subrouter()
	registerV2Routes(v2Router, authService, eventService, apiClient)

	return router
}

// Fonction pour enregistrer toutes les routes de la version 1
func registerV1Routes(v1Router *mux.Router, authService *auth.AuthService, eventService *event.EventService, apiClient *httpclient.APIClient, passwordService *password.PasswordResetService, profilService *profilservice.ProfilService, userService *userservice.UserService, storageService storage.StorageService, db *gorm.DB, jwtSecret string) {
	// Routes d'authentification pour V1
	authRouter := v1Router.PathPrefix("/auth").Subrouter()
	RegisterAuthRoutes(authRouter, authService, apiClient)

	referralService := referral.NewReferralService(db)
	referralRouter := v1Router.PathPrefix("/referrals").Subrouter()
	RegisterReferralRoutes(referralRouter, referralService)

	// Routes de profil avec AuthMiddleware et RoleMiddleware
	profilRouter := v1Router.PathPrefix("/profil").Subrouter()
	RegisterProfilRoutes(profilRouter, profilService, authService, userService, jwtSecret)

	// Routes API protégées
	apiRouter := v1Router.PathPrefix("").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware)

	// Routes de gestion des événements avec protection par rôles
	eventRouter := apiRouter.PathPrefix("/events").Subrouter()
	RegisterEventRoutes(eventRouter, db, storageService, authService)

	// Routes de gestion des organisations avec protection par rôles
	orgService := organisation.NewOrganisationService(db) // Initialise le service organisation
	orgRouter := apiRouter.PathPrefix("/organisations").Subrouter()
	RegisterOrganisationRoutes(orgRouter, orgService, authService)

	// Routes de gestion des mots de passe
	passwordRouter := v1Router.PathPrefix("/password").Subrouter()
	RegisterPasswordRoutes(passwordRouter, passwordService, authService, db)
}

// Fonction pour enregistrer toutes les routes de la version 2 (actuellement vide)
func registerV2Routes(v2Router *mux.Router, authService *auth.AuthService, eventService *event.EventService, apiClient *httpclient.APIClient) {
	// Configuration pour les routes de la version 2 si nécessaire
	// Actuellement vide, ajoutez des routes spécifiques ici quand elles seront prêtes
}
