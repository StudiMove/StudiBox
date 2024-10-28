package auth

import (
	"backend/config"
	request "backend/core/api/request/auth"
	responseGlobal "backend/core/api/response"
	response "backend/core/api/response/auth"
	"backend/core/models"
	"backend/core/services/auth"
	"backend/core/services/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleRegisterUser(c *gin.Context, authService *auth.AuthService, userService *user.UserService) {
	var registerReq request.RegisterUserRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, responseGlobal.ErrorResponse("Invalid input", err))
		return
	}

	// Vérifier si un IDToken est fourni pour une connexion Google
	var user models.User
	if registerReq.IDToken != "" {
		// Vérification du token Firebase
		token, err := config.VerifyIDToken(registerReq.IDToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responseGlobal.ErrorResponse("Invalid Firebase ID token", err))
			return
		}

		// Extraire les informations de l'utilisateur depuis le token
		userEmail := token.Claims["email"].(string)
		user = models.User{
			Email:     userEmail,
			FirstName: token.Claims["given_name"].(string),
			LastName:  token.Claims["family_name"].(string),
		}
	} else {
		// Gestion classique de l'inscription si aucun IDToken
		user = models.User{
			FirstName: registerReq.FirstName,
			LastName:  registerReq.LastName,
			Email:     registerReq.Email,
			Password:  registerReq.Password,
			Phone:     registerReq.Phone,
		}
	}

	// Inscription de l'utilisateur
	if err := authService.RegisterUser(&user); err != nil {
		if strings.Contains(err.Error(), "email déjà utilisé") {
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

	// Extraire les rôles de l'utilisateur
	roleNames, err := userService.Management.ExtractRoleNames(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user roles", err))
		return
	}

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
func HandleRegisterBusinessUser(c *gin.Context, authService *auth.AuthService, userService *user.UserService) {
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
		if strings.Contains(err.Error(), "email déjà utilisé") {
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
	roleNames, err := userService.Management.ExtractRoleNames(businessUser.User.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseGlobal.ErrorResponse("Failed to retrieve user roles", err))
		return
	}

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
