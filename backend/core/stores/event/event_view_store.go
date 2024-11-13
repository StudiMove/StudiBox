package event

import (
	"backend/database/models"

	"gorm.io/gorm"
)

type EventViewStoreType struct {
	db *gorm.DB
}

func EventViewStore(db *gorm.DB) *EventViewStoreType {
	return &EventViewStoreType{db: db}
}

// IncrementViewCount ajoute une vue ou incrémente le compteur de vues pour un événement
func (s *EventViewStoreType) IncrementViewCount(eventView *models.EventView) error {
	eventView.Count++
	return s.db.Save(eventView).Error
}

// CreateEventView crée une nouvelle vue pour un utilisateur et un événement spécifiques
func (s *EventViewStoreType) CreateEventView(userID, eventID uint) error {
	eventView := models.EventView{UserID: userID, EventID: eventID, Count: 1}
	return s.db.Create(&eventView).Error
}

// FindEventView recherche une vue d'événement pour un utilisateur et un événement spécifiques
func (s *EventViewStoreType) FindEventView(userID, eventID uint) (*models.EventView, error) {
	var eventView models.EventView
	err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&eventView).Error
	if err != nil {
		return nil, err
	}
	return &eventView, nil
}

// CountViews compte les vues d'un événement
func (s *EventViewStoreType) CountViews(eventID uint) (int, error) {
	var count int64
	err := s.db.Model(&models.EventView{}).Where("event_id = ?", eventID).Count(&count).Error
	return int(count), err
}

// GetMostViewedEventsByUser récupère les événements les plus vus par un utilisateur
func (s *EventViewStoreType) GetMostViewedEventsByUser(userID uint, limit int) ([]models.Event, error) {
	var events []models.Event
	err := s.db.Joins("JOIN event_views ON events.id = event_views.event_id").
		Where("event_views.user_id = ?", userID).
		Order("event_views.count DESC").
		Limit(limit).
		Preload("Tags").Preload("Categories").
		Find(&events).Error
	return events, err
}
