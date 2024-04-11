package handlers
// SignInHandler handles the signin request
// @Summary Handle signin request
// @Description Handle signin request and authenticate the user
// @Accept json
// @Produce json
// @Param request body SignInRequest true "Sign in request body"
// @Success 200 {object} gin.H{"message": "User signed in successfully"}
// @Failure 400 {object} gin.H{"error": "Bad request"}
// @Failure 401 {object} gin.H{"error": "Unauthorized"}
// @Router /signin [post]
import (
    "Stock_broker_application/models"
    "Stock_broker_application/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

// SignInHandler handles the sign-in request
func SignInHandler(userService *service.SignInService) gin.HandlerFunc {
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
    }
}
