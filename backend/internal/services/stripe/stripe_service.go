// package stripe

// import (
// 	"backend/config"
// 	"backend/internal/db/models"
// 	"fmt"
// 	"log"

// 	"github.com/stripe/stripe-go/v75"
// 	"github.com/stripe/stripe-go/v75/price"
// 	"github.com/stripe/stripe-go/v75/product"
// )

// type StripeService struct{}

// // NewStripeService initialise le service Stripe avec la clé API provenant de config.AppConfig
// func NewStripeService() *StripeService {
// 	stripe.Key = config.AppConfig.StripeSecretKey // Utilisation explicite de la clé
// 	return &StripeService{}
// }

// // CreateProductFromEvent crée un produit Stripe à partir d'un événement
// func (s *StripeService) CreateProductFromEvent(event *models.Event) (string, error) {
// 	// Mapper les données de l'événement vers le produit Stripe
// 	params := &stripe.ProductParams{
// 		Name:        stripe.String(event.Title),
// 		Description: stripe.String(event.Subtitle), // Utiliser le sous-titre comme description
// 		Metadata: map[string]string{
// 			"event_id": fmt.Sprintf("%d", event.ID), // Associer l'event ID
// 		},
// 	}

// 	// Créer le produit Stripe
// 	prod, err := product.New(params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la création du produit Stripe : %v", err)
// 		return "", err
// 	}

// 	log.Printf("Produit Stripe créé pour l'événement : %s (Stripe ID : %s)", event.Title, prod.ID)
// 	return prod.ID, nil
// }

// // CreatePriceFromTarif crée un tarif Stripe pour un produit à partir d'un tarif d'événement
// func (s *StripeService) CreatePriceFromTarif(productID string, tarif models.EventTarif) (string, error) {
// 	params := &stripe.PriceParams{
// 		Product:    stripe.String(productID),
// 		UnitAmount: stripe.Int64(int64(tarif.Price * 100)), // Convertir le prix en centimes
// 		Currency:   stripe.String("eur"),
// 		Metadata: map[string]string{
// 			"title":    tarif.Title,
// 			"stock":    fmt.Sprintf("%d", tarif.Stock),
// 			"isTarif":  "true",
// 			"isOption": "false",
// 		},
// 	}

// 	// Créer le tarif Stripe
// 	p, err := price.New(params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la création du tarif Stripe : %v", err)
// 		return "", err
// 	}

// 	log.Printf("Tarif Stripe créé pour le produit %s : %s", productID, p.ID)
// 	return p.ID, nil
// }

// // CreateOptionPrice crée un prix Stripe pour une option
// func (s *StripeService) CreateOptionPrice(productID string, option models.EventOption) (string, error) {
// 	params := &stripe.PriceParams{
// 		Product:    stripe.String(productID),
// 		UnitAmount: stripe.Int64(int64(option.Price * 100)),
// 		Currency:   stripe.String("eur"),
// 		Metadata: map[string]string{
// 			"title":       option.Title,
// 			"description": option.Description,
// 			"stock":       fmt.Sprintf("%d", option.Stock),
// 			"isTarif":     "false",
// 			"isOption":    "true",
// 		},
// 	}

// 	// Créer le prix dans Stripe
// 	price, err := price.New(params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la création du prix pour l'option : %v", err)
// 		return "", fmt.Errorf("Stripe price creation failed: %w", err)
// 	}

// 	// Validation : Vérifier que price.ID est non vide
// 	if price.ID == "" {
// 		log.Printf("Erreur : Stripe a retourné un PriceID vide pour l'option '%s'", option.Title)
// 		return "", fmt.Errorf("Stripe returned an empty PriceID for option '%s'", option.Title)
// 	}

// 	log.Printf("Option Stripe créée pour le produit %s : %s (Option ID : %s)", productID, price.ID, option.Title)
// 	return price.ID, nil
// }

// // CreateProductAndPrices crée un produit et ses tarifs associés pour un événement
// func (s *StripeService) CreateProductAndPrices(event *models.Event) (string, error) {
// 	// Étape 1 : Créer le produit Stripe
// 	productID, err := s.CreateProductFromEvent(event)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Étape 2 : Créer les tarifs associés
// 	for _, tarif := range event.Tarifs {
// 		_, err := s.CreatePriceFromTarif(productID, tarif)
// 		if err != nil {
// 			log.Printf("Erreur lors de la création d'un tarif pour l'événement : %v", err)
// 			return "", err
// 		}
// 	}

