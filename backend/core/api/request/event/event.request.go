package request

import (
	"time"
)

// package request

type CreateEventRequest struct {
	ImageURL     string                    `json:"image_urls"`
	VideoURL     string                    `json:"video_url"`
	Title        string                    `json:"title" binding:"required,max=100"`
	Subtitle     string                    `json:"subtitle" binding:"max=255"`
	StartDate    time.Time                 `json:"start_date" binding:"required"`
	EndDate      time.Time                 `json:"end_date"`
	StartTime    time.Time                 `json:"start_time"`
	EndTime      time.Time                 `json:"end_time"`
	IsOnline     bool                      `json:"is_online"`
	IsPublic     bool                      `json:"is_visible"`
	Address      string                    `json:"address" binding:"required_if=IsOnline false"`
	City         string                    `json:"city" binding:"required_if=IsOnline false"`
	PostalCode   int32                     `json:"postalcode" binding:"required_if=IsOnline false"`
	Region       string                    `json:"region"`
	Country      string                    `json:"country"`
	CategoryIDs  []int64                   `json:"category_ids"`
	TagIDs       []int64                   `json:"tag_ids"`
	UseStudibox  bool                      `json:"use_studibox"`
	Options      []EventOptionRequest      `json:"options"`
	Tarifs       []EventTarifRequest       `json:"tarifs"`
	Descriptions []EventDescriptionRequest `json:"descriptions"`
}

type UpdateEventRequest struct {
	ImageURL     string                    `json:"image_urls"`
	VideoURL     string                    `json:"video_url"`
	Title        string                    `json:"title" binding:"omitempty,max=100"`
	Subtitle     string                    `json:"subtitle" binding:"omitempty,max=255"`
	StartDate    time.Time                 `json:"start_date"`
	EndDate      time.Time                 `json:"end_date"`
	StartTime    time.Time                 `json:"start_time"`
	EndTime      time.Time                 `json:"end_time"`
	IsOnline     *bool                     `json:"is_online"`
	IsPublic     *bool                     `json:"is_visible"`
	Address      string                    `json:"address" binding:"omitempty,max=100"`
	City         string                    `json:"city" binding:"omitempty,max=100"`
	PostalCode   int32                     `json:"postalcode" binding:"omitempty"`
	Region       string                    `json:"region"`
	Country      string                    `json:"country"`
	CategoryIDs  []int64                   `json:"category_ids"`
	TagIDs       []int64                   `json:"tag_ids"`
	UseStudibox  *bool                     `json:"use_studibox"`
	Options      []EventOptionRequest      `json:"options"`
	Tarifs       []EventTarifRequest       `json:"tarifs"`
	Descriptions []EventDescriptionRequest `json:"descriptions"`
}

// EventDescriptionRequest définit la structure pour chaque description d'un événement
type EventDescriptionRequest struct {
	ID          uint   `json:"id,omitempty"`
	Title       string `json:"title" binding:"required,max=100"`
	Description string `json:"description" binding:"omitempty"`
}

// Définition de la structure pour les options d'un événement
type EventOptionRequest struct {
	ID          uint    `json:"id,omitempty"`
	Title       string  `json:"title" binding:"required,max=100"`
	Description string  `json:"description" binding:"omitempty"`
	Price       float64 `json:"price" binding:"gte=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}

// Définition de la structure pour les tarifs d'un événement
type EventTarifRequest struct {
	ID          uint    `json:"id,omitempty"`
	Title       string  `json:"title" binding:"required,max=100"`
	Description string  `json:"description" binding:"omitempty"`
	Price       float64 `json:"price" binding:"gte=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}
