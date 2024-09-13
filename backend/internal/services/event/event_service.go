package event

import (
    "backend/internal/db/models"
    "gorm.io/gorm"
)

type EventService struct {
    DB *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
    return &EventService{DB: db}
}

func (s *EventService) CreateEvent(event *models.Event) error {
    if err := s.DB.Create(event).Error; err != nil {
        return err
    }
    return nil
}

func (s *EventService) UpdateEvent(event *models.Event) error {
    if err := s.DB.Save(event).Error; err != nil {
        return err
    }
    return nil
}

func (s *EventService) DeleteEvent(eventID uint) error {
    if err := s.DB.Delete(&models.Event{}, eventID).Error; err != nil {
        return err
    }
    return nil
}

func (s *EventService) GetEvent(eventID uint) (*models.Event, error) {
    var event models.Event
    if err := s.DB.First(&event, eventID).Error; err != nil {
        return nil, err
    }
    return &event, nil
}
