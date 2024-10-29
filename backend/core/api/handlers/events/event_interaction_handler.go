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
func HandleGetRecommendations(c *gin.Context, eventService *event.EventService) {
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

	var recommendationResponses []response.EventResponse
	for _, event := range recommendations {
		categories := make([]string, len(event.Categories))
		for i, category := range event.Categories {
			categories[i] = category.Name
		}

		tags := make([]string, len(event.Tags))
		for i, tag := range event.Tags {
			tags[i] = tag.Name
		}

		eventResp := response.EventResponse{
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
			Categories:  categories,
			Tags:        tags,
		}
		recommendationResponses = append(recommendationResponses, eventResp)
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Recommandations récupérées avec succès", recommendationResponses))
}

// HandleLikeEvent gère l'action de liker un événement
func HandleLikeEvent(c *gin.Context, eventService *event.EventService) {
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
func HandleUnlikeEvent(c *gin.Context, eventService *event.EventService) {
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

func HandleGetLikedEvents(c *gin.Context, eventService *event.EventService) {
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

	var eventResponses []response.EventResponse
	for _, event := range likedEvents {
		// Conversion des catégories et tags en []string
		categories := make([]string, len(event.Categories))
		for i, category := range event.Categories {
			categories[i] = category.Name
		}

		tags := make([]string, len(event.Tags))
		for i, tag := range event.Tags {
			tags[i] = tag.Name
		}

		eventResp := response.EventResponse{
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
			Categories:  categories,
			Tags:        tags,
		}
		eventResponses = append(eventResponses, eventResp)
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événements likés récupérés avec succès", eventResponses))
}
