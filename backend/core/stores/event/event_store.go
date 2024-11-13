package event

import (
	"backend/database/models"
	"errors"

	"gorm.io/gorm"
)

type EventStoreType struct {
	db *gorm.DB
}

func EventStore(db *gorm.DB) *EventStoreType {
	return &EventStoreType{db: db}
}

// Accéder à la base de données depuis le store
func (s *EventStoreType) GetDB() *gorm.DB {
	return s.db
}

func (s *EventStoreType) Create(event *models.Event) error {
	return s.db.Create(event).Error
}

func (s *EventStoreType) Update(event *models.Event) error {
	return s.db.Save(event).Error
}

func (s *EventStoreType) Delete(id uint) error {
	return s.db.Delete(&models.Event{}, id).Error
}

func (s *EventStoreType) DeleteWithTx(eventID uint, tx *gorm.DB) error {
	if err := tx.Delete(&models.Event{}, eventID).Error; err != nil {
		return errors.New("Erreur lors de la suppression de l'événement : " + err.Error())
	}
	return nil
}

func (s *EventStoreType) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := s.db.
		Preload("Categories").
		Preload("Tags").
		Preload("Options").
		Preload("Tarifs").
		Preload("Descriptions").
		First(&event, id).
		Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *EventStoreType) List(page, limit int, category, city string) ([]models.Event, int64, error) {
	var events []models.Event
	var total int64

	query := s.db.Model(&models.Event{}).
		Limit(limit).
		Offset((page - 1) * limit).
		Preload("Categories.Category").
		Preload("Tags.Tag")

	if category != "" {
		query = query.Joins("JOIN event_categories ON events.id = event_categories.event_id").
			Joins("JOIN categories ON event_categories.event_category_id = categories.id").
			Where("categories.name = ?", category)
	}
	if city != "" {
		query = query.Where("city = ?", city)
	}

	err := query.Find(&events).Count(&total).Error
	return events, total, err
}

func (s *EventStoreType) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.db.
		Joins("JOIN event_likes ON events.id = event_likes.event_id").
		Where("event_likes.user_id = ?", userID).
		Preload("Categories.Category").
		Preload("Tags.Tag").
		Find(&likedEvents).
		Error
	return likedEvents, err
}

func (s *EventStoreType) GetPopularEvents(limit int) ([]models.Event, error) {
	var popularEvents []models.Event
	err := s.db.
		Joins("LEFT JOIN event_likes ON events.id = event_likes.event_id").
		Select("events.*, COUNT(event_likes.id) as like_count").
		Group("events.id").
		Order("like_count DESC").
		Limit(limit).
		Preload("Tags.Tag").
		Preload("Categories.Category").
		Find(&popularEvents).
		Error
	return popularEvents, err
}

func (s *EventStoreType) GetRecommendedEventsByTags(tags, categories []string, excludedEventIDs []uint, limit int) ([]models.Event, error) {
	var recommendedEvents []models.Event
	err := s.db.
		Joins("JOIN event_event_tags ON events.id = event_event_tags.event_id").
		Joins("JOIN event_tags ON event_event_tags.event_tag_id = event_tags.id").
		Joins("JOIN event_event_categories ON events.id = event_event_categories.event_id").
		Joins("JOIN event_categories ON event_event_categories.event_category_id = event_categories.id").
		Where("event_tags.name IN (?) OR event_categories.name IN (?)", tags, categories).
		Where("events.id NOT IN (?)", excludedEventIDs).
		Group("events.id").
		Limit(limit).
		Preload("Tags").
		Preload("Categories").
		Find(&recommendedEvents).
		Error
	return recommendedEvents, err
}
