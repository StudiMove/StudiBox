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
func HandleGetEvent(c *gin.Context, eventService *event.EventServiceType) {
	claims, err := utils.GetClaimsFromContext(c)
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
	if err := eventService.Interaction.LogEventView(claims.UserID, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de l'enregistrement de la vue", err))
		return
	}

	// Récupère l'événement
	event, err := eventService.Retrieval.GetEvent(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Événement non trouvé", err))
			return
		}
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		return
	}

	// Utilise la fonction buildEventResponse pour générer la réponse de l'événement
	resp := buildEventResponse(event, eventService)
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événement récupéré avec succès", resp))
}

// HandleListEvents gère la récupération paginée des événements
func HandleListEvents(c *gin.Context, eventService *event.EventServiceType) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	category := c.DefaultQuery("category", "")
	city := c.DefaultQuery("city", "")

	// Récupère les événements avec pagination
	events, total, err := eventService.Retrieval.ListEvents(page, limit, category, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la récupération des événements", err))
		return
	}

	eventResponses := make([]response.EventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = buildEventResponse(&event, eventService)
	}

	resp := response.ListEventsResponse{
		Events: eventResponses,
		Page:   page,
		Total:  total,
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Liste des événements récupérée avec succès", resp))
}
