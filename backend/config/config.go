package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	ServerPort          string
	S3Bucket            string
	APIBaseURL          string
	JwtSecretAccessKey  string
	JwtSecretRefreshKey string
	StripeSecretKey     string
	StripeWebhookSecret string
}

var AppConfig Config

func LoadConfig() {
	// Déterminer le chemin du fichier .env
	envPath := "./.env" // Le fichier .env est dans le même répertoire que votre exécutable

	// Vérifier si le fichier existe avant de charger
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("The .env file does not exist at path: %s", envPath)
	}

	viper.SetConfigFile(envPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	fmt.Println("Config file found and loaded successfully.")

	viper.AutomaticEnv() // Supporter les variables d'environnement

	// Charger les configurations
	AppConfig.DB.Host = viper.GetString("DB_HOST")
	AppConfig.DB.Port = viper.GetString("DB_PORT")
	AppConfig.DB.User = viper.GetString("DB_USER")
	AppConfig.DB.Password = viper.GetString("DB_PASSWORD")
	AppConfig.DB.Name = viper.GetString("DB_NAME")
	AppConfig.ServerPort = viper.GetString("SERVER_PORT")
	AppConfig.S3Bucket = viper.GetString("S3_BUCKET")
	AppConfig.APIBaseURL = viper.GetString("API_BASE_URL")
	AppConfig.JwtSecretAccessKey = viper.GetString("JWT_SECRET_ACCESS_KEY")
	AppConfig.JwtSecretRefreshKey = viper.GetString("JWT_SECRET_REFRESH_KEY")
	AppConfig.StripeSecretKey = viper.GetString("STRIPE_SECRET_KEY")
	AppConfig.StripeWebhookSecret = viper.GetString("STRIPE_WEBHOOK_SECRET")
}
