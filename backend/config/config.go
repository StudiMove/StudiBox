package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	DB                 DBConfig
	ServerPort         string
	S3Bucket           string
	APIBaseURL         string
	SMTPHost           string
	SMTPUser           string
	SMTPPass           string
	EmailFrom          string
	JwtSecretAccessKey string
}

var AppConfig *Config

// LoadConfig charge la configuration depuis un fichier .env ou les variables d'environnement
func LoadConfig() error {
	viper.SetConfigFile(".env") // Définit le fichier .env à utiliser

	// Lire le fichier de configuration, s'il existe
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Erreur lors de la lecture du fichier de configuration : %v. Utilisation des variables d'environnement uniquement.", err)
	} else {
		fmt.Println("Fichier de configuration .env chargé avec succès.")
	}

	viper.AutomaticEnv() // Permet de surcharger les valeurs avec les variables d'environnement

	// Charger la configuration en utilisant `os.Getenv` avec fallback sur Viper
	AppConfig = &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", viper.GetString("DB_HOST")),
			Port:     getEnv("DB_PORT", viper.GetString("DB_PORT")),
			User:     getEnv("DB_USER", viper.GetString("DB_USER")),
			Password: getEnv("DB_PASSWORD", viper.GetString("DB_PASSWORD")),
			Name:     getEnv("DB_NAME", viper.GetString("DB_NAME")),
		},
		ServerPort:         getEnv("SERVER_PORT", viper.GetString("SERVER_PORT")),
		S3Bucket:           getEnv("S3_BUCKET", viper.GetString("S3_BUCKET")),
		APIBaseURL:         getEnv("API_BASE_URL", viper.GetString("API_BASE_URL")),
		SMTPHost:           getEnv("SMTP_HOST", viper.GetString("SMTP_HOST")),
		SMTPUser:           getEnv("SMTP_USER", viper.GetString("SMTP_USER")),
		SMTPPass:           getEnv("SMTP_PASS", viper.GetString("SMTP_PASS")),
		EmailFrom:          getEnv("EMAIL_FROM", viper.GetString("EMAIL_FROM")),
		JwtSecretAccessKey: getEnv("JWT_SECRET_ACCESS_KEY", viper.GetString("JWT_SECRET_ACCESS_KEY")),
	}

	fmt.Println("Configuration chargée avec succès.")
	return nil
}

// getEnv vérifie si une variable d'environnement est définie, sinon prend la valeur par défaut
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	fmt.Printf("Variable d'environnement %s: %s\n", key, value)
	return value
}
