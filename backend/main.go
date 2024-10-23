package main

import (
	"fmt"
	"log"

	"backend/config"
	"backend/core/api/routes"
	"backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Charger la configuration depuis le fichier .env
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Erreur lors du chargement de la configuration : %v", err)
	}

	// Connexion à la base de données
	if err := database.ConnectDatabase(); err != nil {
		log.Fatalf("Échec de la connexion à la base de données : %v", err)
	}

	// Migrer automatiquement les modèles vers la base de données
	if err := database.Migrate(); err != nil {
		log.Fatalf("Échec de la migration de la base de données : %v", err)
	}

	// Initialiser les rôles après la migration des modèles
	if err := database.InitRoles(database.DB); err != nil {
		log.Fatalf("Échec de l'initialisation des rôles : %v", err)
	}

	// Initialiser les catégories après la migration des modèles
	if err := database.InitCategories(database.DB); err != nil {
		log.Fatalf("Échec de l'initialisation des catégories: %v", err)
	}

	// Initialiser les tags après la migration des modèles
	if err := database.InitTags(database.DB); err != nil {
		log.Fatalf("Échec de l'initialisation des tags : %v", err)
	}

	// Créer le routeur Gin
	router := gin.Default()

	// Enregistrer les routes, avec le middleware CORS géré dans routes.SetupRouter
	routes.SetupRouter(router)

	// Démarrer le serveur HTTP avec Gin
	serverAddress := fmt.Sprintf(":%s", config.AppConfig.ServerPort)
	log.Printf("Serveur démarré sur le port %s\n", config.AppConfig.ServerPort)
	log.Fatal(router.Run(serverAddress))
}
