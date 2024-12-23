package scheduler

import (
	"backend/internal/services/studiboxcoin"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

func StartScheduler(db *gorm.DB) {
	s := gocron.NewScheduler(time.UTC)

	// Planifiez la tâche quotidienne à minuit
	s.Every(1).Day().At("00:00").Do(func() {
		log.Println("Traitement des transactions en attente...")
		service := studiboxcoin.NewStudiboxCoinService(db)
		if err := service.ProcessPendingTransactions(); err != nil {
			log.Printf("Erreur lors du traitement des transactions : %v", err)
		} else {
			log.Println("Transactions traitées avec succès.")
		}
	})

	// Démarrer le scheduler en mode non bloquant
	s.StartAsync()
}
