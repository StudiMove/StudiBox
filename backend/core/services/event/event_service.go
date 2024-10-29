package event

import (
	stores "backend/core/stores/event"

	"gorm.io/gorm"
)

// EventService regroupe les services de gestion, récupération et interactions des événements
type EventService struct {
	Management  *EventManagementService
	Retrieval   *EventRetrievalService
	Interaction *EventInteractionService
}

// NewEventService crée une nouvelle instance de EventService avec ses sous-services
func NewEventService(db *gorm.DB) *EventService {
	eventStore := stores.NewEventStore(db)
	tagStore := stores.NewTagStore(db)
	categoryStore := stores.NewCategoryStore(db)
	eventLikeStore := stores.NewEventLikeStore(db)
	eventViewStore := stores.NewEventViewStore(db)

	return &EventService{
		Management:  NewEventManagementService(eventStore, tagStore, categoryStore),
		Retrieval:   NewEventRetrievalService(eventStore, eventLikeStore),
		Interaction: NewEventInteractionService(eventStore, eventLikeStore, eventViewStore),
	}
}
