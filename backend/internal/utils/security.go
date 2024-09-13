package utils

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

// HashPassword génère un hash pour le mot de passe donné.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// VerifyPassword vérifie si le mot de passe donné correspond au hash.
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// JWTClaims représente les claims pour le JWT.
type JWTClaims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// GenerateJWT génère un JWT pour l'utilisateur avec un certain temps d'expiration.
func GenerateJWT(userID uint, secret string) (string, error) {
    claims := &JWTClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 72 heures d'expiration
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

// ValidateJWT valide le token et retourne les claims.
func ValidateJWT(tokenStr string, secret string) (*JWTClaims, error) {
    claims := &JWTClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}
