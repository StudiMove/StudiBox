// package webhook

// import (
// 	"backend/internal/db/models"
// 	"errors"
// 	"log"

// 	"github.com/stripe/stripe-go/v75/webhook"
// 	"gorm.io/gorm"
// )

// type WebhookService struct {
// 	db             *gorm.DB
// 	endpointSecret string
// }

// // NewWebhookService initialise le service des webhooks
// func NewWebhookService(db *gorm.DB, endpointSecret string) *WebhookService {
// 	return &WebhookService{
// 		db:             db,
// 		endpointSecret: endpointSecret,
// 	}
// }

// // SaveWebhook sauvegarde le webhook reçu dans la base de données
// func (s *WebhookService) SaveWebhook(stripeEventID string, eventType string, payload []byte) error {
// 	// Vérifie si l'événement existe déjà pour éviter les doublons
// 	var existingWebhook models.StripeWebhook
// 	if err := s.db.Where("stripe_event_id = ?", stripeEventID).First(&existingWebhook).Error; err == nil {
// 		log.Printf("Événement Stripe %s déjà traité, ignoré.", stripeEventID)
// 		return errors.New("webhook already processed")
// 	}

// 	// Crée une nouvelle entrée pour le webhook
// 	newWebhook := models.StripeWebhook{
// 		StripeEventID: stripeEventID,
// 		EventType:     eventType,
// 		Payload:       string(payload), // Convertit le contenu brut en chaîne
// 		Status:        "pending",
// 	}

// 	// Sauvegarde dans la base de données
// 	if err := s.db.Create(&newWebhook).Error; err != nil {
// 		log.Printf("Erreur lors de la sauvegarde du webhook Stripe : %v", err)
// 		return err
// 	}

// 	log.Printf("Webhook Stripe %s sauvegardé avec succès.", stripeEventID)
// 	return nil
// }

// // ProcessWebhook traite les événements en fonction de leur type
// func (s *WebhookService) ProcessWebhook(eventType string, payload []byte) error {
// 	switch eventType {
// 	case "payment_intent.succeeded":
// 		// Traitement spécifique pour un paiement réussi
// 		log.Println("Traitement du paiement réussi...")
// 		// Exécutez ici la logique nécessaire, par exemple :
// 		// - Mise à jour de la commande
// 		// - Notification utilisateur
// 	case "payment_intent.payment_failed":
// 		// Traitement spécifique pour un paiement échoué
// 		log.Println("Traitement du paiement échoué...")
// 		// Exécutez ici la logique pour notifier ou marquer la commande comme échouée
// 	default:
// 		log.Printf("Type d'événement non géré : %s", eventType)
// 		return errors.New("unsupported event type")
// 	}
// 	return nil
// }

// // VerifyAndHandleWebhook vérifie et gère le webhook reçu
// func (s *WebhookService) VerifyAndHandleWebhook(payload []byte, sigHeader string, endpointSecret string) error {
// 	// Étape 1 : Vérifie la signature Stripe
// 	event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
// 	if err != nil {
// 		log.Printf("Signature de webhook invalide : %v", err)
// 		return err
// 	}

// 	// Convertit `event.Type` (stripe.EventType) en `string`
// 	eventType := string(event.Type)

// 	// Étape 2 : Sauvegarde l'événement dans la base de données
// 	if err := s.SaveWebhook(event.ID, eventType, payload); err != nil {
// 		log.Printf("Erreur lors de la sauvegarde du webhook : %v", err)
// 		return err
// 	}

// 	// Étape 3 : Traite l'événement en fonction de son type
// 	if err := s.ProcessWebhook(eventType, payload); err != nil {
// 		log.Printf("Erreur lors du traitement du webhook : %v", err)
// 		return err
// 	}

// 	// Étape 4 : Met à jour le statut du webhook dans la base de données
// 	if err := s.db.Model(&models.StripeWebhook{}).Where("stripe_event_id = ?", event.ID).Update("status", "processed").Error; err != nil {
// 		log.Printf("Erreur lors de la mise à jour du statut du webhook : %v", err)
// 		return err
// 	}

//		log.Printf("Webhook Stripe %s traité avec succès.", event.ID)
//		return nil
//	}

// package webhook

// import (
// 	"backend/config"
// 	"backend/internal/db/models"
// 	"errors"
// 	"log"

// 	"github.com/stripe/stripe-go/v75/webhook"
// 	"gorm.io/gorm"
// )

