// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"backend/config"
// 	"backend/internal/api/routes"
// 	"backend/internal/db"
// 	"backend/internal/services/auth"
// 	"backend/internal/services/event"
// 	"backend/internal/services/password"
// 	"backend/internal/services/profilservice" // Import du service profil
// 	"backend/internal/services/storage"
// 	"backend/pkg/httpclient"
// )

// func main() {
// 	// Charger la configuration de l'application
// 	fmt.Println("Chargement de la configuration de l'application...")
// 	config.LoadConfig()

// 	// Connecter à la base de données et appliquer les migrations
// 	fmt.Println("Connexion à la base de données...")
// 	dbConnection := db.ConnectDatabase()
// 	fmt.Println("Application des migrations...")
// 	db.Migrate()
// 	fmt.Println("Initialisation des rôles dans la base de données...")
// 	db.InitRoles(dbConnection)

// 	// Initialiser les services nécessaires
// 	fmt.Println("Initialisation des services...")
// 	s3Service := storage.NewS3Storage(config.AppConfig.S3Bucket)
// 	apiClient := httpclient.NewAPIClient(config.AppConfig.APIBaseURL)
// 	authService := auth.NewAuthService(dbConnection)
// 	eventService := event.NewEventService(dbConnection, s3Service)
// 	passwordService := password.NewPasswordResetService(dbConnection)
// 	profilService := profilservice.NewProfilService(dbConnection) // Initialisation de profilService

// 	// Configurer le routeur avec les routes et middlewares
// 	fmt.Println("Configuration des routes...")
// 	router := routes.InitRouter(authService, eventService, apiClient, passwordService, profilService, dbConnection)

//		// Démarrer le serveur
//		serverAddress := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
//		fmt.Printf("Démarrage du serveur sur le port %s\n", config.AppConfig.ServerPort)
//		log.Fatal(http.ListenAndServe(serverAddress, router))
//	}
//
// main.go

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"backend/config"
// 	"backend/internal/api/routes"
// 	"backend/internal/db"
// 	"backend/internal/services/auth"
// 	"backend/internal/services/event"
// 	"backend/internal/services/password"
// 	"backend/internal/services/profilservice"
// 	"backend/internal/services/storage"
// 	"backend/internal/services/userservice"
// 	"backend/pkg/httpclient"
// )

// func main() {
// 	// Charger la configuration de l'application
// 	fmt.Println("Chargement de la configuration de l'application...")
// 	config.LoadConfig()

// 	// Connecter à la base de données et appliquer les migrations
// 	fmt.Println("Connexion à la base de données...")
// 	dbConnection := db.ConnectDatabase()
// 	fmt.Println("Application des migrations...")
// 	db.Migrate()
// 	fmt.Println("Initialisation des rôles dans la base de données...")
// 	db.InitRoles(dbConnection)

// 	// Initialiser les services nécessaires
// 	fmt.Println("Initialisation des services...")
// 	s3Service := storage.NewS3Storage(config.AppConfig.S3Bucket)
// 	apiClient := httpclient.NewAPIClient(config.AppConfig.APIBaseURL)
// 	authService := auth.NewAuthService(dbConnection)
// 	eventService := event.NewEventService(dbConnection, s3Service)
// 	passwordService := password.NewPasswordResetService(dbConnection)
// 	profilService := profilservice.NewProfilService(dbConnection)
// 	userService := userservice.NewUserService(dbConnection)

// 	// Configurer le routeur avec les routes et middlewares
// 	fmt.Println("Configuration des routes...")
// 	router := routes.InitRouter(authService, eventService, apiClient, passwordService, profilService, userService, s3Service, dbConnection)

//		// Démarrer le serveur
//		serverAddress := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
//		fmt.Printf("Démarrage du serveur sur le port %s\n", config.AppConfig.ServerPort)
//		log.Fatal(http.ListenAndServe(serverAddress, router))
//	}

package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/internal/api/routes"
	"backend/internal/db"
	"backend/internal/services/auth"
	"backend/internal/services/event"
	"backend/internal/services/password"
	"backend/internal/services/profilservice"
	"backend/internal/services/storage"
	"backend/internal/services/userservice"
	"backend/pkg/httpclient"
)

func main() {
	// Charger la configuration de l'application
	fmt.Println("Chargement de la configuration de l'application...")
	config.LoadConfig()

	// Connecter à la base de données et appliquer les migrations
	fmt.Println("Connexion à la base de données...")
	dbConnection := db.ConnectDatabase()
	fmt.Println("Application des migrations...")
	db.Migrate()
	fmt.Println("Initialisation des rôles dans la base de données...")
	db.InitRoles(dbConnection)

	// Initialiser les services nécessaires
	fmt.Println("Initialisation des services...")
	s3Service := storage.NewS3Storage(config.AppConfig.S3Bucket)
	apiClient := httpclient.NewAPIClient(config.AppConfig.APIBaseURL)
	authService := auth.NewAuthService(dbConnection)
	eventService := event.NewEventService(dbConnection, s3Service)
	passwordService := password.NewPasswordResetService(dbConnection)
	profilService := profilservice.NewProfilService(dbConnection, s3Service)

	userService := userservice.NewUserService(dbConnection)
	jwtSecret := config.AppConfig.JwtSecretAccessKey

	// Configurer le routeur avec les routes et middlewares
	fmt.Println("Configuration des routes...")
	router := routes.InitRouter(authService, eventService, apiClient, passwordService, profilService, userService, s3Service, dbConnection, jwtSecret)

	// Démarrer le serveur
	serverAddress := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
	fmt.Printf("Démarrage du serveur sur le port %s\n", config.AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
