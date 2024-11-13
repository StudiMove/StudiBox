package event

import (
	request "backend/core/api/request/event"
	stores "backend/core/stores/event"
	"backend/database/models"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type EventManagementServiceType struct {
	eventStore       *stores.EventStoreType
	tagStore         *stores.EventTagStoreType
	categoryStore    *stores.EventCategoryStoreType
	tarifStore       *stores.EventTarifStoreType
	optionStore      *stores.EventOptionStoreType
	descriptionStore *stores.EventDescriptionStoreType
}

func EventManagementService(
	eventStore *stores.EventStoreType,
	tagStore *stores.EventTagStoreType,
	categoryStore *stores.EventCategoryStoreType,
	tarifStore *stores.EventTarifStoreType,
	optionStore *stores.EventOptionStoreType,
	descriptionStore *stores.EventDescriptionStoreType,
) *EventManagementServiceType {
	return &EventManagementServiceType{
		eventStore:       eventStore,
		tagStore:         tagStore,
		categoryStore:    categoryStore,
		tarifStore:       tarifStore,
		optionStore:      optionStore,
		descriptionStore: descriptionStore,
	}
}

func (s *EventManagementServiceType) DeleteEvent(eventID uint, userID uint, role string) error {
	// Récupérer l'événement par son ID
	event, err := s.eventStore.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound // Renvoie une erreur spécifique
		}
		return fmt.Errorf("erreur lors de la récupération de l'événement : %w", err)
	}

	// Vérifier si l'utilisateur est Admin ou est le propriétaire de l'événement
	if role != "Admin" && event.OwnerID != userID {
		return errors.New("non autorisé à supprimer cet événement")
	}

	// Démarrer une transaction
	tx := s.eventStore.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Supprimer les descriptions associées
	if err := s.descriptionStore.DeleteByEventIDWithTx(eventID, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("erreur lors de la suppression des descriptions : %w", err)
	}

	// Supprimer les options associées
	if err := s.optionStore.DeleteByEventIDWithTx(eventID, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("erreur lors de la suppression des options : %w", err)
	}

	// Supprimer les tarifs associés
	if err := s.tarifStore.DeleteByEventIDWithTx(eventID, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("erreur lors de la suppression des tarifs : %w", err)
	}

	// Supprimer l'événement lui-même
	if err := s.eventStore.DeleteWithTx(eventID, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("erreur lors de la suppression de l'événement : %w", err)
	}

	// Commit de la transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("erreur lors du commit de la transaction : %w", err)
	}

	return nil
}

func (s *EventManagementServiceType) CreateEvent(event *models.Event, input request.CreateEventRequest) error {
	event.VideoURL = input.VideoURL
	event.Title = input.Title
	event.Subtitle = input.Subtitle
	event.StartDate = input.StartDate
	event.EndDate = input.EndDate
	event.StartTime = input.StartTime
	event.EndTime = input.EndTime
	event.IsOnline = input.IsOnline
	event.IsPublic = input.IsPublic
	event.Address = input.Address
	event.City = input.City
	event.PostalCode = input.PostalCode
	event.Region = input.Region
	event.Country = input.Country
	event.UseStudibox = input.UseStudibox
	event.CategoryIDs = pq.Int64Array(input.CategoryIDs)
	event.TagIDs = pq.Int64Array(input.TagIDs)

	if err := s.eventStore.Create(event); err != nil {
		return errors.New("erreur lors de la création de l'événement : " + err.Error())
	}

	if err := s.manageOptionsForEvent(event.ID, input.Options, false); err != nil {
		return err
	}
	if err := s.manageTarifsForEvent(event.ID, input.Tarifs, false); err != nil {
		return err
	}
	if err := s.manageDescriptionsForEvent(event.ID, input.Descriptions, false); err != nil {
		return err
	}

	return nil
}

func (s *EventManagementServiceType) UpdateEvent(event *models.Event, input request.UpdateEventRequest) error {
	// Met à jour les champs de base de l'événement
	if err := s.updateEventFields(event, input); err != nil {
		return err
	}

	// Mise à jour de l'image si une URL est fournie
	if input.ImageURL != "" {
		event.ImageURL = input.ImageURL
	}

	if err := s.manageOptionsForEvent(event.ID, input.Options, true); err != nil {
		return err
	}
	if err := s.manageTarifsForEvent(event.ID, input.Tarifs, true); err != nil {
		return err
	}
	if err := s.manageDescriptionsForEvent(event.ID, input.Descriptions, true); err != nil {
		return err
	}

	// Enregistre les modifications
	return s.eventStore.Update(event)
}

