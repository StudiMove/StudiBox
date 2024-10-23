package request

type CreateEventRequest struct {
	OwnerID     uint     `json:"owner_id" binding:"required"`
	OwnerType   string   `json:"owner_type" binding:"required"`
	ImageURLs   []string `json:"image_urls"`
	VideoURL    string   `json:"video_url"`
	Title       string   `json:"title" binding:"required"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description" binding:"required"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	IsOnline    bool     `json:"is_online"`
	IsVisible   bool     `json:"is_visible"`
	Price       int      `json:"price"`
	Address     string   `json:"address"`
	City        string   `json:"city"`
	Postcode    string   `json:"postcode"`
	Region      string   `json:"region"`
	Country     string   `json:"country"`
	Category    []string `json:"category"`
	Tags        []string `json:"tags"`
}

type UpdateEventRequest struct {
	OwnerType   string   `json:"owner_type"`
	ImageURLs   []string `json:"image_urls"`
	VideoURL    string   `json:"video_url"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	IsOnline    bool     `json:"is_online"`
	IsVisible   bool     `json:"is_visible"`
	Price       int      `json:"price"`
	Address     string   `json:"address"`
	City        string   `json:"city"`
	Postcode    string   `json:"postcode"`
	Region      string   `json:"region"`
	Country     string   `json:"country"`
	Category    []string `json:"category"`
	Tags        []string `json:"tags"`
}
