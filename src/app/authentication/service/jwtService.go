package service

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(email string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the token string
	tokenString, err := token.SignedString(constants.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