func (s *EventManagementServiceType) updateEventFields(event *models.Event, input request.UpdateEventRequest) error {
	event.Title = input.Title
	event.Subtitle = input.Subtitle
	event.StartDate = input.StartDate
	event.EndDate = input.EndDate
	event.StartTime = input.StartTime
	event.EndTime = input.EndTime
	event.Address = input.Address
	event.City = input.City
	event.PostalCode = input.PostalCode
	event.Region = input.Region
	event.Country = input.Country
	event.VideoURL = input.VideoURL
	event.CategoryIDs = pq.Int64Array(input.CategoryIDs)
	event.TagIDs = pq.Int64Array(input.TagIDs)

	return s.eventStore.Update(event)
}

func (s *EventManagementServiceType) manageTarifsForEvent(eventID uint, tarifs []request.EventTarifRequest, isUpdate bool) error {
	if isUpdate {
		currentTarifs, err := s.tarifStore.GetByEventID(eventID)
		if err != nil {
			return errors.New("erreur lors de la récupération des tarifs actuels")
		}
		for _, newTarif := range tarifs {
			exists := false
			for _, currentTarif := range currentTarifs {
				if currentTarif.ID == newTarif.ID {
					currentTarif.Title = newTarif.Title
					currentTarif.Description = newTarif.Description
					currentTarif.Price = newTarif.Price
					currentTarif.Stock = newTarif.Stock
					s.tarifStore.Update(&currentTarif)
					exists = true
					break
				}
			}
			if !exists {
				eventTarif := models.EventTarif{
					EventID:     eventID,
					Title:       newTarif.Title,
					Description: newTarif.Description,
					Price:       newTarif.Price,
					Stock:       newTarif.Stock,
				}
				s.tarifStore.Create(&eventTarif)
			}
		}
	} else {
		for _, tarif := range tarifs {
			eventTarif := models.EventTarif{
				EventID:     eventID,
				Title:       tarif.Title,
				Description: tarif.Description,
				Price:       tarif.Price,
				Stock:       tarif.Stock,
			}
			if err := s.tarifStore.Create(&eventTarif); err != nil {
				return errors.New("erreur lors de l'ajout des tarifs à l'événement : " + err.Error())
			}
		}
	}
	return nil
}

func (s *EventManagementServiceType) manageOptionsForEvent(eventID uint, options []request.EventOptionRequest, isUpdate bool) error {
	if isUpdate {
		currentOptions, err := s.optionStore.GetByEventID(eventID)
		if err != nil {
			return errors.New("erreur lors de la récupération des options actuelles")
		}
		for _, newOption := range options {
			exists := false
			for _, currentOption := range currentOptions {
				if currentOption.ID == newOption.ID {
					currentOption.Title = newOption.Title
					currentOption.Description = newOption.Description
					currentOption.Price = newOption.Price
					currentOption.Stock = newOption.Stock
					s.optionStore.Update(&currentOption)
					exists = true
					break
				}
			}
			if !exists {
				option := models.EventOption{
					EventID:     eventID,
					Title:       newOption.Title,
					Description: newOption.Description,
					Price:       newOption.Price,
					Stock:       newOption.Stock,
				}
				s.optionStore.Create(&option)
			}
		}
	} else {
		for _, option := range options {
			eventOption := models.EventOption{
				EventID:     eventID,
				Title:       option.Title,
				Description: option.Description,
				Price:       option.Price,
				Stock:       option.Stock,
			}
			if err := s.optionStore.Create(&eventOption); err != nil {
				return errors.New("erreur lors de l'ajout des options à l'événement : " + err.Error())
			}
		}
	}
	return nil
}

func (s *EventManagementServiceType) manageDescriptionsForEvent(eventID uint, descriptions []request.EventDescriptionRequest, isUpdate bool) error {
	if isUpdate {
		currentDescriptions, err := s.descriptionStore.GetByEventID(eventID)
		if err != nil {
			return errors.New("erreur lors de la récupération des descriptions actuelles")
		}
		descriptionMap := make(map[uint]*models.EventDescription)
		for i := range currentDescriptions {
			descriptionMap[currentDescriptions[i].ID] = &currentDescriptions[i]
		}
		for _, newDesc := range descriptions {
			if existingDesc, exists := descriptionMap[newDesc.ID]; exists {
				existingDesc.Title = newDesc.Title
				existingDesc.Description = newDesc.Description
				if err := s.descriptionStore.Update(existingDesc); err != nil {
					return errors.New("erreur lors de la mise à jour des descriptions de l'événement : " + err.Error())
				}
				delete(descriptionMap, newDesc.ID)
			} else {
				eventDescription := models.EventDescription{
					EventID:     eventID,
					Title:       newDesc.Title,
					Description: newDesc.Description,
				}
				s.descriptionStore.Create(&eventDescription)
			}
		}
		for _, descToDelete := range descriptionMap {
			s.descriptionStore.Delete(descToDelete.ID)
		}
	} else {
		for _, desc := range descriptions {
			eventDescription := models.EventDescription{
				EventID:     eventID,
				Title:       desc.Title,
				Description: desc.Description,
			}
			s.descriptionStore.Create(&eventDescription)
		}
	}
	return nil
}
