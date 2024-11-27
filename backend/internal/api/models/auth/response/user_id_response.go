package response

type UserIDResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message,omitempty"`
    UserID  uint   `json:"user_id,omitempty"`
}
