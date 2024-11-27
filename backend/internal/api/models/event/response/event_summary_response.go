package response

import "time"

type EventSummaryResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    StartDate   time.Time `json:"start_date"`
    StartTime     time.Time `json:"start_time"`
    IsOnline    bool      `json:"is_online"`
    IsActivated bool         `json:"is_activated"`
    ImageURLs string `json:"image_urls"`

    
}
