package handlers

import (
	"Stock_broker_application/models"
	"Stock_broker_application/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidateOTPHandler handles the OTP validation request
func ValidateOTPHandler(otpService *service.OTPService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var otpRequest models.OTPRequest

		// Bind JSON request body to OTPRequest struct
		if err := c.ShouldBindJSON(&otpRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call ValidateOTP method to validate OTP
		if err := otpService.ValidateOTP(otpRequest.Email, otpRequest.OTP); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// OTP validation successful
		c.JSON(http.StatusOK, gin.H{"message": "OTP validated successfully"})
	}
}
