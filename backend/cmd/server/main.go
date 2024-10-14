package main

import (
    "fmt"
    "log"
    "net/http"

    "backend/config"
    "backend/internal/api/routes"
    "backend/internal/api/middlewares"
    "backend/internal/api/handlers/authentification"
    "backend/internal/api/handlers/events"
    "backend/internal/api/handlers/user/business/profil"
    "backend/internal/db"
    "backend/internal/services/auth"
    "backend/internal/services/event"
    "backend/internal/services/storage"
    "backend/pkg/httpclient"
    "backend/internal/services/userservice"
    "backend/internal/services/userservice/business/profilservice" // Importer le profilservice
)

func main() {
    // Charger la configuration
    config.LoadConfig()

    // Connexion à la base de données
    db.ConnectDatabase()

    // Migrer automatiquement tous les modèles
    db.Migrate()

    // Initialiser les rôles
    db.InitRoles(db.DB)

    // Créer une instance du service de stockage
    s3Service := storage.NewS3Storage(config.AppConfig.S3Bucket)

    // Créer le client HTTP
    apiClient := httpclient.NewAPIClient(config.AppConfig.APIBaseURL)

    // Créer les services
    authService := auth.NewAuthService(db.DB)        
    eventService := event.NewEventService(db.DB)
    userService := userservice.NewUserService(db.DB) 
    profilService := profilservice.NewProfilService(db.DB) // Initialiser le service Profil

    // Créer les handlers
    authHandler := authentification.NewAuthHandler(authService)
    registerHandler := authentification.NewRegisterHandler(authService, apiClient)
    createEventHandler := events.NewCreateEventHandler(eventService, apiClient)
    getEventHandler := events.NewGetEventHandler(eventService)
    updateEventHandler := events.NewUpdateEventHandler(eventService)
    deleteEventHandler := events.NewDeleteEventHandler(eventService)

    // Créer les handlers pour le profil business avec profilService
    getProfilHandler := profil.NewGetProfilHandler(profilService)
    updateProfilHandler := profil.NewUpdateProfilHandler(profilService)

    // Enregistrer les routes
    mux := http.NewServeMux()
    routes.RegisterRoutes(
        mux,
        s3Service,
        authHandler,
        registerHandler,
        createEventHandler,
        getEventHandler,
        updateEventHandler,
        deleteEventHandler,
        userService,
        authService,
        getProfilHandler, // Ajout du handler de récupération du profil
        updateProfilHandler, // Ajout du handler de mise à jour du profil
    )

    // Démarrer le serveur
    serverAddress := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
    fmt.Printf("Démarrage du serveur sur le port %s\n", config.AppConfig.ServerPort)

    // Utiliser le middleware CORS
    log.Fatal(http.ListenAndServe(serverAddress, middleware.CORSMiddleware(mux)))
}
