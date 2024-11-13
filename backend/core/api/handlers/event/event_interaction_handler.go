package events

import (
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/event"
	"backend/core/services/event"
	"backend/core/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetRecommendations gère la récupération des recommandations pour l'utilisateur
func HandleGetRecommendations(c *gin.Context, eventService *event.EventServiceType) {
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Accès refusé : Token invalide", err))
		return
	}

	recommendations, err := eventService.Interaction.GetRecommendations(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Impossible de générer les recommandations", err))
		return
	}

	recommendationResponses := make([]response.EventResponse, len(recommendations))
	for i, event := range recommendations {
		recommendationResponses[i] = buildEventResponse(&event, eventService)
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Recommandations récupérées avec succès", recommendationResponses))
}

// HandleLikeEvent gère l'action de liker un événement
func HandleLikeEvent(c *gin.Context, eventService *event.EventServiceType) {
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Accès refusé : Token invalide", err))
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID d'événement invalide", err))
		return
	}

	if err := eventService.Interaction.LikeEvent(claims.UserID, uint(eventID)); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Impossible de liker l'événement", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("L'événement a été liké avec succès", nil))
}

// HandleUnlikeEvent gère l'action de retirer un like d'un événement
func HandleUnlikeEvent(c *gin.Context, eventService *event.EventServiceType) {
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Accès refusé : Token invalide", err))
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID d'événement invalide", err))
		return
	}

	if err := eventService.Interaction.UnlikeEvent(claims.UserID, uint(eventID)); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Impossible de retirer le like", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Le like a été retiré avec succès", nil))
}

// HandleGetLikedEvents gère la récupération des événements likés par l'utilisateur
func HandleGetLikedEvents(c *gin.Context, eventService *event.EventServiceType) {
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Accès refusé : Token invalide", err))
		return
	}

	likedEvents, err := eventService.Retrieval.GetLikedEventsByUser(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Impossible de récupérer les événements likés", err))
		return
	}

	// Convertir les événements likés en réponses structurées
	eventResponses := make([]response.EventResponse, len(likedEvents))
	for i, event := range likedEvents {
		eventResponses[i] = buildEventResponse(&event, eventService)
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événements likés récupérés avec succès", eventResponses))
}
