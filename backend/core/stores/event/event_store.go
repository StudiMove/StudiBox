package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{db: db}
}

// Créer un événement avec les relations de tags et catégories
func (s *EventStore) Create(event *models.Event) error {
	return s.db.Create(event).Error
}

// Mettre à jour un événement
func (s *EventStore) Update(event *models.Event) error {
	return s.db.Save(event).Error
}

// Supprimer un événement
func (s *EventStore) Delete(id uint) error {
	return s.db.Delete(&models.Event{}, id).Error
}

// Récupérer un événement avec les relations
func (s *EventStore) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := s.db.Preload("Categories").Preload("Tags").First(&event, id).Error
	return &event, err
}

// Liste des événements avec filtres
func (s *EventStore) List(page, limit int, category, city string) ([]models.Event, int64, error) {
	var events []models.Event
	var total int64
	query := s.db.Model(&models.Event{}).Limit(limit).Offset((page - 1) * limit)

	if category != "" {
		query = query.Joins("JOIN event_categories ON events.id = event_categories.event_id").Where("event_categories.name = ?", category)
	}
	if city != "" {
		query = query.Where("city = ?", city)
	}

	err := query.Preload("Categories").Preload("Tags").Find(&events).Count(&total).Error
	return events, total, err
}

// GetPopularEvents récupère les événements les plus populaires (par nombre de likes)
func (s *EventStore) GetPopularEvents(limit int) ([]models.Event, error) {
	var popularEvents []models.Event
	err := s.db.
		Joins("LEFT JOIN event_likes ON events.id = event_likes.event_id").
		Select("events.*, COUNT(event_likes.id) as like_count").
		Group("events.id").
		Order("like_count DESC").
		Limit(limit).
		Preload("Tags").Preload("Categories").
		Find(&popularEvents).Error
	return popularEvents, err
}

// GetRecommendedEventsByTags récupère les événements recommandés en fonction des tags et des catégories
func (s *EventStore) GetRecommendedEventsByTags(tags, categories []string, excludedEventIDs []uint, limit int) ([]models.Event, error) {
	var recommendedEvents []models.Event
	err := s.db.
		Joins("JOIN event_tags ON events.id = event_tags.event_id").
		Joins("JOIN tags ON event_tags.tag_id = tags.id").
		Joins("JOIN event_categories ON events.id = event_categories.event_id").
		Joins("JOIN categories ON event_categories.category_id = categories.id").
		Where("tags.name IN (?) OR categories.name IN (?)", tags, categories).
		Where("events.id NOT IN (?)", excludedEventIDs).
		Group("events.id").
		Limit(limit).
		Preload("Tags").Preload("Categories").
		Find(&recommendedEvents).Error
	return recommendedEvents, err
}
