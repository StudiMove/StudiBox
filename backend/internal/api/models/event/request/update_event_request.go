package request

import "backend/internal/db/models"

// UpdateEventRequest inclut les champs de CreateEventRequest et l'ID de l'événement.
type UpdateEventRequest struct {
    EventID uint `json:"event_id"`
    CreateEventRequest
}

// UpdateEventDataRequest représente les informations de mise à jour pour un événement,
// en incluant les champs de CreateEventRequest et les champs spécifiques.
// UpdateEventDataRequest représente les informations spécifiques pour la mise à jour d'un événement existant,
// en incluant les champs de mise à jour uniques ainsi que les champs de CreateEventRequest.
type UpdateEventDataRequest struct {
    EventID       uint                       `json:"event_id"`         // ID de l'événement à mettre à jour
    Tags          []string                   `json:"tags"`             // Tags de l'événement
    Categories    []string                   `json:"categories"`       // Catégories de l'événement
    Descriptions  []models.EventDescription  `json:"descriptions"`     // Descriptions de l'événement
    Options       []models.EventOption       `json:"options"`          // Options spécifiques de l'événement
    Tarifs        []models.EventTarif        `json:"tarifs"`           // Tarifs spécifiques de l'événement
    CreateEventRequest                                                // Inclut tous les champs nécessaires pour créer ou mettre à jour l'événement
}