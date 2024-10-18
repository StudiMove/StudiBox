package events

import (
	request "backend/core/api/request/event"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrée invalide", "details": err.Error()})
		return
	}

	// Retrieve the owner's ID from the JWT claims
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Jeton invalide ou expiré"})
		return
	}

	// Set the OwnerID from the JWT claims (UserID)
	createEventReq.OwnerID = claims.UserID

	// Map the request to the Event model
	event := models.Event{
		OwnerID:     createEventReq.OwnerID,
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
		Category:    createEventReq.Category,
		Tags:        createEventReq.Tags,
	}

	// Call the service to create the event
	if err := eventService.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de l'événement", "details": err.Error()})
		return
	}

	// Create the response
	resp := response.CreateEventResponse{
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
		Category:    event.Category,
		Tags:        event.Tags,
	}

	// Return the response
	c.JSON(http.StatusCreated, resp)
}

// HandleUpdateEvent gère la mise à jour d'un événement
func HandleUpdateEvent(c *gin.Context, eventService *event.EventService) {
	id, err := strconv.Atoi(c.Param("id")) // Extract the event ID from the URL
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'événement invalide", "détails": err.Error()})
		return
	}

	var updateEventReq request.UpdateEventRequest
	if err := c.ShouldBindJSON(&updateEventReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrée invalide", "détails": err.Error()})
		return
	}

	// Retrieve the owner's ID from the JWT claims
	claims, err := utils.GetClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Jeton invalide ou expiré"})
		return
	}

	event, err := eventService.GetEvent(uint(id)) // Fetch the event by ID
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Événement non trouvé", "détails": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne lors de la récupération de l'événement", "détails": err.Error()})
		return
	}

	// Ensure that only the owner of the event can update it
	if event.OwnerID != claims.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Vous n'êtes pas autorisé à modifier cet événement"})
		return
	}

	// Update event fields
	event.OwnerType = updateEventReq.OwnerType
	event.ImageURLs = updateEventReq.ImageURLs
	event.VideoURL = updateEventReq.VideoURL
	event.Title = updateEventReq.Title
	event.Subtitle = updateEventReq.Subtitle
	event.Description = updateEventReq.Description
	event.StartDate = updateEventReq.StartDate
	event.EndDate = updateEventReq.EndDate
	event.StartTime = updateEventReq.StartTime
	event.EndTime = updateEventReq.EndTime
	event.IsOnline = updateEventReq.IsOnline
	event.IsVisible = updateEventReq.IsVisible
	event.Price = updateEventReq.Price
	event.Address = updateEventReq.Address
	event.City = updateEventReq.City
	event.Postcode = updateEventReq.Postcode
	event.Region = updateEventReq.Region
	event.Country = updateEventReq.Country
	event.Category = updateEventReq.Category
	event.Tags = updateEventReq.Tags

	// Call the service to update the event
	if err := eventService.UpdateEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour de l'événement", "détails": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Événement mis à jour avec succès"})
}

// HandleDeleteEvent gère la suppression d'un événement
func HandleDeleteEvent(c *gin.Context, eventService *event.EventService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'événement invalide", "détails": err.Error()})
		return
	}

	if err := eventService.DeleteEvent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression de l'événement", "détails": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Événement supprimé avec succès"})
}

// HandleGetEvent gère la récupération d'un événement par ID
func HandleGetEvent(c *gin.Context, eventService *event.EventService) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de l'événement invalide", "détails": err.Error()})
		return
	}

	event, err := eventService.GetEvent(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Événement non trouvé", "détails": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur interne lors de la récupération de l'événement", "détails": err.Error()})
		return
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
		Category:    event.Category,
		Tags:        event.Tags,
	}

	c.JSON(http.StatusOK, resp)
}
