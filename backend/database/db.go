package database

import (
	"backend/config"
	"backend/core/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// retryConnect tente de se connecter à la base de données plusieurs fois avant de retourner une erreur
func retryConnect(dsn string, maxRetries int, retryDelay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err == nil {
			log.Println("✅ Connexion à la base de données réussie.")
			return db, nil
		}
		log.Printf("❌ Tentative de connexion à la base de données échouée (%d/%d) : %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}
	return nil, fmt.Errorf("impossible de se connecter à la base de données après %d tentatives: %v", maxRetries, err)
}

// ConnectDatabase établit la connexion avec PostgreSQL
func ConnectDatabase() error {
	// Créer la Data Source Name (DSN) pour la connexion à la base de données
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DB.Host,
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Name,
		config.AppConfig.DB.Port)

	log.Println("🔄 Tentative de connexion à la base de données...")

	// Tenter plusieurs connexions à la base de données
	db, err := retryConnect(dsn, 5, 5*time.Second) // 5 tentatives avec 5 secondes d'intervalle
	if err != nil {
		return fmt.Errorf("❌ Erreur de connexion à la base de données : %v", err)
	}

	// Affecter la base de données globale
	DB = db
	log.Println("✅ Connexion à la base de données principale réussie.")

	return nil
}

// Migrate effectue la migration des modèles vers la base de données
func Migrate() error {
	log.Println("🔄 Démarrage de la migration des modèles...")

	// Migrer les modèles définis dans l'application
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.BusinessUser{},
		&models.EducationalInstitution{},
		&models.Association{},
		&models.Event{},
		&models.Ticket{},
		&models.Payment{},
		&models.PaymentTransaction{},
		&models.StudiboxTransaction{},
		&models.SchoolMembership{},
		&models.AssociationMembership{},
		&models.PasswordReset{},
		&models.PointHistory{},
		&models.EventOption{},
	)
	if err != nil {
		return fmt.Errorf("❌ Échec de la migration des modèles : %v", err)
	}

	log.Println("✅ Migration des modèles terminée avec succès.")
	return nil
}

// InitRoles initialise les rôles de base dans la base de données
func InitRoles(db *gorm.DB) error {
	// Liste des rôles à initialiser
	roles := []models.Role{
		{Name: "Admin"},
		{Name: "User"},
		{Name: "Business"},
	}

	// Initialiser chaque rôle s'il n'existe pas déjà
	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.Role{Name: role.Name}).Error; err != nil {
			return err
		}
		log.Printf("✅ Rôle '%s' vérifié ou créé avec succès.", role.Name)
	}

	log.Println("✅ Initialisation des rôles terminée avec succès.")
	return nil
}
