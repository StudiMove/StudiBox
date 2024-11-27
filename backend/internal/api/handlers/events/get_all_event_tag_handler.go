package events

import (
	// Import des ODT pour structurer les réponses
	"backend/internal/services/event"
	"encoding/json"
	"log"
	"net/http"
)

type GetAllTagsHandler struct {
	EventService *event.EventService
}

func NewGetAllTagsHandler(eventService *event.EventService) *GetAllTagsHandler {
	return &GetAllTagsHandler{EventService: eventService}
}

func (h *GetAllTagsHandler) HandleGetAllTags(w http.ResponseWriter, r *http.Request) {
	// Récupérer tous les tags via le service
	tags, err := h.EventService.GetAllTags()
	if err != nil {
		log.Printf("Error fetching tags: %v", err)
		http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
		return
	}

	// Envoyer les tags dans la réponse JSON en utilisant l'ODT
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tags)
}