// type WebhookService struct {
// 	db *gorm.DB
// }

// // NewWebhookService initialise le service des webhooks
// func NewWebhookService(db *gorm.DB) *WebhookService {
// 	return &WebhookService{
// 		db: db,
// 	}
// }

// // SaveWebhook sauvegarde le webhook reçu dans la base de données
// func (s *WebhookService) SaveWebhook(stripeEventID string, eventType string, payload []byte) error {
// 	var existingWebhook models.StripeWebhook
// 	if err := s.db.Where("stripe_event_id = ?", stripeEventID).First(&existingWebhook).Error; err == nil {
// 		log.Printf("Événement Stripe %s déjà traité, ignoré.", stripeEventID)
// 		return errors.New("webhook already processed")
// 	}

// 	newWebhook := models.StripeWebhook{
// 		StripeEventID: stripeEventID,
// 		EventType:     eventType,
// 		Payload:       string(payload),
// 		Status:        "pending",
// 	}

// 	if err := s.db.Create(&newWebhook).Error; err != nil {
// 		log.Printf("Erreur lors de la sauvegarde du webhook Stripe : %v", err)
// 		return err
// 	}

// 	log.Printf("Webhook Stripe %s sauvegardé avec succès.", stripeEventID)
// 	return nil
// }

// // ProcessWebhook traite les événements en fonction de leur type
// func (s *WebhookService) ProcessWebhook(eventType string, payload []byte) error {
// 	switch eventType {
// 	case "payment_intent.succeeded":
// 		return s.handlePaymentSucceeded(payload)
// 	case "payment_intent.payment_failed":
// 		return s.handlePaymentFailed(payload)
// 	default:
// 		log.Printf("Type d'événement non géré : %s", eventType)
// 		return errors.New("unsupported event type")
// 	}
// }

// // handlePaymentSucceeded traite les paiements réussis
// func (s *WebhookService) handlePaymentSucceeded(payload []byte) error {
// 	log.Println("Traitement du paiement réussi...")
// 	// Ajouter ici la logique métier spécifique
// 	return nil
// }

// // handlePaymentFailed traite les paiements échoués
// func (s *WebhookService) handlePaymentFailed(payload []byte) error {
// 	log.Println("Traitement du paiement échoué...")
// 	// Ajouter ici la logique métier spécifique
// 	return nil
// }

// // VerifyAndHandleWebhook vérifie et traite les webhooks Stripe
// func (s *WebhookService) VerifyAndHandleWebhook(payload []byte, sigHeader string) error {
// 	// Utilisation de AppConfig pour récupérer le StripeWebhookSecret
// 	endpointSecret := config.AppConfig.StripeWebhookSecret

// 	event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
// 	if err != nil {
// 		log.Printf("Signature invalide pour le webhook : %v", err)
// 		return err
// 	}

// 	eventType := string(event.Type)

// 	if err := s.SaveWebhook(event.ID, eventType, payload); err != nil {
// 		log.Printf("Erreur lors de la sauvegarde du webhook : %v", err)
// 		return err
// 	}

// 	if err := s.ProcessWebhook(eventType, payload); err != nil {
// 		log.Printf("Erreur lors du traitement du webhook : %v", err)
// 		return err
// 	}

// 	if err := s.db.Model(&models.StripeWebhook{}).Where("stripe_event_id = ?", event.ID).Update("status", "processed").Error; err != nil {
// 		log.Printf("Erreur lors de la mise à jour du statut du webhook : %v", err)
// 		return err
// 	}

//		log.Printf("Webhook Stripe traité avec succès : %s", event.ID)
//		return nil
//	}

package webhook

import (
	"backend/config"
	"backend/internal/db/models"
	"backend/internal/services/studiboxcoin"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/price"
	"github.com/stripe/stripe-go/v81/webhook"
	"gorm.io/gorm"
)

type WebhookService struct {
	db *gorm.DB
}

// NewWebhookService initialise le service des webhooks
func NewWebhookService(db *gorm.DB) *WebhookService {
	return &WebhookService{
		db: db,
	}
}

