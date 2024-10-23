// event_service.go
package event

import (
	stores "backend/core/stores/event"

	"gorm.io/gorm"
)

// EventService gère les opérations liées aux événements
type EventService struct {
	eventStore     *stores.EventStore
	eventLikeStore *stores.EventLikeStore
	eventViewStore *stores.EventViewStore
	tagStore       *stores.TagStore
	categoryStore  *stores.CategoryStore
}

// NewEventService crée une nouvelle instance de EventService
func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		eventStore:     stores.NewEventStore(db),
		eventLikeStore: stores.NewEventLikeStore(db),
		eventViewStore: stores.NewEventViewStore(db),
		tagStore:       stores.NewTagStore(db),
		categoryStore:  stores.NewCategoryStore(db),
	}
}
