package routes

import (
	"backend/internal/api/handlers/events"
	"backend/internal/api/middleware"
	"backend/internal/services/auth"
	"backend/internal/services/event"
	"backend/internal/services/storage"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterEventRoutes enregistre les routes pour la gestion des événements
func RegisterEventRoutes(router *mux.Router, db *gorm.DB, storageService storage.StorageService, authService *auth.AuthService) {
	eventService := event.NewEventService(db, storageService)     // Initialise le service événement
	eventUserActionService := event.NewEventUserActionService(db) // Nouveau service

	// Créer une instance du handler pour les événements
	createEventHandler := events.NewCreateEventHandler(eventService) // Créez votre handler ici

	// Définir les rôles requis pour l'accès aux routes d'événements
	requiredRoles := []string{"admin", "business", "school", "association"}

	// Route pour créer un nouvel événement
	router.Handle("/create",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(createEventHandler.HandleCreateEvent)), // Utilisez le handler
		),
	).Methods("POST", "OPTIONS")

	// Route pour mettre à jour un événement existant
	router.Handle("/update",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				updateHandler := events.NewUpdateEventHandler(eventService) // Créer le handler pour l'update
				updateHandler.HandleUpdateEvent(w, r)                       // Appeler le handler pour traiter la requête
			})),
		),
	).Methods("PUT", "OPTIONS")

	// Route pour obtenir les détails d'un événement
	router.Handle("/get",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				getHandler := events.NewGetEventHandler(eventService)
				getHandler.HandleGetEvent(w, r)
			})),
		),
	).Methods("GET", "POST", "OPTIONS")

	// Route pour récupérer tous les événements
	router.Handle("/all",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				allEventsHandler := events.NewGetAllEventsHandler(eventService)
				allEventsHandler.HandleGetAllEvents(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Route pour récupérer uniquement les événements en ligne
	router.Handle("/all/online",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				onlineEventsHandler := events.NewGetOnlineEventsHandler(eventService)
				onlineEventsHandler.HandleGetOnlineEvents(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Nouvelle route pour récupérer les événements d'un utilisateur spécifique (utilisant son UserID à partir du token)
	router.Handle("/list",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				getEventListHandler := events.NewGetEventListHandler(eventService)
				getEventListHandler.HandleGetEventList(w, r)
			})),
		),
	).Methods("GET", "OPTIONS")

	// Nouvelle route pour récupérer les événements d'un utilisateur cible spécifique (UserTargetID passé dans le body)
	router.Handle("/list/target",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				getEventListByTargetIDHandler := events.NewGetEventListByTargetIDHandler(eventService)
				getEventListByTargetIDHandler.HandleGetEventListByTargetID(w, r)
			})),
		),
	).Methods("POST", "OPTIONS")

	categoriesHandler := events.NewGetAllCategoriesHandler(eventService)
	tagsHandler := events.NewGetAllTagsHandler(eventService)

	// Route pour récupérer toutes les catégories
	router.Handle("/categories",
		middleware.AuthMiddleware(http.HandlerFunc(categoriesHandler.HandleGetAllCategories)),
	).Methods("GET", "OPTIONS")

	// Route pour récupérer tous les tags
	router.Handle("/tags",
		middleware.AuthMiddleware(http.HandlerFunc(tagsHandler.HandleGetAllTags)),
	).Methods("GET", "OPTIONS")

	uploadEventImageHandler := events.NewUploadEventImageHandler(eventService)
	router.Handle("/upload-image",
		middleware.AuthMiddleware(
			middleware.RoleMiddleware(authService, requiredRoles)(http.HandlerFunc(uploadEventImageHandler.HandleUploadEventImage)),
		),
	).Methods("POST", "OPTIONS")

	// Route pour récupérer tous les événements (sans middleware)
	router.HandleFunc("/all/events", func(w http.ResponseWriter, r *http.Request) {
		allEventsHandler := events.NewGetAllEventsMobileHandler(eventService)
		allEventsHandler.HandleGetAllEventsMobile(w, r)
	}).Methods("GET", "OPTIONS")

	// APP MOBILE

	eventUserActionHandler := events.NewEventUserActionHandler(eventUserActionService)

	// Nouvelle route pour mettre à jour les actions utilisateur (intéressé ou favori)
	router.Handle("/user/actions",
		middleware.AuthMiddleware(http.HandlerFunc(eventUserActionHandler.UpdateUserAction)),
	).Methods("POST", "OPTIONS")

	// Nouvelle route pour récupérer les événements favoris d'un utilisateur
	router.Handle("/user/favorites",
		middleware.AuthMiddleware(http.HandlerFunc(eventUserActionHandler.GetFavoriteEvents)),
	).Methods("GET", "OPTIONS")

	// Nouvelle route pour récupérer les événements où l'utilisateur est intéressé
	router.Handle("/user/interested",
		middleware.AuthMiddleware(http.HandlerFunc(eventUserActionHandler.GetInterestedEvents)),
	).Methods("GET", "OPTIONS")

	// Nouvelle route pour supprimer une action utilisateur
	router.Handle("/user/actions/remove",
		middleware.AuthMiddleware(http.HandlerFunc(eventUserActionHandler.RemoveUserAction)),
	).Methods("DELETE", "OPTIONS")

}