// SaveWebhook sauvegarde le webhook reçu dans la base de données
func (s *WebhookService) SaveWebhook(stripeEventID string, eventType string, payload []byte) error {
	var existingWebhook models.StripeWebhook
	// Vérifie si l'événement a déjà été traité
	if err := s.db.Where("stripe_event_id = ?", stripeEventID).First(&existingWebhook).Error; err == nil {
		log.Printf("Événement Stripe %s déjà traité, ignoré.", stripeEventID)
		return errors.New("webhook already processed")
	}

	// Crée un nouvel enregistrement pour le webhook
	newWebhook := models.StripeWebhook{
		StripeEventID: stripeEventID,
		EventType:     eventType,
		Payload:       string(payload),
		Status:        "pending",
	}

	if err := s.db.Create(&newWebhook).Error; err != nil {
		log.Printf("Erreur lors de la sauvegarde du webhook Stripe : %v", err)
		return err
	}

	log.Printf("Webhook Stripe %s sauvegardé avec succès.", stripeEventID)
	return nil
}

// ProcessWebhook traite les événements en fonction de leur type
func (s *WebhookService) ProcessWebhook(eventType string, payload []byte) error {
	switch eventType {
	case "payment_intent.succeeded":
		return s.handlePaymentSucceeded(payload)
	case "payment_intent.payment_failed":
		return s.handlePaymentFailed(payload)
	case "payment_intent.created":
		return s.handlePaymentIntentCreated(payload)
	case "charge.succeeded":
		return s.handleChargeSucceeded(payload)
	case "charge.updated":
		return s.handleChargeUpdated(payload)
	case "payment_intent.dispute.created":
		return s.handleDisputeCreated(payload)
	case "payment_intent.dispute.closed":
		return s.handleDisputeClosed(payload) // Nouveau cas pour les disputes fermées
	default:
		log.Printf("Type d'événement non géré : %s", eventType)
		return errors.New("unsupported event type")
	}
}

// handleChargeSucceeded traite l'événement charge.succeeded
func (s *WebhookService) handleChargeSucceeded(payload []byte) error {
	log.Println("Traitement de charge.succeeded...")
	// Ajouter la logique métier ici
	return nil
}

// handleChargeUpdated traite l'événement charge.updated
func (s *WebhookService) handleChargeUpdated(payload []byte) error {
	log.Println("Traitement de charge.updated...")
	// Ajouter la logique métier ici
	return nil
}

// handlePaymentFailed traite les paiements échoués
func (s *WebhookService) handlePaymentFailed(payload []byte) error {
	log.Println("Traitement du paiement échoué...")
	// Ajouter la logique métier ici
	return nil
}

// VerifyAndHandleWebhook vérifie et traite les webhooks Stripe
func (s *WebhookService) VerifyAndHandleWebhook(payload []byte, sigHeader string) error {
	endpointSecret := config.AppConfig.StripeWebhookSecret

	// Utilisation de ConstructEventWithOptions pour ignorer les conflits de version d'API
	opts := webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	}

	event, err := webhook.ConstructEventWithOptions(payload, sigHeader, endpointSecret, opts)
	if err != nil {
		log.Printf("Signature invalide pour le webhook : %v", err)
		return err
	}

	eventType := string(event.Type)

	if err := s.SaveWebhook(event.ID, eventType, payload); err != nil {
		log.Printf("Erreur lors de la sauvegarde du webhook : %v", err)
		return err
	}

	if err := s.ProcessWebhook(eventType, payload); err != nil {
		log.Printf("Erreur lors du traitement du webhook : %v", err)
		return err
	}

	// Met à jour le statut à "processed" une fois le traitement terminé
	if err := s.db.Model(&models.StripeWebhook{}).Where("stripe_event_id = ?", event.ID).Update("status", "processed").Error; err != nil {
		log.Printf("Erreur lors de la mise à jour du statut du webhook : %v", err)
		return err
	}

	log.Printf("Webhook Stripe traité avec succès : %s", event.ID)
	return nil
}
func (s *WebhookService) handlePaymentIntentCreated(payload []byte) error {
	log.Println("Traitement de payment_intent.created...")
	// Ajouter ici la logique métier si nécessaire
	return nil
}

// getEventEndDate récupère la date de fin de l'événement en fonction du produit
func (s *WebhookService) getEventEndDate(productIDStr string) (time.Time, error) {
	var event models.Event
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		log.Printf("Erreur de conversion product_id : %v", err)
		return time.Time{}, errors.New("erreur de conversion du product_id")
	}

	if err := s.db.First(&event, uint(productID)).Error; err != nil {
		log.Printf("Erreur lors de la récupération de l'événement : %v", err)
		return time.Time{}, errors.New("événement introuvable")
	}

	return event.EndDate, nil
}

