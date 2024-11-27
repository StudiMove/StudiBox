package events

import (
	"backend/internal/services/event"
	"encoding/json"
	"log"
	"net/http"
)

type GetAllCategoriesHandler struct {
	EventService *event.EventService
}

func NewGetAllCategoriesHandler(eventService *event.EventService) *GetAllCategoriesHandler {
	return &GetAllCategoriesHandler{EventService: eventService}
}

func (h *GetAllCategoriesHandler) HandleGetAllCategories(w http.ResponseWriter, r *http.Request) {
    // Récupérer toutes les catégories via le service
    categories, err := h.EventService.GetAllCategories()
    if err != nil {
        log.Printf("Error fetching categories: %v", err)
        http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
        return
    }

    // Transformer les catégories en un tableau de chaînes
    categoryNames := make([]string, len(categories))
    for i, category := range categories {
        categoryNames[i] = category.Name
    }

    // Envoyer la réponse JSON en tant que tableau de chaînes
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(categoryNames)
}

