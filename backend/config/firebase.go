package config

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitializeFirebase() error {
	sa := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if sa == "" {
		return fmt.Errorf("chemin des credentials Firebase manquant")
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(sa))
	if err != nil {
		return fmt.Errorf("erreur d'initialisation de Firebase : %v", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("erreur d'initialisation du client Firebase : %v", err)
	}
	log.Println("Firebase a été initialisé avec succès.")
	return nil
}

func VerifyIDToken(idToken string) (*auth.Token, error) {
	token, err := FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, fmt.Errorf("token ID non valide : %v", err)
	}
	return token, nil
}