// handlePaymentSucceeded traite les paiements réussis et met à jour les stocks

// func (s *WebhookService) handlePaymentSucceeded(payload []byte) error {
// 	log.Println("Traitement du paiement réussi...")

// 	// 1. Désérialiser le payload de l'événement
// 	var event stripe.Event
// 	if err := json.Unmarshal(payload, &event); err != nil {
// 		log.Printf("Erreur lors du parsing de l'événement Stripe : %v", err)
// 		return err
// 	}

// 	// 2. Convertir `event.Data.Object` en PaymentIntent
// 	var paymentIntent stripe.PaymentIntent
// 	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
// 		log.Printf("Erreur lors du parsing du PaymentIntent : %v", err)
// 		return err
// 	}

// 	// 3. Extraire les métadonnées
// 	userID := paymentIntent.Metadata["user_id"] // Assurez-vous que `user_id` est bien transmis dans les métadonnées
// 	selectedTarifs := paymentIntent.Metadata["selected_tarifs"]
// 	selectedOptions := paymentIntent.Metadata["selected_options"]
// 	studiboxCoins := paymentIntent.Metadata["studibox_coins"]
// 	productID := paymentIntent.Metadata["product_id"]

// 	if userID == "" {
// 		log.Println("Erreur : Aucun user_id trouvé dans les métadonnées.")
// 		return fmt.Errorf("user_id manquant dans les métadonnées")
// 	}

// 	if selectedTarifs == "" && selectedOptions == "" {
// 		log.Println("Aucun tarif ou option sélectionné dans les métadonnées.")
// 		return nil
// 	}

// 	// Convertir userID en uint
// 	userIDUint, err := strconv.ParseUint(userID, 10, 32)
// 	if err != nil {
// 		log.Printf("Erreur lors de la conversion de user_id : %v", err)
// 		return err
// 	}

// 	// Convertir studiboxCoins en int
// 	studiboxCoinsInt := 0
// 	if studiboxCoins != "" {
// 		studiboxCoinsInt, err = strconv.Atoi(studiboxCoins)
// 		if err != nil {
// 			log.Printf("Erreur lors de la conversion de studibox_coins : %v", err)
// 			return fmt.Errorf("studibox_coins invalide : %v", err)
// 		}
// 	}
// 	// 4. Traiter les `selected_tarifs` (tarifs)
// 	if selectedTarifs != "" {
// 		if err := s.processSelectedItems(selectedTarifs, "event_tarifs"); err != nil {
// 			log.Printf("Erreur lors du traitement des tarifs sélectionnés : %v", err)
// 			return err
// 		}
// 	}

// 	// 5. Traiter les `selected_options` (options)
// 	if selectedOptions != "" {
// 		if err := s.processSelectedItems(selectedOptions, "event_options"); err != nil {
// 			log.Printf("Erreur lors du traitement des options sélectionnées : %v", err)
// 			return err
// 		}
// 	}
// 	// 6. Rechercher l'ID du parrain
// 	referrerID, err := s.getReferrerID(uint(userIDUint))
// 	if err != nil {
// 		log.Printf("Erreur lors de la récupération du parrain pour l'utilisateur %d : %v", userIDUint, err)
// 		return err
// 	}

// 	if referrerID != nil {
// 		log.Printf("Parrain trouvé pour l'utilisateur %d : %d", userIDUint, *referrerID)
// 		// Vous pouvez maintenant utiliser `referrerID` pour attribuer des StudiboxCoins ou autre logique métier
// 	} else {
// 		log.Printf("Aucun parrain trouvé pour l'utilisateur %d", userIDUint)
// 	}
// 	// 7. Récupérer la date de fin de l'événement via getEventEndDate
// 	eventEndDate, err := s.getEventEndDate(productID)
// 	if err != nil {
// 		log.Printf("Erreur lors de la récupération de la date de fin de l'événement : %v", err)
// 		return err
// 	}
// 	// 8. Créer la transaction Studibox
// 	amount := float64(paymentIntent.AmountReceived) / 100.0 // Convertir en euros si nécessaire

