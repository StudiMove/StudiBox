package routes

import (
	webhookhandler "backend/internal/api/handlers/webhook"
	"backend/internal/services/webhook"

	"github.com/gorilla/mux"
)

// RegisterWebhookRoutes configure les routes pour gérer les webhooks
func RegisterWebhookRoutes(router *mux.Router, webhookService *webhook.WebhookService) {
	handler := webhookhandler.NewWebhookHandler(webhookService)

	// Route pour Stripe Webhooks
	router.HandleFunc("/stripe", handler.HandleStripeWebhook).Methods("POST", "OPTIONS")
}
