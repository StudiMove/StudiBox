package auth

import (
	request "backend/core/api/request/auth"
	response "backend/core/api/response/auth"
	"backend/core/models"
	"backend/core/services/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandleLogin gère la connexion des utilisateurs
func HandleLogin(c *gin.Context, authService *auth.AuthService) {
	var loginReq request.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "Entrée invalide", "détails": err.Error()})
		return
	}

	token, err := authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erreur": "Email ou mot de passe invalide", "détails": err.Error()})
		return
	}

	resp := response.LoginResponse{
		Token:           token,
		IsAuthenticated: true,
	}
	c.JSON(http.StatusOK, resp)
}

// HandleRegisterUser gère l'inscription des utilisateurs normaux
func HandleRegisterUser(c *gin.Context, authService *auth.AuthService) {
	var registerReq request.RegisterUserRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "Entrée invalide", "détails": err.Error()})
		return
	}

	// Mapping manuel des champs de la requête vers le modèle User
	user := models.User{
		FirstName: registerReq.FirstName,
		LastName:  registerReq.LastName,
		Email:     registerReq.Email,
		Password:  registerReq.Password, // le mot de passe sera haché plus tard
		Phone:     registerReq.Phone,
	}

	if err := authService.RegisterUser(&user); err != nil {
		if err.Error() == "email déjà utilisé" {
			c.JSON(http.StatusConflict, gin.H{"erreur": "Cet email est déjà utilisé", "détails": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de l'inscription", "détails": err.Error()})
		}
		return
	}

	resp := response.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  "user",
	}
	c.JSON(http.StatusCreated, resp)
}

// HandleRegisterBusinessUser gère l'inscription des utilisateurs entreprises
func HandleRegisterBusinessUser(c *gin.Context, authService *auth.AuthService) {
	var registerReq request.RegisterBusinessRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		log.Printf("Erreur de validation du JSON : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"erreur": "Entrée invalide", "détails": err.Error()})
		return
	}

	// Mapping manuel des champs de la requête vers les modèles BusinessUser et User
	businessUser := models.BusinessUser{
		CompanyName: registerReq.CompanyName,
		Address:     registerReq.Address,
		Postcode:    registerReq.Postcode,
		City:        registerReq.City,
		Country:     registerReq.Country,
		User: models.User{
			Email:    registerReq.Email,
			Password: registerReq.Password,
			Phone:    registerReq.Phone,
		},
	}

	if err := authService.RegisterBusinessUser(&businessUser); err != nil {
		log.Printf("Erreur lors de l'enregistrement de l'utilisateur business : %v", err)
		if strings.Contains(err.Error(), "email déjà utilisé") {
			c.JSON(http.StatusConflict, gin.H{"erreur": "Cet email est déjà utilisé", "détails": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"erreur": "Échec de l'inscription", "détails": err.Error()})
		}
		return
	}

	resp := response.RegisterResponse{
		ID:    businessUser.User.ID,
		Email: businessUser.User.Email,
		Role:  "business",
	}
	c.JSON(http.StatusCreated, resp)
}