// 	studiboxCoinService := studiboxcoin.NewStudiboxCoinService(s.db)
// 	if err := studiboxCoinService.AddTransaction(
// 		uint(userIDUint), // UserID
// 		amount,
// 		studiboxCoinsInt, // Coins utilisés
// 		eventEndDate,     // Date de fin de l'événement
// 		"succeeded",      // Statut de paiement
// 		paymentIntent.ID, // ID du paiement
// 	); err != nil {
// 		log.Printf("Erreur lors de la création de la transaction Studibox : %v", err)
// 		return err
// 	}

// 	log.Println("Transaction Studibox créée avec succès.")
// 	log.Println("Traitement du paiement réussi terminé.")
// 	return nil
// }

func (s *WebhookService) handlePaymentSucceeded(payload []byte) error {
	log.Println("Traitement du paiement réussi...")

	// 1. Désérialiser le payload de l'événement
	var event stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Erreur lors du parsing de l'événement Stripe : %v", err)
		return err
	}

	// 2. Convertir `event.Data.Object` en PaymentIntent
	var paymentIntent stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentIntent); err != nil {
		log.Printf("Erreur lors du parsing du PaymentIntent : %v", err)
		return err
	}

	// 3. Extraire les métadonnées
	userID := paymentIntent.Metadata["user_id"] // Assurez-vous que `user_id` est bien transmis dans les métadonnées
	selectedTarifs := paymentIntent.Metadata["selected_tarifs"]
	selectedOptions := paymentIntent.Metadata["selected_options"]
	studiboxCoins := paymentIntent.Metadata["studibox_coins"]
	productID := paymentIntent.Metadata["product_id"]

	if userID == "" {
		log.Println("Erreur : Aucun user_id trouvé dans les métadonnées.")
		return fmt.Errorf("user_id manquant dans les métadonnées")
	}

	// 4. Convertir userID en uint
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		log.Printf("Erreur lors de la conversion de user_id : %v", err)
		return err
	}

	// Convertir studiboxCoins en int
	studiboxCoinsInt := 0
	if studiboxCoins != "" {
		studiboxCoinsInt, err = strconv.Atoi(studiboxCoins)
		if err != nil {
			log.Printf("Erreur lors de la conversion de studibox_coins : %v", err)
			return fmt.Errorf("studibox_coins invalide : %v", err)
		}
	}

	// 5. Traiter les `selected_tarifs` (tarifs)
	if selectedTarifs != "" {
		if err := s.processSelectedItems(selectedTarifs, "event_tarifs"); err != nil {
			log.Printf("Erreur lors du traitement des tarifs sélectionnés : %v", err)
			return err
		}
	}

	// 6. Traiter les `selected_options` (options)
	if selectedOptions != "" {
		if err := s.processSelectedItems(selectedOptions, "event_options"); err != nil {
			log.Printf("Erreur lors du traitement des options sélectionnées : %v", err)
			return err
		}
	}

	// 7. Rechercher l'ID du parrain
	referrerID, err := s.getReferrerID(uint(userIDUint))
	if err != nil {
		log.Printf("Erreur lors de la récupération du parrain pour l'utilisateur %d : %v", userIDUint, err)
		return err
	}

	if referrerID != nil {
		log.Printf("Parrain trouvé pour l'utilisateur %d : %d", userIDUint, *referrerID)
	} else {
		log.Printf("Aucun parrain trouvé pour l'utilisateur %d", userIDUint)
	}

	// 8. Récupérer la date de fin de l'événement via getEventEndDate
	eventEndDate, err := s.getEventEndDate(productID)
	if err != nil {
		log.Printf("Erreur lors de la récupération de la date de fin de l'événement : %v", err)
		return err
	}

	// 9. Extraire le ChargeID
	chargeID := ""
	if paymentIntent.LatestCharge != nil {
		chargeID = paymentIntent.LatestCharge.ID
	}

	// 10. Gérer le KlarnaStatus (si applicable)
	klarnaStatus := ""
	if paymentIntent.PaymentMethodOptions != nil && paymentIntent.PaymentMethodOptions.Klarna != nil {
		klarnaStatus = paymentIntent.PaymentMethodOptions.Klarna.PreferredLocale
	}

	// 11. Créer la transaction Studibox
	amount := float64(paymentIntent.AmountReceived) / 100.0 // Convertir en euros si nécessaire

	studiboxCoinService := studiboxcoin.NewStudiboxCoinService(s.db)
	if err := studiboxCoinService.AddTransaction(
		uint(userIDUint), // UserID
		amount,           // Montant reçu
		studiboxCoinsInt, // Coins utilisés
		eventEndDate,     // Date de fin de l'événement
		"succeeded",      // Statut de paiement
		paymentIntent.ID, // ID du paiement
		chargeID,         // ChargeID
		klarnaStatus,     // KlarnaStatus
	); err != nil {
		log.Printf("Erreur lors de la création de la transaction Studibox : %v", err)
		return err
	}

	log.Println("Transaction Studibox créée avec succès.")
	log.Println("Traitement du paiement réussi terminé.")
	return nil
}

