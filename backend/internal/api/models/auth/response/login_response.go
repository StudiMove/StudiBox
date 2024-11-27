package response

// LoginResponse représente les données renvoyées après une connexion réussie.
type LoginResponse struct {
    Token           string `json:"token"`
    ProfileImage    string `json:"profile_image"`
    IsAuthenticated bool   `json:"isAuthenticated"`
    Message         string `json:"message,omitempty"`
}
