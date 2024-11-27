package models

// Modèle de liaison entre les événements et les tags
type EventTagEvent struct {
    EventID    uint `gorm:"primaryKey" json:"eventId"`
    EventTagID uint `gorm:"primaryKey" json:"eventTagId"`
}
