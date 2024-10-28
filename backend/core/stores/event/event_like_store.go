package stores

import (
	"backend/core/models"

	"gorm.io/gorm"
)

type EventLikeStore struct {
	db *gorm.DB
}

func NewEventLikeStore(db *gorm.DB) *EventLikeStore {
	return &EventLikeStore{db: db}
}

// LikeEvent ajoute un like à un événement
func (s *EventLikeStore) LikeEvent(userID, eventID uint) error {
	var existingLike models.EventLike
	if err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&existingLike).Error; err == nil {
		return nil // Le like existe déjà
	} else if err != gorm.ErrRecordNotFound {
		return err // Si une autre erreur survient, la retourner
	}

	// Ajouter un nouveau like
	like := models.EventLike{UserID: userID, EventID: eventID}
	return s.db.Create(&like).Error
}

// UnlikeEvent retire un like d'un événement
func (s *EventLikeStore) UnlikeEvent(userID, eventID uint) error {
	return s.db.Where("user_id = ? AND event_id = ?", userID, eventID).Delete(&models.EventLike{}).Error
}

// GetLikedEventsByUser récupère les événements likés par un utilisateur
func (s *EventLikeStore) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.db.Joins("JOIN event_likes ON events.id = event_likes.event_id").
		Where("event_likes.user_id = ?", userID).Find(&likedEvents).Error
	return likedEvents, err
}

// GetLikedEventsByUserWithRelations récupère les événements likés avec relations
func (s *EventLikeStore) GetLikedEventsByUserWithRelations(userID uint, likedEvents *[]models.Event) error {
	return s.db.Joins("JOIN event_likes ON events.id = event_likes.event_id").
		Where("event_likes.user_id = ?", userID).
		Preload("Categories").Preload("Tags").
		Find(likedEvents).Error
}

// CountLikes retourne le nombre de likes pour un événement donné
func (s *EventLikeStore) CountLikes(eventID uint) (int, error) {
	var count int64
	err := s.db.Model(&models.EventLike{}).Where("event_id = ?", eventID).Count(&count).Error
	return int(count), err
}
