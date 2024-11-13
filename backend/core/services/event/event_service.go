package event

import (
	stores "backend/core/stores/event"

	"gorm.io/gorm"
)

// EventService regroupe les services de gestion, récupération et interactions des événements
type EventServiceType struct {
	Management  *EventManagementServiceType
	Retrieval   *EventRetrievalServiceType
	Interaction *EventInteractionServiceType
}

// EventService crée une nouvelle instance de EventService avec ses sous-services
func EventService(db *gorm.DB) *EventServiceType {
	// Initialiser les stores
	eventStore := stores.EventStore(db)
	eventTagStore := stores.EventTagStore(db)
	eventCategoryStore := stores.EventCategoryStore(db)
	eventLikeStore := stores.EventLikeStore(db)
	eventViewStore := stores.EventViewStore(db)
	eventTarifStore := stores.EventTarifStore(db)
	eventOptionStore := stores.EventOptionStore(db)
	eventDescriptionStore := stores.EventDescriptionStore(db)

	// Initialiser les services avec les stores appropriés
	managementService := EventManagementService(
		eventStore,
		eventTagStore,
		eventCategoryStore,
		eventTarifStore,
		eventOptionStore,
		eventDescriptionStore,
	)

	retrievalService := EventRetrievalService(
		eventStore,
		eventLikeStore,
		eventDescriptionStore,
		eventOptionStore,
		eventTarifStore,
		eventCategoryStore,
		eventTagStore,
	)

	interactionService := EventInteractionService(
		eventStore,
		eventLikeStore,
		eventViewStore,
		eventTarifStore,
		eventOptionStore,
		eventDescriptionStore,
		eventTagStore,
		eventCategoryStore,
	)

	return &EventServiceType{
		Management:  managementService,
		Retrieval:   retrievalService,
		Interaction: interactionService,
	}
}
