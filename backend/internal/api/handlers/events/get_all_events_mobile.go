package events

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/internal/api/models/event/response"
	"backend/internal/db/models"
	"backend/internal/services/event"
)

type GetAllEventsMobileHandler struct {
	eventService *event.EventService
}

func NewGetAllEventsMobileHandler(eventService *event.EventService) *GetAllEventsMobileHandler {
	return &GetAllEventsMobileHandler{eventService: eventService}
}

// HandleGetAllEventsMobile gère les requêtes pour récupérer tous les événements pour mobile
func (h *GetAllEventsMobileHandler) HandleGetAllEventsMobile(w http.ResponseWriter, r *http.Request) {
	// Récupérer tous les événements via GetAllEventsMobile
	events, err := h.eventService.GetAllEventsMobile()
	if err != nil {
		log.Printf("Failed to retrieve events: %v", err)
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	// Transformer chaque événement en `EventDetailResponse`
	eventResponses := make([]response.EventDetailResponse, len(events))
	for i, event := range events {
		eventResponses[i] = h.transformEventToDetailResponse(event)
	}

	// Retourner la liste complète des événements en JSON pour mobile
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventResponses)
}

// transformEventToDetailResponse transforme un modèle Event en réponse EventDetailResponse
func (h *GetAllEventsMobileHandler) transformEventToDetailResponse(event models.Event) response.EventDetailResponse {
	return response.EventDetailResponse{
		ID:                 event.ID,
		Title:              event.Title,
		Subtitle:           event.Subtitle,
		StartDate:          event.StartDate,
		EndDate:            event.EndDate,
		StartTime:          event.StartTime,
		EndTime:            event.EndTime,
		IsOnline:           event.IsOnline,
		TicketPrice:        event.TicketPrice,
		TicketStock:        event.TicketStock,
		Address:            event.Address,
		City:               event.City,
		Postcode:           event.Postcode,
		Region:             event.Region,
		Country:            event.Country,
		Categories:         h.extractCategoryNames(event.Categories),
		Tags:               h.extractTagNames(event.Tags),
		Options:            h.transformOptions(event.Options),
		Tarifs:             h.transformTarifs(event.Tarifs),
		Descriptions:       h.transformDescriptions(event.Descriptions),
		ImageURLs:          event.ImageURLs,
		IsValidatedByAdmin: event.IsValidatedByAdmin,
		UseStudibox:        event.UseStudibox,
		VideoURL:           event.VideoURL,
	}
}

// extractCategoryNames extrait les noms des catégories d'un événement
func (h *GetAllEventsMobileHandler) extractCategoryNames(categories []models.EventCategory) []string {
	var names []string
	for _, category := range categories {
		names = append(names, category.Name)
	}
	return names
}

// extractTagNames extrait les noms des tags d'un événement
func (h *GetAllEventsMobileHandler) extractTagNames(tags []models.EventTag) []string {
	var names []string
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

// transformOptions transforme les options de l'événement en réponse JSON
func (h *GetAllEventsMobileHandler) transformOptions(options []models.EventOption) []response.EventOptionResponse {
	var optionResponses []response.EventOptionResponse
	for _, option := range options {
		optionResponses = append(optionResponses, response.EventOptionResponse{
			ID:          option.ID,
			Title:       option.Title,
			Description: option.Description,
			Price:       option.Price,
			Stock:       option.Stock,
			PriceID:     option.PriceID,
		})
	}
	return optionResponses
}

// transformTarifs transforme les tarifs de l'événement en réponse JSON
func (h *GetAllEventsMobileHandler) transformTarifs(tarifs []models.EventTarif) []response.EventTarifResponse {
	var tarifResponses []response.EventTarifResponse
	for _, tarif := range tarifs {
		tarifResponses = append(tarifResponses, response.EventTarifResponse{
			ID:          tarif.ID,
			Title:       tarif.Title,
			Description: tarif.Description,
			Price:       tarif.Price,
			Stock:       tarif.Stock,
			PriceID:     tarif.PriceID,
		})
	}
	return tarifResponses
}

// transformDescriptions transforme les descriptions de l'événement en réponse JSON
func (h *GetAllEventsMobileHandler) transformDescriptions(descriptions []models.EventDescription) []response.EventDescriptionResponse {
	var descriptionResponses []response.EventDescriptionResponse
	for _, description := range descriptions {
		descriptionResponses = append(descriptionResponses, response.EventDescriptionResponse{
			Title:       description.Title,
			Description: description.Description,
		})
	}
	return descriptionResponses
}
