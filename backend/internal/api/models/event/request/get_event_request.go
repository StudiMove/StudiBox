package request

type GetEventRequest struct {
	TargetEventID uint `json:"event_id"`
}

type GetEventListByTargetIDRequest struct {
    UserTargetID uint `json:"user_id"`
}