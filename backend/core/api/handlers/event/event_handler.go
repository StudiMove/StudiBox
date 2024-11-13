package events

import (
	response "backend/core/api/response/event"
	"backend/core/services/event"
	"backend/database/models"
)

func buildEventResponse(event *models.Event, eventService *event.EventServiceType) response.EventResponse {
	// Obtenir le nombre de likes et de vues
	likes, _ := eventService.Interaction.GetLikesCount(event.ID)
	views, _ := eventService.Interaction.GetViewsCount(event.ID)

	// Convertir les CategoryIDs en slice d'entiers
	categoryIDs := make([]int64, len(event.CategoryIDs))
	for i, id := range event.CategoryIDs {
		categoryIDs[i] = int64(id)
	}

	// Récupérer les noms des catégories
	categories, err := eventService.Retrieval.GetCategoryNamesByIDs(categoryIDs)
	if err != nil {
		categories = []string{}
	}

	// Convertir les TagIDs en slice d'entiers
	tagIDs := make([]int64, len(event.TagIDs))
	for i, id := range event.TagIDs {
		tagIDs[i] = int64(id)
	}

	// Récupérer les noms des tags
	tags, err := eventService.Retrieval.GetTagNamesByIDs(tagIDs)
	if err != nil {
		tags = []string{}
	}

	// Préparer les réponses pour les descriptions, options et tarifs
	descriptions := make([]response.DescriptionResponse, len(event.Descriptions))
	for i, desc := range event.Descriptions {
		descriptions[i] = response.DescriptionResponse{
			Title:       desc.Title,
			Description: desc.Description,
		}
	}

	options := make([]response.OptionResponse, len(event.Options))
	for i, opt := range event.Options {
		options[i] = response.OptionResponse{
			Title:       opt.Title,
			Description: opt.Description,
			Price:       opt.Price,
			Stock:       opt.Stock,
		}
	}

	tarifs := make([]response.TarifResponse, len(event.Tarifs))
	for i, tarif := range event.Tarifs {
		tarifs[i] = response.TarifResponse{
			Title:       tarif.Title,
			Description: tarif.Description,
			Price:       tarif.Price,
			Stock:       tarif.Stock,
		}
	}

	// Construire et retourner la réponse
	return response.EventResponse{
		ID:           event.ID,
		OwnerID:      event.OwnerID,
		OwnerType:    event.OwnerType,
		ImageURL:     event.ImageURL,
		VideoURL:     event.VideoURL,
		Title:        event.Title,
		Subtitle:     event.Subtitle,
		StartDate:    event.StartDate,
		EndDate:      event.EndDate,
		StartTime:    event.StartTime,
		EndTime:      event.EndTime,
		IsOnline:     event.IsOnline,
		IsPublic:     event.IsPublic,
		Address:      event.Address,
		City:         event.City,
		PostalCode:   event.PostalCode,
		Region:       event.Region,
		Country:      event.Country,
		Categories:   categories,
		Tags:         tags,
		Descriptions: descriptions,
		Options:      options,
		Tarifs:       tarifs,
		Likes:        likes,
		Views:        views,
	}
}
