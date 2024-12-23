package routes

import (
	stripehandler "backend/internal/api/handlers/stripe"
	"backend/internal/services/stripe"

	"github.com/gorilla/mux"
)

// RegisterStripeRoutes configure les routes pour Stripe
func RegisterStripeRoutes(router *mux.Router, stripeService *stripe.StripeService) {
	handler := stripehandler.NewStripeHandler(stripeService)

	// Routes Stripe
	router.HandleFunc("/products", handler.CreateProductFromEventHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/prices", handler.CreatePriceFromTarifHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/payment-intents", handler.CreatePaymentIntentHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/payment-intents/confirm", handler.ConfirmPaymentIntentHandler).Methods("POST", "OPTIONS")
}
