package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EventViewStore struct {
	db *gorm.DB
}

func NewEventViewStore(db *gorm.DB) *EventViewStore {
	return &EventViewStore{db: db}
}

// AddEventView ajoute ou met à jour une vue pour un événement par un utilisateur
func (s *EventViewStore) AddEventView(userID, eventID uint) error {
	var eventView models.EventView

	// Vérifier si l'utilisateur a déjà vu cet événement
	err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&eventView).Error
	if err == nil {
		// Si l'utilisateur a déjà vu l'événement, incrémenter le compteur
		eventView.Count++
		return s.db.Save(&eventView).Error
	} else if err == gorm.ErrRecordNotFound {
		// Si c'est la première vue de l'utilisateur pour cet événement
		eventView = models.EventView{UserID: userID, EventID: eventID, Count: 1}
		return s.db.Create(&eventView).Error
	}
	return err
}

// GetMostViewedEventsByUser récupère les événements les plus vus par un utilisateur
func (s *EventViewStore) GetMostViewedEventsByUser(userID uint, limit int) ([]models.Event, error) {
	var events []models.Event
	err := s.db.Joins("JOIN event_views ON events.id = event_views.event_id").
		Where("event_views.user_id = ?", userID).
		Order("event_views.count DESC").
		Limit(limit).
		Preload("Tags").Preload("Categories").
		Find(&events).Error
	return events, err
}

func (s *EventViewStore) CountViews(eventID uint) (int, error) {
	var count int64
	err := s.db.Model(&models.EventView{}).Where("event_id = ?", eventID).Count(&count).Error
	return int(count), err
}
