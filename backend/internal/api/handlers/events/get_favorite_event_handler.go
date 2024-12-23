package events

import (
	"encoding/json"
	"log"
	"net/http"

	"backend/internal/api/models/event/response"
	"backend/internal/db/models"
	"backend/internal/services/event"
	"backend/internal/utils"
)

type GetFavoriteEventsHandler struct {
	eventService *event.EventService
}

func NewGetFavoriteEventsHandler(eventService *event.EventService) *GetFavoriteEventsHandler {
	return &GetFavoriteEventsHandler{eventService: eventService}
}

// HandleGetFavoriteEvents gère la requête pour récupérer les événements favoris d'un utilisateur
func (h *GetFavoriteEventsHandler) HandleGetFavoriteEvents(w http.ResponseWriter, r *http.Request) {
	// Récupérer les claims utilisateur depuis le contexte
	userClaims, ok := r.Context().Value("user").(*utils.JWTClaims)
	if !ok || userClaims == nil {
		http.Error(w, "Unauthorized request: invalid or missing JWT claims", http.StatusUnauthorized)
		return
	}

	userID := userClaims.UserID

	// Appeler le service pour récupérer les événements favoris
	events, err := h.eventService.GetFavoriteEventsByUserId(userID)
	if err != nil {
		log.Printf("Failed to retrieve favorite events for user %d: %v", userID, err)
		http.Error(w, "Failed to retrieve favorite events", http.StatusInternalServerError)
		return
	}

	// Transformer chaque événement en `EventDetailResponse`
	eventResponses := make([]response.EventDetailResponse, len(events))
	for i, event := range events {
		eventResponses[i] = h.transformEventToDetailResponse(event)
	}

	// Retourner les événements favoris en JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventResponses)
}

// transformEventToDetailResponse transforme un modèle Event en réponse EventDetailResponse
func (h *GetFavoriteEventsHandler) transformEventToDetailResponse(event models.Event) response.EventDetailResponse {
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
func (h *GetFavoriteEventsHandler) extractCategoryNames(categories []models.EventCategory) []string {
	var names []string
	for _, category := range categories {
		names = append(names, category.Name)
	}
	return names
}

// extractTagNames extrait les noms des tags d'un événement
func (h *GetFavoriteEventsHandler) extractTagNames(tags []models.EventTag) []string {
	var names []string
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

// transformOptions transforme les options de l'événement en réponse JSON
func (h *GetFavoriteEventsHandler) transformOptions(options []models.EventOption) []response.EventOptionResponse {
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
func (h *GetFavoriteEventsHandler) transformTarifs(tarifs []models.EventTarif) []response.EventTarifResponse {
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
func (h *GetFavoriteEventsHandler) transformDescriptions(descriptions []models.EventDescription) []response.EventDescriptionResponse {
	var descriptionResponses []response.EventDescriptionResponse
	for _, description := range descriptions {
		descriptionResponses = append(descriptionResponses, response.EventDescriptionResponse{
			Title:       description.Title,
			Description: description.Description,
		})
	}
	return descriptionResponses
}
