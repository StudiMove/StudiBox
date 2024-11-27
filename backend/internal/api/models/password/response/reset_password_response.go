package response

type SendResetCodeResponse struct {
    Success   bool   `json:"success"`
    Message   string `json:"message,empty"`
    ResetCode int `json:"reset_code,empty"` // Optionnel si nécessaire pour des tests ou logs
}

type UpdatePasswordResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message,empty"`
}

type GetResetCodeResponse struct {
    Success   bool   `json:"success"`
    Message   string `json:"message,empty"`
    ResetCode int `json:"reset_code"`
}


type VerifyResetCodeResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message,empty"`
}