package events

import (
	request "backend/core/api/request/event"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/event"
	"backend/core/models"
	"backend/core/services/event"
	"backend/core/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleCreateEvent gère la création d'un événement
func HandleCreateEvent(c *gin.Context, eventService *event.EventService) {
	var createEventReq request.CreateEventRequest
	if err := c.ShouldBindJSON(&createEventReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Entrée invalide", err))
		return
	}

	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Jeton invalide ou expiré", err))
		return
	}

	event := models.Event{
		OwnerID:     claims.UserID,
		OwnerType:   createEventReq.OwnerType,
		ImageURLs:   createEventReq.ImageURLs,
		VideoURL:    createEventReq.VideoURL,
		Title:       createEventReq.Title,
		Subtitle:    createEventReq.Subtitle,
		Description: createEventReq.Description,
		StartDate:   createEventReq.StartDate,
		EndDate:     createEventReq.EndDate,
		StartTime:   createEventReq.StartTime,
		EndTime:     createEventReq.EndTime,
		IsOnline:    createEventReq.IsOnline,
		IsVisible:   createEventReq.IsVisible,
		Price:       createEventReq.Price,
		Address:     createEventReq.Address,
		City:        createEventReq.City,
		Postcode:    createEventReq.Postcode,
		Region:      createEventReq.Region,
		Country:     createEventReq.Country,
	}

	// Créer l'événement avec les catégories et les tags
	if err := eventService.Management.CreateEvent(&event, createEventReq.Tags, createEventReq.Category); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la création de l'événement", err))
		return
	}

	resp := response.EventResponse{
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
		Categories:  createEventReq.Category,
		Tags:        createEventReq.Tags,
	}

	c.JSON(http.StatusCreated, responseGlobal.SuccessResponse("Événement créé avec succès", resp))
}

// HandleUpdateEvent gère la mise à jour d'un événement
func HandleUpdateEvent(c *gin.Context, eventService *event.EventService) {
	// Conversion de l'ID d'événement depuis le paramètre URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID de l'événement invalide", err))
		return
	}

	// Lier la requête JSON à la structure de mise à jour d'événement
	var updateEventReq request.UpdateEventRequest
	if err := c.ShouldBindJSON(&updateEventReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Entrée invalide", err))
		return
	}

	// Récupérer l'événement par ID
	event, err := eventService.Retrieval.GetEvent(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Événement non trouvé", err))
			return
		}
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		return
	}

	// Vérifier les droits de l'utilisateur
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil || event.OwnerID != claims.UserID {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Non autorisé à modifier cet événement", nil))
		return
	}

	// Mettre à jour l'événement en utilisant la requête
	if err := eventService.Management.UpdateEvent(event, updateEventReq); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la mise à jour", err))
		return
	}

	// Réponse après la mise à jour de l'événement
	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événement mis à jour avec succès", &response.EventResponse{
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
		Categories:  updateEventReq.Category,
		Tags:        updateEventReq.Tags,
	}))
}

// HandleDeleteEvent gère la suppression d'un événement
func HandleDeleteEvent(c *gin.Context, eventService *event.EventService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("ID de l'événement invalide", err))
		return
	}

	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Jeton invalide ou expiré", err))
		return
	}

	event, err := eventService.Retrieval.GetEvent(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responseGlobal.ErrorResponse("Événement non trouvé", err))
			return
		}
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur interne", err))
		return
	}

	if event.OwnerID != claims.UserID {
		c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Non autorisé à supprimer cet événement", nil))
		return
	}

	// Supprimer l'événement
	if err := eventService.Management.DeleteEvent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Erreur lors de la suppression", err))
		return
	}

	c.JSON(http.StatusOK, responseGlobal.SuccessResponse("Événement supprimé avec succès", nil))
}
