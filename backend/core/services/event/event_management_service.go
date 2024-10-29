package event

import (
	request "backend/core/api/request/event"
	"backend/core/models"
	stores "backend/core/stores/event"
	"encoding/json"
	"errors"
)

type EventManagementService struct {
	eventStore    *stores.EventStore
	tagStore      *stores.TagStore
	categoryStore *stores.CategoryStore
}

func NewEventManagementService(eventStore *stores.EventStore, tagStore *stores.TagStore, categoryStore *stores.CategoryStore) *EventManagementService {
	return &EventManagementService{
		eventStore:    eventStore,
		tagStore:      tagStore,
		categoryStore: categoryStore,
	}
}

func (s *EventManagementService) CreateEvent(event *models.Event, tagNames, categoryNames []string) error {
	if err := s.assignTagsAndCategories(event, tagNames, categoryNames); err != nil {
		return err
	}

	imageURLsJSON, err := json.Marshal(event.ImageURLs)
	if err != nil {
		return errors.New("erreur de conversion des URLs d'images en JSON: " + err.Error())
	}
	event.ImageURLsJSON = string(imageURLsJSON)

	return s.eventStore.Create(event)
}

func (s *EventManagementService) UpdateEvent(event *models.Event, input request.UpdateEventRequest) error {
	fieldsToUpdate := map[string]interface{}{
		"OwnerType":   input.OwnerType,
		"Title":       input.Title,
		"Subtitle":    input.Subtitle,
		"Description": input.Description,
		"StartDate":   input.StartDate,
		"EndDate":     input.EndDate,
		"StartTime":   input.StartTime,
		"EndTime":     input.EndTime,
		"Price":       input.Price,
		"Address":     input.Address,
		"City":        input.City,
		"Postcode":    input.Postcode,
		"Region":      input.Region,
		"Country":     input.Country,
		"VideoURL":    input.VideoURL,
	}

	for field, value := range fieldsToUpdate {
		if v, ok := value.(string); ok && v != "" {
			// Mise à jour des champs conditionnels
			switch field {
			case "OwnerType":
				event.OwnerType = v
			case "Title":
				event.Title = v
			case "Subtitle":
				event.Subtitle = v
			case "Description":
				event.Description = v
			case "StartDate":
				event.StartDate = v
			case "EndDate":
				event.EndDate = v
			case "StartTime":
				event.StartTime = v
			case "EndTime":
				event.EndTime = v
			case "Address":
				event.Address = v
			case "City":
				event.City = v
			case "Postcode":
				event.Postcode = v
			case "Region":
				event.Region = v
			case "Country":
				event.Country = v
			case "VideoURL":
				event.VideoURL = v
			}
		} else if p, ok := value.(int); ok && p >= 0 {
			event.Price = p
		}
	}

	if len(input.ImageURLs) > 0 {
		event.ImageURLs = input.ImageURLs
		imageURLsJSON, err := json.Marshal(input.ImageURLs)
		if err != nil {
			return errors.New("erreur de conversion des URLs d'images en JSON: " + err.Error())
		}
		event.ImageURLsJSON = string(imageURLsJSON)
	}

	if input.IsOnline != nil {
		event.IsOnline = *input.IsOnline
	}
	if input.IsVisible != nil {
		event.IsVisible = *input.IsVisible
	}

	if len(input.Tags) > 0 || len(input.Category) > 0 {
		if err := s.assignTagsAndCategories(event, input.Tags, input.Category); err != nil {
			return err
		}
	}

	return s.eventStore.Update(event)
}

func (s *EventManagementService) DeleteEvent(id uint) error {
	if err := s.eventStore.Delete(id); err != nil {
		return errors.New("Erreur lors de la suppression de l'événement: " + err.Error())
	}
	return nil
}

func (s *EventManagementService) assignTagsAndCategories(event *models.Event, tagNames, categoryNames []string) error {
	for _, tagName := range tagNames {
		tag, err := s.tagStore.FindOrCreateTag(tagName)
		if err != nil {
			return err
		}
		event.Tags = append(event.Tags, *tag)
	}

	for _, categoryName := range categoryNames {
		category, err := s.categoryStore.FindOrCreateCategory(categoryName)
		if err != nil {
			return err
		}
		event.Categories = append(event.Categories, *category)
	}

	return nil
}
