// package webhookhandler

// import (
// 	"backend/internal/services/webhook"
// 	"io/ioutil"
// 	"net/http"
// )

// // WebhookHandler structure pour gérer les webhooks
// type WebhookHandler struct {
// 	Service *webhook.WebhookService
// }

// // NewWebhookHandler initialise un nouveau WebhookHandler
// func NewWebhookHandler(service *webhook.WebhookService) *WebhookHandler {
// 	return &WebhookHandler{Service: service}
// }

// // HandleStripeWebhook gère les webhooks Stripe
// func (h *WebhookHandler) HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
// 	// Lire le payload brut
// 	payload, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Unable to read request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Récupérer l'en-tête Stripe-Signature
// 	sigHeader := r.Header.Get("Stripe-Signature")

// 	// Appeler le service pour vérifier et traiter le webhook
// 	if err := h.Service.VerifyAndHandleWebhook(payload, sigHeader); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Répondre avec un succès
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "Webhook processed successfully"}`))
// }

package webhookhandler

import (
	"backend/config"
	"backend/internal/services/webhook"
	"io"
	"log"
	"net/http"
)

// WebhookHandler structure pour gérer les webhooks
type WebhookHandler struct {
	Service *webhook.WebhookService
}

// NewWebhookHandler initialise un nouveau WebhookHandler
func NewWebhookHandler(service *webhook.WebhookService) *WebhookHandler {
	return &WebhookHandler{Service: service}
}

// HandleStripeWebhook gère les webhooks Stripe
func (h *WebhookHandler) HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Lire le payload
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Récupérer la signature Stripe
	sigHeader := r.Header.Get("Stripe-Signature")
	log.Printf("Payload: %s", string(payload))
	log.Printf("Signature header: %s", sigHeader)
	log.Printf("Configured webhook secret: %s", config.AppConfig.StripeWebhookSecret)

	// Vérification et traitement
	if err := h.Service.VerifyAndHandleWebhook(payload, sigHeader); err != nil {
		log.Printf("Signature invalide : %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Répondre au webhook
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Webhook processed successfully"}`))
}