// processSelectedItems traite les items sélectionnés (tarifs ou options) et met à jour les stocks
func (s *WebhookService) processSelectedItems(selectedItems string, tableName string) error {
	// Diviser la liste des items
	items := strings.Split(selectedItems, ",")
	itemQuantities := make(map[string]int)

	// Compter les occurrences de chaque item
	for _, item := range items {
		itemQuantities[item]++
	}

	// Mettre à jour les stocks
	for priceID, quantity := range itemQuantities {
		log.Printf("Mise à jour du stock pour PriceID: %s, Quantité: %d dans la table %s", priceID, quantity, tableName)

		// Rechercher le stock actuel
		var stock int
		if err := s.db.Table(tableName).
			Where("price_id = ?", priceID).
			Select("stock").
			Row().
			Scan(&stock); err != nil {
			log.Printf("Erreur : Impossible de trouver le stock pour PriceID: %s dans la table %s", priceID, tableName)
			return err
		}

		// Vérifier si le stock est suffisant
		if stock < quantity {
			log.Printf("Stock insuffisant pour PriceID: %s dans la table %s. Stock actuel: %d, Demandé: %d", priceID, tableName, stock, quantity)
			return fmt.Errorf("stock insuffisant pour PriceID: %s", priceID)
		}

		// Réduire le stock
		if err := s.db.Table(tableName).
			Where("price_id = ?", priceID).
			Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {
			log.Printf("Erreur lors de la mise à jour du stock pour PriceID: %s dans la table %s", priceID, tableName)
			return err
		}

		// Mettre à jour le stock dans Stripe
		newStock := stock - quantity
		if err := updateStripePriceStock(priceID, newStock); err != nil {
			log.Printf("Erreur lors de la mise à jour du stock Stripe pour PriceID: %s", priceID)
			return err
		}
	}

	return nil
}

// updateStripePriceStock met à jour le stock dans les métadonnées Stripe pour un PriceID
func updateStripePriceStock(priceID string, newStock int) error {
	params := &stripe.PriceParams{
		Metadata: map[string]string{
			"stock": fmt.Sprintf("%d", newStock),
		},
	}

	_, err := price.Update(priceID, params)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du stock Stripe pour le PriceID %s : %w", priceID, err)
	}

	return nil
}

// getReferrerID récupère l'ID du parrain en utilisant l'ID de l'utilisateur
func (s *WebhookService) getReferrerID(userID uint) (*uint, error) {
	var user models.User
	// Récupérer le ParrainageCode de l'utilisateur
	if err := s.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Utilisateur non trouvé avec l'ID : %d", userID)
			return nil, nil // Pas de parrain, donc pas d'erreur critique
		}
		return nil, err
	}

	// Si l'utilisateur n'a pas de ParrainCode, retourner nil
	if user.ParrainCode == "" {
		log.Printf("Utilisateur %d n'a pas de parrainage actif.", userID)
		return nil, nil
	}

	// Rechercher le parrain basé sur le ParrainCode
	var referrer models.User
	if err := s.db.Where("parrainage_code = ?", user.ParrainCode).First(&referrer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Aucun parrain trouvé pour le ParrainCode : %s", user.ParrainCode)
			return nil, nil
		}
		return nil, err
	}

	// Retourner l'ID du parrain
	return &referrer.ID, nil
}

