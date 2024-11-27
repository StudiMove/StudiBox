package request

type VerifyResetCodeRequest struct {
    Email     string `json:"email"`
    ResetCode int    `json:"reset_code"`
}

type SendResetCodeRequest struct {
    Email string `json:"email"`
}


type GetResetCodeRequest struct {
    Email string `json:"email"`
}
