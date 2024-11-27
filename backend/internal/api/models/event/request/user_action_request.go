package request

type UpdateUserActionRequest struct {
	UserID           uint `json:"userId" validate:"required"`
	EventID          uint `json:"eventId" validate:"required"`
	IsInterested     bool `json:"isInterested"`
	IsFavorite       bool `json:"isFavorite"`
	UpdateInterested bool `json:"updateInterested"` // Indique si `isInterested` doit être mis à jour
	UpdateFavorite   bool `json:"updateFavorite"`   // Indique si `isFavorite` doit être mis à jour
}

type RemoveUserActionRequest struct {
	UserID  uint `json:"userId"`
	EventID uint `json:"eventId"`
}
