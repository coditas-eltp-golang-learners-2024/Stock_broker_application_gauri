package handlers

import (
    "Stock_broker_application/models"
    "Stock_broker_application/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

// SignInHandler handles the sign-in request
// @Summary Handle sign-in request
// @Description Handle sign-in request and authenticate the user
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {object} string "User authenticated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Router /signin [post]
func SignInHandler(userService *service.SignInService, otpService *service.OTPService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var signInRequest models.SignInRequest

        // Bind JSON request body to SignInRequest struct
        if err := c.ShouldBindJSON(&signInRequest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Call SignIn method to authenticate user
        if err := userService.SignIn(signInRequest); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
            return
        }

        // Authentication successful
        c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})

        // Generate and save OTP for the email
        if err := otpService.GenerateAndSaveOTP(signInRequest.Email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate and save OTP"})
            return
        }
    }
}
