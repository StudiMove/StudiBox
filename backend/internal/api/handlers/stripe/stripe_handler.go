package stripehandler

import (
	"backend/internal/db/models"
	"backend/internal/services/stripe"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// StripeHandler gère les routes liées à Stripe
type StripeHandler struct {
	Service *stripe.StripeService
}

// NewStripeHandler initialise le handler pour Stripe
func NewStripeHandler(service *stripe.StripeService) *StripeHandler {
	return &StripeHandler{Service: service}
}

// CreateProductFromEventHandler gère la création d'un produit Stripe pour un événement
func (h *StripeHandler) CreateProductFromEventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	// Récupérer les données de l'événement depuis la requête
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Appeler le service Stripe pour créer un produit
	productID, err := h.Service.CreateProductFromEvent(&event)
	if err != nil {
		http.Error(w, "Failed to create product in Stripe", http.StatusInternalServerError)
		return
	}

	// Répondre avec le productID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"product_id": productID})
}

// CreatePriceFromTarifHandler gère la création d'un tarif Stripe pour un produit existant
func (h *StripeHandler) CreatePriceFromTarifHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProductID string            `json:"product_id"`
		Tarif     models.EventTarif `json:"tarif"`
	}

	// Récupérer les données depuis la requête
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Appeler le service Stripe pour créer un tarif
	priceID, err := h.Service.CreatePriceFromTarif(req.ProductID, req.Tarif)
	if err != nil {
		http.Error(w, "Failed to create price in Stripe", http.StatusInternalServerError)
		return
	}

	// Répondre avec le priceID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"price_id": priceID})
}

// CreatePaymentIntentHandler crée un PaymentIntent Stripe
// func (h *StripeHandler) CreatePaymentIntentHandler(w http.ResponseWriter, r *http.Request) {
// 	var req struct {
// 		Amount      int64  `json:"amount"`
// 		Currency    string `json:"currency"`
// 		Description string `json:"description"`
// 	}

// 	// Récupérer les données depuis la requête
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	// Appeler le service Stripe pour créer un PaymentIntent
// 	pi, err := h.Service.CreatePaymentIntent(req.Amount, req.Currency, req.Description)
// 	if err != nil {
// 		http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
// 		return
// 	}

//		// Répondre avec le PaymentIntent
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(pi)
//	}

func (h *StripeHandler) CreatePaymentIntentHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID        int64    `json:"user_id"` // ID de l'utilisateur
		Amount        int64    `json:"amount"`
		Currency      string   `json:"currency"`
		Description   string   `json:"description"`
		Email         string   `json:"email"`
		ProductID     string   `json:"product_id"`               // ID du produit
		TarifIDs      []string `json:"tarif_ids"`                // Liste des IDs des tarifs sélectionnés
		OptionsIds    []string `json:"options_ids"`              // Liste des IDs des options sélectionnées
		StudiboxCoins int64    `json:"studibox_coins,omitempty"` // Nombre de StudiboxCoins utilisés

	}

	// Décoder la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Erreur lors du décodage de la requête JSON : %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Assurez-vous que StudiboxCoins a une valeur par défaut
	if req.StudiboxCoins == 0 {
		req.StudiboxCoins = 0
	}

	// Construire les métadonnées
	metadata := map[string]string{
		"user_id":         fmt.Sprintf("%d", req.UserID), // Ajouter l'ID utilisateur
		"product_id":      req.ProductID,
		"selected_tarifs": strings.Join(req.TarifIDs, ","), // Convertir la liste en une chaîne séparée par des virgules
		"description":     req.Description,
		"studibox_coins":  fmt.Sprintf("%d", req.StudiboxCoins), // Ajouter les StudiboxCoins aux métadonnées
	}

	// Ajouter les options dans les métadonnées
	if len(req.OptionsIds) > 0 {
		metadata["selected_options"] = strings.Join(req.OptionsIds, ",")
	}

	// Appeler le service Stripe pour créer un PaymentIntent
	pi, err := h.Service.CreatePaymentIntent(req.Amount, req.Currency, req.Email, req.UserID, metadata)
	if err != nil {
		log.Printf("Erreur lors de la création du PaymentIntent : %v", err)
		http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
		return
	}

	// Répondre avec le PaymentIntent
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pi); err != nil {
		log.Printf("Erreur lors de l'encodage de la réponse JSON : %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// ConfirmPaymentIntentHandler confirme un PaymentIntent Stripe
func (h *StripeHandler) ConfirmPaymentIntentHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}

	// Récupérer les données depuis la requête
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Appeler le service Stripe pour confirmer le PaymentIntent
	pi, err := h.Service.ConfirmPaymentIntent(req.PaymentIntentID)
	if err != nil {
		http.Error(w, "Failed to confirm payment intent", http.StatusInternalServerError)
		return
	}

	// Répondre avec le PaymentIntent confirmé
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pi)
}
