package event

import (
	"backend/internal/db/models"
	"log"

	"gorm.io/gorm"
)

type EventUserActionService struct {
	db *gorm.DB
}

func NewEventUserActionService(db *gorm.DB) *EventUserActionService {
	return &EventUserActionService{db: db}
}

// Mettre à jour une action utilisateur sur un événement
func (s *EventUserActionService) UpdateUserAction(userID uint, eventID uint, isInterested bool, isFavorite bool, updateInterested bool, updateFavorite bool) error {
	var action models.EventUserAction

	// Rechercher une action existante
	err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&action).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Si aucune action n'existe, en créer une nouvelle
			newAction := models.EventUserAction{
				UserID:       userID,
				EventID:      eventID,
				IsInterested: false, // Par défaut
				IsFavorite:   false, // Par défaut
			}

			// Mettre à jour uniquement si indiqué
			if updateInterested {
				newAction.IsInterested = isInterested
			}
			if updateFavorite {
				newAction.IsFavorite = isFavorite
			}

			if err := s.db.Create(&newAction).Error; err != nil {
				log.Printf("Erreur lors de la création d'une action utilisateur : %v", err)
				return err
			}
			return nil
		}
		log.Printf("Erreur lors de la recherche d'une action utilisateur existante : %v", err)
		return err
	}

	// Si une action existe déjà, mettre à jour uniquement les champs nécessaires
	if updateInterested {
		action.IsInterested = isInterested
	}
	if updateFavorite {
		action.IsFavorite = isFavorite
	}

	if err := s.db.Save(&action).Error; err != nil {
		log.Printf("Erreur lors de la mise à jour de l'action utilisateur : %v", err)
		return err
	}

	return nil
}

// Récupérer les événements marqués comme favoris par un utilisateur
func (s *EventUserActionService) GetFavoriteEventsByUser(userID uint) ([]models.Event, error) {
	var events []models.Event
	err := s.db.Joins("JOIN event_user_actions ON events.id = event_user_actions.event_id").
		Where("event_user_actions.user_id = ? AND event_user_actions.is_favorite = ?", userID, true).
		Find(&events).Error

	if err != nil {
		log.Printf("Erreur lors de la récupération des événements favoris pour l'utilisateur %d : %v", userID, err)
		return nil, err
	}

	return events, nil
}

// Récupérer les événements où l'utilisateur est intéressé
func (s *EventUserActionService) GetInterestedEventsByUser(userID uint) ([]models.Event, error) {
	var events []models.Event
	err := s.db.Joins("JOIN event_user_actions ON events.id = event_user_actions.event_id").
		Where("event_user_actions.user_id = ? AND event_user_actions.is_interested = ?", userID, true).
		Find(&events).Error

	if err != nil {
		log.Printf("Erreur lors de la récupération des événements intéressés pour l'utilisateur %d : %v", userID, err)
		return nil, err
	}

	return events, nil
}

// Supprimer une action spécifique
func (s *EventUserActionService) RemoveUserAction(userID uint, eventID uint) error {
	if err := s.db.Where("user_id = ? AND event_id = ?", userID, eventID).Delete(&models.EventUserAction{}).Error; err != nil {
		log.Printf("Erreur lors de la suppression d'une action utilisateur : %v", err)
		return err
	}
	return nil
}