// 	return productID, nil
// }

// // UpdatePrice met à jour un prix Stripe existant

// func (s *StripeService) UpdatePrice(priceID string, newPrice models.EventTarif) error {
// 	params := &stripe.PriceParams{
// 		Metadata: map[string]string{
// 			"title":    newPrice.Title,
// 			"stock":    fmt.Sprintf("%d", newPrice.Stock),
// 			"isTarif":  "true",
// 			"isOption": "false",
// 		},
// 	}

// 	_, err := price.Update(priceID, params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la mise à jour du prix Stripe '%s' : %v", priceID, err)
// 		return fmt.Errorf("failed to update Stripe price '%s': %w", priceID, err)
// 	}

// 	log.Printf("Prix Stripe '%s' mis à jour avec succès", priceID)
// 	return nil
// }

// // UpdateOptionPrice met à jour une option Stripe existante
// func (s *StripeService) UpdateOptionPrice(priceID string, newOption models.EventOption) error {
// 	params := &stripe.PriceParams{
// 		Metadata: map[string]string{
// 			"title":       newOption.Title,
// 			"description": newOption.Description,
// 			"stock":       fmt.Sprintf("%d", newOption.Stock),
// 			"isTarif":     "false",
// 			"isOption":    "true",
// 		},
// 	}

// 	_, err := price.Update(priceID, params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la mise à jour de l'option Stripe '%s' : %v", priceID, err)
// 		return fmt.Errorf("failed to update Stripe option price '%s': %w", priceID, err)
// 	}

// 	log.Printf("Option Stripe '%s' mise à jour avec succès", priceID)
// 	return nil
// }

// func (s *StripeService) DisablePrice(priceID string) error {
// 	_, err := price.Update(priceID, &stripe.PriceParams{Active: stripe.Bool(false)})
// 	if err != nil {
// 		log.Printf("Erreur lors de la désactivation du prix Stripe : %v", err)
// 		return fmt.Errorf("failed to disable Stripe price '%s': %w", priceID, err)
// 	}

// 	log.Printf("Prix Stripe désactivé avec succès : %s", priceID)
// 	return nil
// }

// // SyncPrices synchronise les tarifs et options avec Stripe
// func (s *StripeService) SyncPrices(productID string, currentPrices []models.EventTarif, currentOptions []models.EventOption) error {
// 	// Étape 1 : Récupérer tous les prix actuels de Stripe pour le produit
// 	iter := price.List(&stripe.PriceListParams{
// 		Product: stripe.String(productID),
// 		Active:  stripe.Bool(true),
// 	})

// 	// Créer un ensemble de PriceIDs existants dans Stripe
// 	existingStripePrices := make(map[string]bool)
// 	for iter.Next() {
// 		p := iter.Price()
// 		existingStripePrices[p.ID] = true
// 	}

// 	// Étape 2 : Synchroniser les tarifs
// 	for _, tarif := range currentPrices {
// 		if tarif.PriceID == "" || !existingStripePrices[tarif.PriceID] {
// 			// Créer un nouveau prix si le PriceID est vide ou non trouvé
// 			priceID, err := s.CreatePriceFromTarif(productID, tarif)
// 			if err != nil {
// 				log.Printf("Erreur lors de la création d'un nouveau tarif Stripe : %v", err)
// 				return fmt.Errorf("failed to sync Stripe price for tarif '%s': %w", tarif.Title, err)
// 			}
// 			tarif.PriceID = priceID
// 			log.Printf("Nouveau tarif synchronisé dans Stripe : %s", priceID)
// 		} else {
// 			// Mettre à jour le tarif existant
// 			err := s.UpdatePrice(tarif.PriceID, tarif)
// 			if err != nil {
// 				log.Printf("Erreur lors de la mise à jour d'un tarif Stripe : %v", err)
// 				return fmt.Errorf("failed to update Stripe price for tarif '%s': %w", tarif.Title, err)
// 			}
// 		}
// 	}

