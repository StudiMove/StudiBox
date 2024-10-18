package response

type CreateEventResponse struct {
	ID          uint     `json:"id"`
	OwnerID     uint     `json:"owner_id"`
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
	Category    string   `json:"category"`
	Tags        string   `json:"tags"`
}

type GetEventResponse struct {
	ID          uint     `json:"id"`
	OwnerID     uint     `json:"owner_id"`
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
	Category    string   `json:"category"`
	Tags        string   `json:"tags"`
}
