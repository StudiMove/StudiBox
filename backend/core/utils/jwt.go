package utils

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims contient les informations stockées dans le JWT, y compris l'ID utilisateur.
type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// HashPassword génère un hash sécurisé pour un mot de passe donné.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword compare le mot de passe en clair avec le hash stocké.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateJWT génère un token JWT avec une durée d'expiration personnalisée.
func GenerateJWT(userID uint, secret string, expirationHours int) (string, error) {
	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT valide un JWT et retourne les claims contenues dans le token.

func ValidateJWT(tokenStr, secret string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GetClaimsFromContext récupère les claims JWT à partir du contexte Gin.
func GetClaimsFromContext(c *gin.Context) (*JWTClaims, error) {
	claimsValue, exists := c.Get("user")
	if !exists {
		return nil, errors.New("token absent ou invalide")
	}
	claims, ok := claimsValue.(*JWTClaims)
	if !ok {
		return nil, errors.New("type de claims invalide")
	}
	return claims, nil
}

// ExtractUserIDFromToken extrait l'ID utilisateur depuis un token JWT valide.
func ExtractUserIDFromToken(tokenStr, secret string) (uint, error) {
	claims, err := ValidateJWT(tokenStr, secret)
	if err != nil {
		return 0, errors.New("unable to extract user ID: invalid token")
	}
	return claims.UserID, nil
}
