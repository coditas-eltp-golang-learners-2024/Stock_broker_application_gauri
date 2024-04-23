package handlers

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/models"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware function to validate JWT token
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			//c.Abort()
			return
		}
		const bearerPrefix = "Bearer "
        if strings.HasPrefix(tokenString, bearerPrefix) {
            // Remove the "Bearer " prefix
            tokenString = tokenString[len(bearerPrefix):]
        }

        // Trim any leading or trailing whitespace from the token string
        tokenString = strings.TrimSpace(tokenString)
		// Parse JWT token
		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return constants.JwtKey, nil
		})

		if err != nil {
			log.Printf("Error parsing JWT token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			//c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			c.Set("Email", claims.Email)
			c.Next()
		} else {
			log.Println("Invalid token or claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
