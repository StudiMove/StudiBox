package migrations

import (
	"backend/internal/db/models"
	"backend/internal/utils"
	"log"

	"gorm.io/gorm"
)

func AddParrainCodeColumn(db *gorm.DB) error {
	// Étape 1 : Ajouter la colonne ParrainCode sans contrainte NOT NULL
	log.Println("Ajout de la colonne ParrainCode sans contrainte NOT NULL...")
	if err := db.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS parrain_code TEXT").Error; err != nil {
		log.Printf("Erreur lors de l'ajout de la colonne ParrainCode : %v\n", err)
		return err
	}

	// Étape 2 : Initialiser tous les enregistrements existants avec un code généré
	log.Println("Initialisation des codes ParrainCode pour les utilisateurs existants...")
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Printf("Erreur lors de la récupération des utilisateurs : %v\n", err)
		return err
	}

	for _, user := range users {
		if user.ParrainCode == "" { // Ajouter un code uniquement si manquant
			user.ParrainCode = utils.GenerateParrainCode()
			if err := db.Save(&user).Error; err != nil {
				log.Printf("Erreur lors de la mise à jour de l'utilisateur %d : %v\n", user.ID, err)
				return err
			}
		}
	}

	// Étape 3 : Appliquer la contrainte NOT NULL après avoir rempli les données
	log.Println("Application de la contrainte NOT NULL sur la colonne ParrainCode...")
	if err := db.Exec("ALTER TABLE users ALTER COLUMN parrain_code SET NOT NULL").Error; err != nil {
		log.Printf("Erreur lors de l'application de la contrainte NOT NULL : %v\n", err)
		return err
	}

	log.Println("Migration de la colonne ParrainCode terminée avec succès.")
	return nil
}
