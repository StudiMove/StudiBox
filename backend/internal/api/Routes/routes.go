// backend/internal/api/routes/routes.go

package routes

import (
    "net/http"
    "backend/internal/api/handlers/authentification"
    "backend/internal/api/handlers/events"
    "backend/internal/services/storage"
    "backend/internal/services/userservice"
    "backend/internal/services/auth"  // Import correct du service auth
    "backend/internal/api/routes/auth"
    "backend/internal/api/routes/user" // Correctement référencé
    "backend/internal/api/routes/business" // Ajout du package business
    "backend/internal/api/handlers/user/business/profil"



   

)

// RegisterRoutes enregistre toutes les routes de l'API
func RegisterRoutes(
    mux *http.ServeMux,
    storage storage.StorageService,
    authHandler *authentification.AuthHandler,
    registerHandler *authentification.RegisterHandler,
    createEventHandler *events.CreateEventHandler,
    getEventHandler *events.GetEventHandler,
    updateEventHandler *events.UpdateEventHandler,
    deleteEventHandler *events.DeleteEventHandler,
    userService *userservice.UserService,
    authService *auth.AuthService,  // Ajouter authService ici
    getProfilHandler *profil.GetProfilHandler,  // Nouveau handler pour le profil
    updateProfilHandler *profil.UpdateProfilHandler, // Nouveau handler pour la mise à jour du profil
) {
    // Enregistrer les routes d'authentification - Pas protégé
    auth_routes.RegisterAuthRoutes(mux, authHandler, registerHandler) // Appel à RegisterAuthRoutes

    // Enregistrer les routes des événements - Protégé
    mux.HandleFunc("/events", createEventHandler.HandleCreateEvent)
    mux.HandleFunc("/events/{id}", getEventHandler.HandleGetEvent)
    mux.HandleFunc("/events/{id}/update", updateEventHandler.HandleUpdateEvent)
    mux.HandleFunc("/events/{id}/delete", deleteEventHandler.HandleDeleteEvent)

    // Enregistrer les routes utilisateur - Protéger en passant authService
    user_routes.RegisterUserRoutes(mux, userService)
    business_routes.RegisterBusinessRoutes(mux, getProfilHandler, updateProfilHandler, authService)

}