// 	// Étape 3 : Synchroniser les options
// 	for _, option := range currentOptions {
// 		if option.PriceID == "" || !existingStripePrices[option.PriceID] {
// 			// Créer un nouveau prix si le PriceID est vide ou non trouvé
// 			priceID, err := s.CreateOptionPrice(productID, option)
// 			if err != nil {
// 				log.Printf("Erreur lors de la création d'une nouvelle option Stripe : %v", err)
// 				return fmt.Errorf("failed to sync Stripe price for option '%s': %w", option.Title, err)
// 			}
// 			option.PriceID = priceID
// 			log.Printf("Nouvelle option synchronisée dans Stripe : %s", priceID)
// 		} else {
// 			// Mettre à jour l'option existante
// 			err := s.UpdateOptionPrice(option.PriceID, option)
// 			if err != nil {
// 				log.Printf("Erreur lors de la mise à jour d'une option Stripe : %v", err)
// 				return fmt.Errorf("failed to update Stripe price for option '%s': %w", option.Title, err)
// 			}
// 		}
// 	}

// 	// Étape 4 : Désactiver les anciens prix qui ne sont plus utilisés
// 	for stripePriceID := range existingStripePrices {
// 		found := false
// 		for _, tarif := range currentPrices {
// 			if tarif.PriceID == stripePriceID {
// 				found = true
// 				break
// 			}
// 		}
// 		for _, option := range currentOptions {
// 			if option.PriceID == stripePriceID {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			// Désactiver le prix sur Stripe
// 			_, err := price.Update(stripePriceID, &stripe.PriceParams{Active: stripe.Bool(false)})
// 			if err != nil {
// 				log.Printf("Erreur lors de la désactivation du prix Stripe : %v", err)
// 				return fmt.Errorf("failed to deactivate unused Stripe price '%s': %w", stripePriceID, err)
// 			}
// 			log.Printf("Prix Stripe désactivé : %s", stripePriceID)
// 		}
// 	}

//		log.Printf("Synchronisation des prix Stripe terminée pour le produit : %s", productID)
//		return nil
//	}
//
// ////////

package stripe

import (
	"backend/config"
	"backend/internal/db/models"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/price"
	"github.com/stripe/stripe-go/v81/product"
)

type StripeService struct{}

// NewStripeService initialise le service Stripe avec la clé API provenant de config.AppConfig
func NewStripeService() *StripeService {
	stripe.Key = config.AppConfig.StripeSecretKey
	log.Println("Stripe service initialized with secret key.")
	return &StripeService{}
}

// CreateProductFromEvent crée un produit Stripe à partir d'un événement
func (s *StripeService) CreateProductFromEvent(event *models.Event) (string, error) {
	params := &stripe.ProductParams{
		Name:        stripe.String(event.Title),
		Description: stripe.String(event.Subtitle),
		Metadata: map[string]string{
			"event_id": fmt.Sprintf("%d", event.ID),
		},
	}

	prod, err := product.New(params)
	if err != nil {
		log.Printf("Erreur lors de la création du produit Stripe : %v", err)
		return "", err
	}

	log.Printf("Produit Stripe créé pour l'événement : %s (Stripe ID : %s)", event.Title, prod.ID)
	return prod.ID, nil
}

// CreatePriceFromTarif crée un tarif Stripe pour un produit à partir d'un tarif d'événement
func (s *StripeService) CreatePriceFromTarif(productID string, tarif models.EventTarif) (string, error) {
	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		UnitAmount: stripe.Int64(int64(tarif.Price * 100)),
		Currency:   stripe.String("eur"),
		Metadata: map[string]string{
			"title":    tarif.Title,
			"stock":    fmt.Sprintf("%d", tarif.Stock),
			"isTarif":  "true",
			"isOption": "false",
		},
	}

	p, err := price.New(params)
	if err != nil {
		log.Printf("Erreur lors de la création du tarif Stripe : %v", err)
		return "", err
	}

	log.Printf("Tarif Stripe créé pour le produit %s : %s", productID, p.ID)
	return p.ID, nil
}

// CreateOptionPrice crée un prix Stripe pour une option
func (s *StripeService) CreateOptionPrice(productID string, option models.EventOption) (string, error) {
	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		UnitAmount: stripe.Int64(int64(option.Price * 100)),
		Currency:   stripe.String("eur"),
		Metadata: map[string]string{
			"title":       option.Title,
			"description": option.Description,
			"stock":       fmt.Sprintf("%d", option.Stock),
			"isTarif":     "false",
			"isOption":    "true",
		},
	}

	p, err := price.New(params)
	if err != nil {
		log.Printf("Erreur lors de la création du prix pour l'option : %v", err)
		return "", fmt.Errorf("Stripe price creation failed: %w", err)
	}

	if p.ID == "" {
		log.Printf("Erreur : Stripe a retourné un PriceID vide pour l'option '%s'", option.Title)
		return "", fmt.Errorf("Stripe returned an empty PriceID for option '%s'", option.Title)
	}

	log.Printf("Option Stripe créée pour le produit %s : %s (Option ID : %s)", productID, p.ID, option.Title)
	return p.ID, nil
}

