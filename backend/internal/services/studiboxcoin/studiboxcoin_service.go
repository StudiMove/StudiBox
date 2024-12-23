package studiboxcoin

import (
	"backend/internal/db/models"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type StudiboxCoinService struct {
	db *gorm.DB
}

// NewStudiboxCoinService initialise le service
func NewStudiboxCoinService(db *gorm.DB) *StudiboxCoinService {
	return &StudiboxCoinService{db: db}
}

// AddTransaction ajoute une nouvelle transaction Studibox
func (s *StudiboxCoinService) AddTransaction(userID uint, amount float64, coinsUsed int, eventEndDate time.Time, paymentStatus, paymentID, chargeID, klarnaStatus string) error {
	// Récupérer l'utilisateur
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return errors.New("utilisateur introuvable")
	}

	// Chercher le parrain (via le champ `ParrainCode` dans la table `User`)
	var referrerID *uint
	if user.ParrainCode != "" {
		var referrer models.User
		if err := s.db.Where("parrainage_code = ?", user.ParrainCode).First(&referrer).Error; err == nil {
			referrerID = &referrer.ID
		}
	}

	// Définir les coins à créditer
	coinsToUser := int(calculateUserCoins(amount)) // Conversion en int
	coinsToReferrer := int(calculateReferrerCoins(amount))

	// Créer la transaction
	transaction := models.StudiboxTransaction{
		UserID:          userID,
		Amount:          amount,
		ReferrerID:      referrerID,
		CoinsToUser:     coinsToUser,
		CoinsToReferrer: coinsToReferrer,
		CoinsUsed:       coinsUsed,
		EventEndDate:    eventEndDate,
		PaymentStatus:   paymentStatus,
		GlobalStatus:    "pending",
		PaymentID:       paymentID,
		ChargeID:        chargeID,
		KlarnaStatus:    klarnaStatus,
	}

	// Sauvegarder la transaction
	if err := s.db.Create(&transaction).Error; err != nil {
		return errors.New("échec de la création de la transaction")
	}

	return nil
}

// UpdateTransactionStatus met à jour le statut global d'une transaction
func (s *StudiboxCoinService) UpdateTransactionStatus(transactionID uint, status string) error {
	if err := s.db.Model(&models.StudiboxTransaction{}).
		Where("id = ?", transactionID).
		Update("global_status", status).Error; err != nil {
		return errors.New("échec de la mise à jour du statut de la transaction")
	}
	return nil
}

// RefundCoins gère le remboursement des coins après une annulation
func (s *StudiboxCoinService) RefundCoins(transactionID uint) error {
	var transaction models.StudiboxTransaction
	if err := s.db.First(&transaction, transactionID).Error; err != nil {
		return errors.New("transaction introuvable")
	}

	// Ajouter les coins utilisés à l'utilisateur
	if transaction.CoinsUsed > 0 {
		if err := s.db.Model(&models.User{}).
			Where("id = ?", transaction.UserID).
			Update("studibox_coins", gorm.Expr("studibox_coins + ?", transaction.CoinsUsed)).Error; err != nil {
			return errors.New("échec du remboursement des coins")
		}
	}

	// Mettre à jour le statut du paiement et de la transaction
	transaction.PaymentStatus = "refunded"
	transaction.GlobalStatus = "refunded"
	if err := s.db.Save(&transaction).Error; err != nil {
		return errors.New("échec de la mise à jour de la transaction")
	}

	return nil
}

// calculateUserCoins calcule les coins à créditer à l'utilisateur (logique métier personnalisable)
func calculateUserCoins(amount float64) float64 {
	// Exemple : Créditer 10% des coins utilisés
	return amount / 10
}

// calculateReferrerCoins calcule les coins à créditer au parrain (logique métier personnalisable)
func calculateReferrerCoins(amount float64) float64 {
	// Exemple : Créditer 5% des coins utilisés
	return amount / 20
}

