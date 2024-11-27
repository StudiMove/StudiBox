package models

// Modèle de liaison entre les événements et les catégories
type EventCategoryEvent struct {
    EventID         uint `gorm:"primaryKey" json:"eventId"`
    EventCategoryID uint `gorm:"primaryKey" json:"eventCategoryId"`
}
