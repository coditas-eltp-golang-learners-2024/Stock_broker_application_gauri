package handlers

import (
	"Stock_broker_application/models"
	"Stock_broker_application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChangePasswordHandler handles HTTP requests for changing passwords
func ChangePasswordHandler(passwordService *service.PasswordService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput models.ChangePassword

		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if err := passwordService.ChangePasswordService(&userInput); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
	}
}
