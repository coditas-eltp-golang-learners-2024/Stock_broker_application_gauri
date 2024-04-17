package handlers

import (
	"Stock_broker_application/models"
	"Stock_broker_application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Validate OTP
// @Description Validates the OTP for a user
// @Tags OTP
// @Accept json
// @Produce json
// @Param otpRequest body models.OTPRequest true "OTP Request"
// @Success 200 {object} string "OTP validated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "OTP is expired or invalid"
// @Router /validateotp [post]
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
