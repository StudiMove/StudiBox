// event_management.go
package event

import (
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

// Met à jour un événement existant
func (s *EventService) UpdateEvent(event *models.Event, tagNames, categoryNames []string) error {
	if err := s.assignTagsAndCategories(event, tagNames, categoryNames); err != nil {
		return err
	}

	imageURLsJSON, err := json.Marshal(event.ImageURLs)
	if err != nil {
		return errors.New("erreur de conversion des URLs d'images en JSON: " + err.Error())
	}
	event.ImageURLsJSON = string(imageURLsJSON)

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
