package db

import (
    "backend/config"
    "backend/internal/db/models" // Assurez-vous que ce chemin est correct
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=" + config.AppConfig.DB.Host +
        " user=" + config.AppConfig.DB.User +
        " password=" + config.AppConfig.DB.Password +
        " dbname=" + config.AppConfig.DB.Name +
        " port=" + config.AppConfig.DB.Port +
        " sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = database
}

// Migrate effectue les migrations des modèles vers la base de données
func Migrate() {
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
    )
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}
