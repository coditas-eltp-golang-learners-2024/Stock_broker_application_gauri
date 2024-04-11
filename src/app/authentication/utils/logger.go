package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("[Recovery] Panic recovered: %v\n", err)
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            }
        }()
        c.Next()
    }
}