// UpdatePrice recrée un prix Stripe existant (nouveau prix + désactivation de l'ancien)
func (s *StripeService) UpdatePrice(productID string, currentPriceID string, newPrice models.EventTarif) (string, error) {
	// Vérification préalable de l'ancien prix
	if currentPriceID == "" {
		log.Printf("Erreur : PriceID actuel manquant pour le tarif '%s'", newPrice.Title)
		return "", fmt.Errorf("current PriceID is missing for tarif '%s'", newPrice.Title)
	}

	// Désactiver l'ancien prix
	if err := s.DisablePrice(currentPriceID); err != nil {
		log.Printf("Erreur lors de la désactivation du prix Stripe '%s' : %v", currentPriceID, err)
		return "", fmt.Errorf("failed to disable old Stripe price '%s': %w", currentPriceID, err)
	}

	// Créer un nouveau prix avec les nouvelles données
	newPriceID, err := s.CreatePriceFromTarif(productID, newPrice)
	if err != nil {
		log.Printf("Erreur lors de la recréation du prix Stripe pour '%s': %v", newPrice.Title, err)
		return "", fmt.Errorf("failed to create new Stripe price for tarif '%s': %w", newPrice.Title, err)
	}

	log.Printf("Prix Stripe mis à jour pour '%s'. Ancien: %s, Nouveau: %s", newPrice.Title, currentPriceID, newPriceID)
	return newPriceID, nil
}

// UpdateOptionPrice recrée une option Stripe existante (nouveau prix + désactivation de l'ancien)
func (s *StripeService) UpdateOptionPrice(productID string, currentPriceID string, newOption models.EventOption) (string, error) {
	// Vérification préalable de l'ancien prix
	if currentPriceID == "" {
		log.Printf("Erreur : PriceID actuel manquant pour l'option '%s'", newOption.Title)
		return "", fmt.Errorf("current PriceID is missing for option '%s'", newOption.Title)
	}

	// Désactiver l'ancien prix
	if err := s.DisablePrice(currentPriceID); err != nil {
		log.Printf("Erreur lors de la désactivation de l'option Stripe '%s' : %v", currentPriceID, err)
		return "", fmt.Errorf("failed to disable old Stripe option price '%s': %w", currentPriceID, err)
	}

	// Créer un nouveau prix avec les nouvelles données
	newPriceID, err := s.CreateOptionPrice(productID, newOption)
	if err != nil {
		log.Printf("Erreur lors de la recréation de l'option Stripe pour '%s': %v", newOption.Title, err)
		return "", fmt.Errorf("failed to create new Stripe price for option '%s': %w", newOption.Title, err)
	}

	log.Printf("Option Stripe mise à jour pour '%s'. Ancien: %s, Nouveau: %s", newOption.Title, currentPriceID, newPriceID)
	return newPriceID, nil
}

// DisablePrice désactive un prix Stripe existant
func (s *StripeService) DisablePrice(priceID string) error {
	_, err := price.Update(priceID, &stripe.PriceParams{Active: stripe.Bool(false)})
	if err != nil {
		log.Printf("Erreur lors de la désactivation du prix Stripe '%s' : %v", priceID, err)
		return fmt.Errorf("failed to disable Stripe price '%s': %w", priceID, err)
	}

	log.Printf("Prix Stripe désactivé avec succès : %s", priceID)
	return nil
}

//
//
//

// /
//
// ///
//
//

// /
//
// /

// CreatePaymentIntent pour un paiement immédiat
func (s *StripeService) CreatePaymentIntent(amount int64, currency string, email string, userId int64, metadata map[string]string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),                           // Montant en centimes
		Currency:           stripe.String(currency),                        // Devise (ex. : "eur")
		ReceiptEmail:       stripe.String(email),                           // Adresse email pour le reçu
		PaymentMethodTypes: stripe.StringSlice([]string{"card", "klarna"}), // Mode de paiement (CB ici)
		Metadata:           metadata,                                       // Ajout des métadonnées

	}

	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Erreur lors de la création du Payment Intent : %v", err)
		return nil, fmt.Errorf("failed to create Payment Intent: %w", err)
	}

	log.Printf("Payment Intent créé avec succès : %s", pi.ID)
	return pi, nil
}

