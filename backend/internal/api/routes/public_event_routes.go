package routes

import (
	"backend/internal/api/handlers/events"
	"backend/internal/api/middleware"
	"backend/internal/services/event"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterPublicEventRoutes enregistre les routes d'événements publiques, accessibles sans authentification
func RegisterPublicEventRoutes(router *mux.Router, eventService *event.EventService) {
	// Route pour récupérer tous les événements (sans authentification)

	router.HandleFunc("/eventsMobile", func(w http.ResponseWriter, r *http.Request) {
		allEventsHandler := events.NewGetAllEventsMobileHandler(eventService)
		allEventsHandler.HandleGetAllEventsMobile(w, r)
	}).Methods("GET", "OPTIONS")

	router.Handle("/events/favorites",
		middleware.AuthMiddleware(http.HandlerFunc(events.NewGetFavoriteEventsHandler(eventService).HandleGetFavoriteEvents)),
	).Methods("GET", "OPTIONS")

	router.HandleFunc("/events/tags", func(w http.ResponseWriter, r *http.Request) {
		getAllTagsHandler := events.NewGetAllTagsHandler(eventService)
		getAllTagsHandler.HandleGetAllTags(w, r)
	}).Methods("GET", "OPTIONS")

}
