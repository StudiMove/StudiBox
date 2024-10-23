package auth

import (
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/models"
	"backend/core/services/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HandleRegisterUser gère l'inscription des utilisateurs normaux, suivi d'une connexion automatique
func HandleRegisterUser(c *gin.Context, authService *auth.AuthService) {
	var registerReq request.RegisterUserRequest

	// Validation des données d'entrée
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	// Mapping manuel des champs de la requête vers le modèle User
	user := models.User{
		FirstName: registerReq.FirstName,
		LastName:  registerReq.LastName,
		Email:     registerReq.Email,
		Password:  registerReq.Password, // Le mot de passe sera haché plus tard
		Phone:     registerReq.Phone,
	}

	// Inscription de l'utilisateur
	if err := authService.RegisterUser(&user); err != nil {
		if strings.Contains(err.Error(), "email already used") {
			c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("Email already used", err))
		} else {
			c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Registration failed", err))
		}
		return
	}

	// Connexion automatique après l'inscription
	token, err := authService.Login(user.Email, registerReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to login after registration", err))
		return
	}

	// Extraire les noms des rôles via le service
	roleNames := authService.ExtractRoleNames(user.Roles)

	// Créer la réponse
	resp := response.LoginResponse{
		Token:           token,
		IsAuthenticated: true,
		User: response.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Roles:     roleNames,
		},
	}
	c.JSON(http.StatusCreated, responseGlobal.SuccessResponse("User registered and logged in successfully", resp))
}

// HandleRegisterBusinessUser gère l'inscription des utilisateurs entreprises, suivi d'une connexion automatique
func HandleRegisterBusinessUser(c *gin.Context, authService *auth.AuthService) {
	var registerReq request.RegisterBusinessRequest

	// Validation des données d'entrée
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		log.Printf("Erreur de validation du JSON : %v", err)
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
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

	// Inscription de l'utilisateur business
	if err := authService.RegisterBusinessUser(&businessUser); err != nil {
		log.Printf("Erreur lors de l'enregistrement de l'utilisateur business : %v", err)
		if strings.Contains(err.Error(), "email already used") {
			c.JSON(http.StatusConflict, responseGlobal.ErrorResponse("Email already used", err))
		} else {
			c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Registration failed", err))
		}
		return
	}

	// Connexion automatique après l'inscription
	token, err := authService.Login(businessUser.User.Email, registerReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to login after registration", err))
		return
	}

	// Extraire les noms des rôles via le service
	roleNames := authService.ExtractRoleNames(businessUser.User.Roles)

	// Créer la réponse
	resp := response.LoginResponse{
		Token:           token,
		IsAuthenticated: true,
		User: response.UserResponse{
			ID:        businessUser.ID,
			Email:     businessUser.User.Email,
			FirstName: businessUser.User.FirstName,
			LastName:  businessUser.User.LastName,
			Roles:     roleNames,
		},
	}
	c.JSON(http.StatusCreated, responseGlobal.SuccessResponse("Business user registered and logged in successfully", resp))
}
