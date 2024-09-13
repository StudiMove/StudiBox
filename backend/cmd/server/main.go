package main

import (
    "fmt"
    "log"
    "net/http"
    "backend/config"
    "backend/internal/db"
    "backend/internal/api"
    "backend/internal/services/auth"
    "backend/internal/services/storage"
    "backend/internal/api/handlers/authentification"
    "backend/pkg/httpclient"
    "backend/internal/services/event"
    "backend/internal/api/handlers/events"
    "backend/internal/db/models" // Importez les modèles
)

func main() {
    // Charger la configuration
    config.LoadConfig()

    // Connexion à la base de données
    db.ConnectDatabase()

    // Migrer automatiquement les modèles
    db.DB.AutoMigrate(&models.Event{})

    // Créer une instance du service de stockage
    s3Service := storage.NewS3Storage(config.AppConfig.S3Bucket)

    // Créer le client HTTP
    apiClient := httpclient.NewAPIClient(config.AppConfig.APIBaseURL)

    // Créer les services
    authService := auth.NewAuthService(db.DB)
    eventService := event.NewEventService(db.DB)

    // Créer les handlers
    authHandler := authentification.NewAuthHandler(authService)
    registerHandler := authentification.NewRegisterHandler(authService, apiClient)
    createEventHandler := events.NewCreateEventHandler(eventService, apiClient)
    getEventHandler := events.NewGetEventHandler(eventService)
    updateEventHandler := events.NewUpdateEventHandler(eventService)
    deleteEventHandler := events.NewDeleteEventHandler(eventService)

    // Enregistrer les routes
    api.RegisterRoutes(s3Service, authHandler, registerHandler, createEventHandler, getEventHandler, updateEventHandler, deleteEventHandler)

    // Démarrer le serveur
    fmt.Println("Starting server on port", config.AppConfig.ServerPort)
    log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, nil))
}
