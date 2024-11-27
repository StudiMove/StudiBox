package events

import (
	"backend/internal/api/models/event/request"
	"backend/internal/db/models"
	"backend/internal/services/event"
	"encoding/json"
	"net/http"
)

type UpdateEventHandler struct {
	eventService *event.EventService
}

func NewUpdateEventHandler(eventService *event.EventService) *UpdateEventHandler {
	return &UpdateEventHandler{eventService: eventService}
}

func (h *UpdateEventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateEventDataRequest

	// Décodage de la requête
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.EventID == 0 {
		http.Error(w, "Event ID is required", http.StatusBadRequest)
		return
	}

	// Récupérer ou créer les catégories et tags en convertissant vers les types `models`
	categoryModels, err := h.eventService.GetOrCreateCategories(req.Categories)
	if err != nil {
		http.Error(w, "Failed to get or create categories", http.StatusInternalServerError)
		return
	}
	tagModels, err := h.eventService.GetOrCreateTags(req.Tags)
	if err != nil {
		http.Error(w, "Failed to get or create tags", http.StatusInternalServerError)
		return
	}

	// Transformer les catégories et tags en types `models`
	var categories []models.EventCategory
	for _, cat := range categoryModels {
		categories = append(categories, models.EventCategory{Name: cat.Name})
	}

	var tags []models.EventTag
	for _, tag := range tagModels {
		tags = append(tags, models.EventTag{Name: tag.Name})
	}

	// Créer l'objet événement à mettre à jour sans conversion supplémentaire pour `time.Time`
	updatedEvent := models.Event{
		IsOnline:    req.IsOnline,
		IsVisible:   req.IsVisible,
		Title:       req.Title,
		Subtitle:    req.Subtitle,
		StartDate:   req.StartDate,
		StartTime:   req.StartTime,
		EndDate:     req.EndDate,
		EndTime:     req.EndTime,
		UseStudibox: req.UseStudibox,
		TicketPrice: req.TicketPrice,
		TicketStock: req.TicketStock,
		Address:     req.Location.Address,
		City:        req.Location.City,
		Postcode:    req.Location.Postcode,
		Region:      req.Location.Region,
		Country:     req.Location.Country,
		Options:     req.Options,
		Tarifs:      req.Tarifs,
		VideoURL:    req.VideoURL,
		Descriptions: req.Descriptions,
		Categories:   categories,
		Tags:         tags,
		ImageURLs: 		req.Images,	
	}

	// Appel du service pour mettre à jour l'événement
	if err := h.eventService.UpdateEvent(req.EventID, &updatedEvent, req.Categories, req.Tags); err != nil {
		http.Error(w, "Failed to update event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Réponse réussie avec les détails mis à jour
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEvent)
}
