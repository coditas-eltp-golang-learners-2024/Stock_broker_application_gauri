package handlers

// SignUpHandler handles the signup request
// @Summary Handle signup request
// @Description Handle signup request and create a new user
// @Accept json
// @Produce json
// @Param request body models.SignUpRequest true "Sign up request body"
// @Success 200 {object} gin.H{"message": "User signed up successfully"}
// @Failure 400 {object} gin.H{"error": "Bad request"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}

// @Router /signup [post]
import (
	"Stock_broker_application/models"
	"Stock_broker_application/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
