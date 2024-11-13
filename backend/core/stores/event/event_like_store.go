package event

import (
	"backend/database/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type EventLikeStoreType struct {
	db *gorm.DB
}

func EventLikeStore(db *gorm.DB) *EventLikeStoreType {
	return &EventLikeStoreType{db: db}
}

// IsLikeExists vérifie si un utilisateur a déjà liké un événement
func (s *EventLikeStoreType) IsLikeExists(userID, eventID uint) (bool, error) {
	var existingLike models.EventLike
	err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&existingLike).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // Pas de like existant
	} else if err != nil {
		return false, err // Erreur inattendue
	}
	return true, nil // Le like existe déjà
}

// AddLike ajoute un like pour un événement sans vérification préalable
func (s *EventLikeStoreType) AddLike(userID, eventID uint) error {
	like := models.EventLike{UserID: userID, EventID: eventID}
	return s.db.Create(&like).Error
}

// UnlikeEvent retire un like d'un événement
func (s *EventLikeStoreType) UnlikeEvent(userID, eventID uint) error {
	return s.db.Where("user_id = ? AND event_id = ?", userID, eventID).Delete(&models.EventLike{}).Error
}

// CountLikes compte les likes d'un événement
func (s *EventLikeStoreType) CountLikes(eventID uint) (int, error) {
	var count int64
	err := s.db.Model(&models.EventLike{}).Where("event_id = ?", eventID).Count(&count).Error
	return int(count), err
}

// GetLikedEventsByUser récupère les événements likés par un utilisateur avec leurs catégories et tags
func (s *EventLikeStoreType) GetLikedEventsByUser(userID uint) ([]models.Event, error) {
	var likedEvents []models.Event
	err := s.db.Joins("JOIN event_likes ON events.id = event_likes.event_id").
		Where("event_likes.user_id = ?", userID).
		Preload("Categories").
		Preload("Tags").
		Preload("Descriptions").
		Preload("Options").
		Preload("Tarifs").
		Find(&likedEvents).Error
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des événements likés par l'utilisateur : %w", err)
	}
	return likedEvents, nil
}
