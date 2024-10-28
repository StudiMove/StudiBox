// event_management.go
package event

import (
	request "backend/core/api/request/event"
	"backend/core/models"
	"encoding/json"
	"errors"
)

// Crée un événement
func (s *EventService) CreateEvent(event *models.Event, tagNames, categoryNames []string) error {
	if err := s.assignTagsAndCategories(event, tagNames, categoryNames); err != nil {
		return err
	}

	// Sauvegarder les URLs des images sous format JSON
	imageURLsJSON, err := json.Marshal(event.ImageURLs)
	if err != nil {
		return errors.New("erreur de conversion des URLs d'images en JSON: " + err.Error())
	}
	event.ImageURLsJSON = string(imageURLsJSON)

	return s.eventStore.Create(event)
}

// UpdateEvent met à jour un événement existant de manière conditionnelle
func (s *EventService) UpdateEvent(event *models.Event, input request.UpdateEventRequest) error {
	// Map des champs à mettre à jour
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

	// Appliquer les changements conditionnels
	for field, value := range fieldsToUpdate {
		if v, ok := value.(string); ok && v != "" {
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
		} else if p, ok := value.(int); ok && p >= 0 { // Gestion du prix, en s'assurant qu'il est non-négatif
			event.Price = p
		}
	}

	// Mise à jour des images et conversion en JSON
	if len(input.ImageURLs) > 0 {
		event.ImageURLs = input.ImageURLs
		imageURLsJSON, err := json.Marshal(input.ImageURLs)
		if err != nil {
			return errors.New("erreur de conversion des URLs d'images en JSON: " + err.Error())
		}
		event.ImageURLsJSON = string(imageURLsJSON)
	}

	// Mise à jour des booléens (IsOnline et IsVisible) avec des pointeurs pour gérer la nullabilité
	if input.IsOnline != nil {
		event.IsOnline = *input.IsOnline
	}
	if input.IsVisible != nil {
		event.IsVisible = *input.IsVisible
	}

	// Mise à jour des tags et catégories uniquement si présents
	if len(input.Tags) > 0 || len(input.Category) > 0 {
		if err := s.assignTagsAndCategories(event, input.Tags, input.Category); err != nil {
			return err
		}
	}

	// Mise à jour dans la base de données
	return s.eventStore.Update(event)
}

// Supprime un événement par ID
func (s *EventService) DeleteEvent(id uint) error {
	if err := s.eventStore.Delete(id); err != nil {
		return errors.New("Erreur lors de la suppression de l'événement: " + err.Error())
	}
	return nil
}

// Fonction pour assigner les tags et catégories à un événement
func (s *EventService) assignTagsAndCategories(event *models.Event, tagNames, categoryNames []string) error {
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
