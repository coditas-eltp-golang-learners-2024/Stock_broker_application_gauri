// handlers/signUpHandler.go
package handlers

import (
    "Stock_broker_application/models"
    "Stock_broker_application/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

// SignUpHandler handles the signup request
func SignUpHandler(userService *service.SignUpService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var signUpRequest models.SignUpRequest

        // Bind JSON request body to SignUpRequest struct
        if err := c.ShouldBindJSON(&signUpRequest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Call SignUp method to process signup request
        if err := userService.SignUp(signUpRequest); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
    }
}
