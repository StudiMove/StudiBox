package event

import (
	stores "backend/core/stores/event"
	"backend/database/models"
	"encoding/json"
	"fmt"
)

type EventRetrievalServiceType struct {
	eventStore            *stores.EventStoreType
	eventLikeStore        *stores.EventLikeStoreType
	eventDescriptionStore *stores.EventDescriptionStoreType
	eventOptionStore      *stores.EventOptionStoreType
	eventTarifStore       *stores.EventTarifStoreType
	eventCategoryStore    *stores.EventCategoryStoreType
	eventTagStore         *stores.EventTagStoreType
}

// NewEventRetrievalService initialise un nouveau service de récupération d'événements
func EventRetrievalService(
	eventStore *stores.EventStoreType,
	eventLikeStore *stores.EventLikeStoreType,
	eventDescriptionStore *stores.EventDescriptionStoreType,
	eventOptionStore *stores.EventOptionStoreType,
	eventTarifStore *stores.EventTarifStoreType,
	eventCategoryStore *stores.EventCategoryStoreType,
	eventTagStore *stores.EventTagStoreType,
) *EventRetrievalServiceType {
	return &EventRetrievalServiceType{
		eventStore:            eventStore,
		eventLikeStore:        eventLikeStore,
		eventDescriptionStore: eventDescriptionStore,
		eventOptionStore:      eventOptionStore,
		eventTarifStore:       eventTarifStore,
		eventCategoryStore:    eventCategoryStore,
		eventTagStore:         eventTagStore,
	}
}

// GetEvent récupère un événement par son ID avec tous les détails
func (s *EventRetrievalServiceType) GetEvent(id uint) (*models.Event, error) {
	event, err := s.eventStore.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération de l'événement : %w", err)
	}

	// Conversion de l'URL de l'image en JSON si nécessaire
	if err := json.Unmarshal([]byte(event.ImageURL), &event.ImageURL); err != nil {
		event.ImageURL = string(event.ImageURL)
	}

	// Utiliser la méthode populateEventDetails pour récupérer les détails supplémentaires
	if err := s.populateEventDetails(event); err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des détails de l'événement : %w", err)
	}

	return event, nil
}

// ListEvents récupère une liste d'événements paginée, filtrée par catégorie et ville
func (s *EventRetrievalServiceType) ListEvents(page, limit int, category, city string) ([]models.Event, int64, error) {
	events, total, err := s.eventStore.List(page, limit, category, city)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de la récupération des événements : %w", err)
	}

	// Récupérer les détails supplémentaires pour chaque événement
	for i := range events {
		if err := s.populateEventDetails(&events[i]); err != nil {
			return nil, 0, fmt.Errorf("erreur lors de la récupération des détails de l'événement ID %d : %w", events[i].ID, err)
		}
	}

	return events, total, nil
}

// GetLikedEventsByUser récupère les événements likés par un utilisateur spécifique avec tous les détails
func (s *EventRetrievalServiceType) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	// Récupérer les événements likés
	likedEvents, err := s.eventLikeStore.GetLikedEventsByUser(userID)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des événements likés par l'utilisateur : %w", err)
	}

	// Enrichir chaque événement avec des détails supplémentaires (s'ils ne sont pas déjà chargés)
	for i := range likedEvents {
		if err := s.populateEventDetails(&likedEvents[i]); err != nil {
			return nil, fmt.Errorf("erreur lors de la récupération des détails de l'événement ID %d : %w", likedEvents[i].ID, err)
		}
	}

	return likedEvents, nil
}

// populateEventDetails récupère les descriptions, options et tarifs pour un événement donné
func (s *EventRetrievalServiceType) populateEventDetails(event *models.Event) error {
	var err error

	// Récupérer les descriptions
	if event.Descriptions, err = s.eventDescriptionStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des descriptions : %w", err)
	}

	// Récupérer les options
	if event.Options, err = s.eventOptionStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des options : %w", err)
	}

	// Récupérer les tarifs
	if event.Tarifs, err = s.eventTarifStore.GetByEventID(event.ID); err != nil {
		return fmt.Errorf("erreur lors de la récupération des tarifs : %w", err)
	}

	return nil
}

func (s *EventRetrievalServiceType) GetCategoryNamesByIDs(categoryIDs []int64) ([]string, error) {
	categories, err := s.eventCategoryStore.GetCategoriesByIDs(categoryIDs)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(categories))
	for i, category := range categories {
		names[i] = category.Name
	}
	return names, nil
}

func (s *EventRetrievalServiceType) GetTagNamesByIDs(tagIDs []int64) ([]string, error) {
	tags, err := s.eventTagStore.GetTagsByIDs(tagIDs)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names, nil
}