// CreatePaymentIntentWithKlarna pour un paiement en plusieurs fois avec Klarna
func (s *StripeService) CreatePaymentIntentWithKlarna(amount int64, currency string, email string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),                   // Montant en centimes
		Currency:           stripe.String(currency),                // Devise (ex. : "eur")
		ReceiptEmail:       stripe.String(email),                   // Adresse email pour le reçu
		PaymentMethodTypes: stripe.StringSlice([]string{"klarna"}), // Klarna comme méthode de paiement

	}

	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Erreur lors de la création du Payment Intent Klarna : %v", err)
		return nil, fmt.Errorf("failed to create Klarna Payment Intent: %w", err)
	}

	log.Printf("Payment Intent Klarna créé avec succès : %s", pi.ID)
	return pi, nil
}

// ConfirmPaymentIntent pour confirmer manuellement un Payment Intent (si nécessaire)
func (s *StripeService) ConfirmPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentConfirmParams{}
	pi, err := paymentintent.Confirm(paymentIntentID, params)
	if err != nil {
		log.Printf("Erreur lors de la confirmation du Payment Intent : %v", err)
		return nil, fmt.Errorf("failed to confirm Payment Intent: %w", err)
	}

	log.Printf("Payment Intent confirmé avec succès : %s", pi.ID)
	return pi, nil
}

// CancelPaymentIntent pour annuler un Payment Intent
func (s *StripeService) CancelPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	pi, err := paymentintent.Cancel(paymentIntentID, nil)
	if err != nil {
		log.Printf("Erreur lors de l'annulation du Payment Intent : %v", err)
		return nil, fmt.Errorf("failed to cancel Payment Intent: %w", err)
	}

	log.Printf("Payment Intent annulé avec succès : %s", pi.ID)
	return pi, nil
}

// UpdatePriceMetadata met à jour les métadonnées d'un prix Stripe (ex. : stock)
func (s *StripeService) UpdatePriceMetadata(priceID string, newStock int) error {
	params := &stripe.PriceParams{
		Metadata: map[string]string{
			"stock": fmt.Sprintf("%d", newStock), // Mettre à jour le stock
		},
	}

	_, err := price.Update(priceID, params)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour des métadonnées du prix '%s' : %v", priceID, err)
		return fmt.Errorf("failed to update metadata for price '%s': %w", priceID, err)
	}

	log.Printf("Mise à jour réussie des métadonnées du prix '%s' avec un stock de %d", priceID, newStock)
	return nil
}

// // CreatePaymentIntentWithCoins crée un Payment Intent Stripe avec gestion des coins
// func (s *StripeService) CreatePaymentIntentWithCoins(
// 	productID string,
// 	totalAmount int64,
// 	userCoins int64,
// 	coinsToUse int64,
// 	currency string,
// 	description string,
// 	metadata map[string]string,
// ) (string, error) {

// 	// Validation des coins
// 	if coinsToUse > userCoins {
// 		return "", fmt.Errorf("L'utilisateur ne peut pas utiliser plus de coins qu'il n'en possède")
// 	}
// 	if coinsToUse > totalAmount {
// 		return "", fmt.Errorf("Les coins ne peuvent pas excéder le montant total")
// 	}

// 	// Calculer le montant restant
// 	remainingAmount := totalAmount - coinsToUse

// 	// Ajouter les métadonnées liées aux coins
// 	if metadata == nil {
// 		metadata = make(map[string]string)
// 	}
// 	metadata["product_id"] = productID
// 	metadata["coins_used"] = fmt.Sprintf("%d", coinsToUse)

// 	// Créer le Payment Intent
// 	params := &stripe.PaymentIntentParams{
// 		Amount:      stripe.Int64(remainingAmount),
// 		Currency:    stripe.String(currency),
// 		Description: stripe.String(description),
// 		Metadata:    metadata,
// 	}

// 	pi, err := paymentintent.New(params)
// 	if err != nil {
// 		log.Printf("Erreur lors de la création du Payment Intent : %v", err)
// 		return "", fmt.Errorf("Erreur lors de la création du Payment Intent : %w", err)
// 	}

// 	log.Printf("Payment Intent créé avec succès : %s", pi.ID)
// 	return pi.ID, nil
// }
