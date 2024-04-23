package models

import "github.com/dgrijalva/jwt-go"

// Claims represents the JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
