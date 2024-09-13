package api

import (
    "net/http"
    "backend/internal/services/storage"
    "backend/internal/api/handlers/file"
    "backend/internal/api/handlers/authentification"
    "backend/internal/api/routes"
    "backend/internal/api/handlers/events"
)

// RegisterRoutes enregistre toutes les routes de l'API
func RegisterRoutes(storage storage.StorageService, authHandler *authentification.AuthHandler, registerHandler *authentification.RegisterHandler, createEventHandler *events.CreateEventHandler, getEventHandler *events.GetEventHandler, updateEventHandler *events.UpdateEventHandler, deleteEventHandler *events.DeleteEventHandler) {
    // Créer des handlers pour les différentes actions de fichiers
    fileUploadHandler := file.NewUploadFileHandler(storage)
    fileDeleteHandler := file.NewDeleteFileHandler(storage)
    fileGetURLHandler := file.NewGetFileURLHandler(storage)

    // Enregistrer les routes pour le téléchargement de fichiers
    http.HandleFunc("/upload", fileUploadHandler.HandleUpload)

    // Enregistrer les routes pour les opérations sur les fichiers
    http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodDelete:
            fileDeleteHandler.HandleDelete(w, r)
        case http.MethodGet:
            fileGetURLHandler.HandleGetFileURL(w, r)
        default:
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        }
    })

    // Enregistrer les routes d'authentification
    routes.RegisterAuthRoutes(authHandler, registerHandler)

    // Enregistrer les routes des événements
    routes.RegisterEventRoutes(createEventHandler, getEventHandler, updateEventHandler, deleteEventHandler)
}
