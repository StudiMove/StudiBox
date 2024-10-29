package utils

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims contient les informations stockées dans le JWT, y compris l'ID utilisateur.
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Issuer   string `json:"iss"` // Emis par
	Audience string `json:"aud"` // Destiné à
	jwt.RegisteredClaims
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

// GenerateJWT génère un token JWT avec une durée d'expiration personnalisée et des claims supplémentaires.
func GenerateJWT(userID uint, secret, issuer, audience string, expirationHours int) (string, error) {
	claims := &JWTClaims{
		UserID:   userID,
		Issuer:   issuer,
		Audience: audience,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user_authentication",
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

	// Ajout d'une vérification supplémentaire sur l'audience et l'issuer
	if claims.Issuer != "StudiMove" || claims.Audience != "studi_users" {
		return nil, errors.New("invalid token claims: issuer or audience")
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