func (s *WebhookService) handleDisputeCreated(payload []byte) error {
	log.Println("Traitement d'une dispute créée...")

	// Désérialiser le payload Stripe
	var event stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Erreur lors du parsing de l'événement Stripe : %v", err)
		return err
	}

	var dispute stripe.Dispute
	if err := json.Unmarshal(event.Data.Raw, &dispute); err != nil {
		log.Printf("Erreur lors du parsing de la dispute : %v", err)
		return err
	}

	chargeID := dispute.Charge.ID
	log.Printf("Dispute détectée : Charge ID = %s, Raison = %s, Montant = %.2f€", chargeID, dispute.Reason, float64(dispute.Amount)/100)

	// Rechercher la transaction associée à la charge
	var transaction models.StudiboxTransaction
	if err := s.db.Where("charge_id = ?", chargeID).First(&transaction).Error; err != nil {
		log.Printf("Erreur : Transaction non trouvée pour Charge ID %s", chargeID)
		return fmt.Errorf("Transaction non trouvée pour le Charge ID : %s", chargeID)
	}

	// Vérifier si l'événement associé à la transaction est terminé
	if time.Now().After(transaction.EventEndDate) {
		log.Printf("Dispute refusée : L'événement est terminé pour la transaction associée au Charge ID %s", chargeID)
		transaction.PaymentStatus = "not_refundable"
		if err := s.db.Save(&transaction).Error; err != nil {
			log.Printf("Erreur lors de la mise à jour du statut de la transaction pour Charge ID %s : %v", chargeID, err)
			return fmt.Errorf("Erreur lors de la mise à jour de la transaction : %v", err)
		}
		return nil
	}

	// Mettre à jour la transaction avec l'état "disputed" et enregistrer l'ID de la dispute
	transaction.PaymentStatus = "disputed"
	transaction.DisputeID = dispute.ID
	if err := s.db.Save(&transaction).Error; err != nil {
		log.Printf("Erreur lors de la mise à jour de la transaction pour Charge ID %s : %v", chargeID, err)
		return fmt.Errorf("Erreur lors de la mise à jour de la transaction : %v", err)
	}

	log.Printf("Transaction mise à jour avec succès : Charge ID = %s, Dispute ID = %s", chargeID, dispute.ID)
	log.Println("Traitement de la dispute terminé.")
	return nil
}

func (s *WebhookService) handleDisputeClosed(payload []byte) error {
	log.Println("Traitement d'une dispute fermée...")

	var event stripe.Event
	if err := json.Unmarshal(payload, &event); err != nil {
		log.Printf("Erreur lors du parsing de l'événement Stripe : %v", err)
		return err
	}

	var dispute stripe.Dispute
	if err := json.Unmarshal(event.Data.Raw, &dispute); err != nil {
		log.Printf("Erreur lors du parsing de la dispute : %v", err)
		return err
	}

	chargeID := dispute.Charge.ID
	status := dispute.Status

	log.Printf("Dispute fermée : Charge ID = %s, Statut = %s", chargeID, status)

	// Mettre à jour la transaction en fonction du statut
	switch status {
	case "won":
		if err := s.markTransactionAsResolved(chargeID, "dispute_won"); err != nil {
			log.Printf("Erreur lors de la mise à jour de la transaction gagnée pour Charge ID %s : %v", chargeID, err)
			return err
		}
		// Restaurer les ressources si nécessaire
		log.Printf("Ressources restaurées pour la transaction gagnée avec Charge ID %s", chargeID)
	case "lost":
		if err := s.markTransactionAsResolved(chargeID, "dispute_lost"); err != nil {
			log.Printf("Erreur lors de la mise à jour de la transaction perdue pour Charge ID %s : %v", chargeID, err)
			return err
		}
		log.Printf("Aucune action supplémentaire pour la transaction perdue avec Charge ID %s", chargeID)
	default:
		log.Printf("Statut de dispute non reconnu : %s pour Charge ID %s", status, chargeID)
		return nil
	}

	log.Println("Traitement de la dispute fermée terminé avec succès.")
	return nil
}

func (s *WebhookService) markTransactionAsResolved(chargeID, resolutionStatus string) error {
	var transaction models.StudiboxTransaction
	if err := s.db.Where("charge_id = ?", chargeID).First(&transaction).Error; err != nil {
		log.Printf("Transaction introuvable pour Charge ID : %s", chargeID)
		return fmt.Errorf("transaction non trouvée pour le charge ID : %s", chargeID)
	}

	transaction.PaymentStatus = resolutionStatus
	transaction.DisputeID = "" // Nettoyage du champ DisputeID si nécessaire
	if err := s.db.Save(&transaction).Error; err != nil {
		log.Printf("Échec de la mise à jour de la transaction pour Charge ID %s : %v", chargeID, err)
		return fmt.Errorf("échec de la mise à jour de la transaction : %v", err)
	}

	log.Printf("Transaction avec Charge ID %s mise à jour avec le statut : %s.", chargeID, resolutionStatus)
	return nil
}
