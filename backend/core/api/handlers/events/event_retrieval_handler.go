package events

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/event"
	"backend/core/services/event"
	"backend/core/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleGetEvent gère la récupération d'un événement par ID et enregistre une vue
func HandleGetEvent(c *gin.Context, eventService *event.EventService) {
	claims, err := utils.GetClaimsFromContext(c) // Récupère l'utilisateur connecté
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Accès refusé : Token invalide", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID de l'événement invalide", err))
		return
	}

	// Enregistre la vue de l'événement
	if err := eventService.LogEventView(claims.UserID, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de l'enregistrement de la vue", err))
		return
	}

	// Récupère l'événement
	event, err := eventService.GetEvent(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Événement non trouvé", err))
			return
		}
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		return
	}

	// Conversion des catégories et tags en []string
	categories := make([]string, len(event.Categories))
	for i, category := range event.Categories {
		categories[i] = category.Name
	}

	tags := make([]string, len(event.Tags))
	for i, tag := range event.Tags {
		tags[i] = tag.Name
	}

	resp := response.GetEventResponse{
		ID:          event.ID,
		OwnerID:     event.OwnerID,
		OwnerType:   event.OwnerType,
		ImageURLs:   event.ImageURLs,
		VideoURL:    event.VideoURL,
		Title:       event.Title,
		Subtitle:    event.Subtitle,
		Description: event.Description,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		IsOnline:    event.IsOnline,
		IsVisible:   event.IsVisible,
		Price:       event.Price,
		Address:     event.Address,
		City:        event.City,
		Postcode:    event.Postcode,
		Region:      event.Region,
		Country:     event.Country,
		Categories:  categories, // Conversion en []string
		Tags:        tags,       // Conversion en []string
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événement récupéré avec succès", resp))
}

// HandleListEvents gère la récupération paginée des événements
func HandleListEvents(c *gin.Context, eventService *event.EventService) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	category := c.DefaultQuery("category", "")
	city := c.DefaultQuery("city", "")

	events, total, err := eventService.ListEvents(page, limit, category, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la récupération des événements", err))
		return
	}

	var eventResponses []response.GetEventResponse
	for _, event := range events {
		// Conversion des catégories et tags en []string
		categories := make([]string, len(event.Categories))
		for i, category := range event.Categories {
			categories[i] = category.Name
		}

		tags := make([]string, len(event.Tags))
		for i, tag := range event.Tags {
			tags[i] = tag.Name
		}

		eventResp := response.GetEventResponse{
			ID:          event.ID,
			OwnerID:     event.OwnerID,
			OwnerType:   event.OwnerType,
			ImageURLs:   event.ImageURLs,
			VideoURL:    event.VideoURL,
			Title:       event.Title,
			Subtitle:    event.Subtitle,
			Description: event.Description,
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			IsOnline:    event.IsOnline,
			IsVisible:   event.IsVisible,
			Price:       event.Price,
			Address:     event.Address,
			City:        event.City,
			Postcode:    event.Postcode,
			Region:      event.Region,
			Country:     event.Country,
			Categories:  categories, // Conversion en []string
			Tags:        tags,       // Conversion en []string
		}
		eventResponses = append(eventResponses, eventResp)
	}

	resp := response.ListEventsResponse{
		Events: eventResponses,
		Page:   page,
		Total:  total,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Liste des événements récupérée avec succès", resp))
}