// ProcessPendingTransactions traite automatiquement les transactions en attente après la fin de l'événement
func (s *StudiboxCoinService) ProcessPendingTransactions() error {
	now := time.Now()
	var pendingTransactions []models.StudiboxTransaction

	log.Println("Démarrage du traitement des transactions en attente...")

	// Récupérer les transactions en attente dont l'événement est terminé
	if err := s.db.Where("global_status = ? AND event_end_date <= ?", "pending", now).
		Find(&pendingTransactions).Error; err != nil {
		log.Printf("Erreur lors de la récupération des transactions en attente : %v", err)
		return errors.New("échec de la récupération des transactions en attente")
	}

	log.Printf("Nombre de transactions en attente trouvées : %d", len(pendingTransactions))

	for _, transaction := range pendingTransactions {
		log.Printf("Traitement de la transaction ID: %d pour l'utilisateur ID: %d", transaction.ID, transaction.UserID)

		// Créditez les coins à l'utilisateur
		if transaction.CoinsToUser > 0 {
			if err := s.db.Model(&models.User{}).
				Where("id = ?", transaction.UserID).
				Update("studibox_coins", gorm.Expr("studibox_coins + ?", transaction.CoinsToUser)).Error; err != nil {
				log.Printf("Erreur lors de l'attribution des coins à l'utilisateur ID: %d : %v", transaction.UserID, err)
				return errors.New("échec de l'attribution des coins à l'utilisateur")
			}
			log.Printf("Coins crédités à l'utilisateur ID: %d : %d coins", transaction.UserID, transaction.CoinsToUser)
		}

		// Créditez les coins au parrain
		if transaction.ReferrerID != nil && transaction.CoinsToReferrer > 0 {
			if err := s.db.Model(&models.User{}).
				Where("id = ?", *transaction.ReferrerID).
				Update("studibox_coins", gorm.Expr("studibox_coins + ?", transaction.CoinsToReferrer)).Error; err != nil {
				log.Printf("Erreur lors de l'attribution des coins au parrain ID: %d : %v", *transaction.ReferrerID, err)
				return errors.New("échec de l'attribution des coins au parrain")
			}
			log.Printf("Coins crédités au parrain ID: %d : %d coins", *transaction.ReferrerID, transaction.CoinsToReferrer)
		}

		// Mettre à jour le statut global de la transaction
		transaction.GlobalStatus = "done"
		if err := s.db.Save(&transaction).Error; err != nil {
			log.Printf("Erreur lors de la mise à jour de la transaction ID: %d après traitement : %v", transaction.ID, err)
			return errors.New("échec de la mise à jour de la transaction après traitement")
		}

		log.Printf("Transaction ID: %d mise à jour avec le statut 'done'", transaction.ID)
	}

	log.Println("Traitement des transactions en attente terminé.")
	return nil
}

// // ProcessPendingTransactions traite automatiquement les transactions en attente après la fin de l'événement
// func (s *StudiboxCoinService) ProcessPendingTransactions() error {
// 	now := time.Now()
// 	var pendingTransactions []models.StudiboxTransaction

// 	// Récupérer les transactions en attente dont l'événement est terminé
// 	if err := s.db.Where("global_status = ? AND event_end_date <= ?", "pending", now).
// 		Find(&pendingTransactions).Error; err != nil {
// 		return errors.New("échec de la récupération des transactions en attente")
// 	}

// 	for _, transaction := range pendingTransactions {
// 		// Créditez les coins à l'utilisateur
// 		if transaction.CoinsToUser > 0 {
// 			if err := s.db.Model(&models.User{}).
// 				Where("id = ?", transaction.UserID).
// 				Update("studibox_coins", gorm.Expr("studibox_coins + ?", transaction.CoinsToUser)).Error; err != nil {
// 				return errors.New("échec de l'attribution des coins à l'utilisateur")
// 			}
// 		}

// 		// Créditez les coins au parrain
// 		if transaction.ReferrerID != nil && transaction.CoinsToReferrer > 0 {
// 			if err := s.db.Model(&models.User{}).
// 				Where("id = ?", *transaction.ReferrerID).
// 				Update("studibox_coins", gorm.Expr("studibox_coins + ?", transaction.CoinsToReferrer)).Error; err != nil {
// 				return errors.New("échec de l'attribution des coins au parrain")
// 			}
// 		}

// 		// Mettre à jour le statut global de la transaction
// 		transaction.GlobalStatus = "done"
// 		if err := s.db.Save(&transaction).Error; err != nil {
// 			return errors.New("échec de la mise à jour de la transaction après traitement")
// 		}
// 	}

// 	return nil
// }
