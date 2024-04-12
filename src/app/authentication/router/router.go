package router

import (
	"Stock_broker_application/constants"
	_ "Stock_broker_application/docs" // Import generated Swagger docs
	"Stock_broker_application/handlers"
	"Stock_broker_application/repo"
	"Stock_broker_application/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// SetupRouter sets up the routes for the application
// @title Stock Broker Application API
// @description API endpoints for a stock broker application
// @version 1.0
// @host localhost:8080
// @BasePath /api
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Initialize UserRepository
	userRepository := repo.NewUserRepositoryImpl(db)

	// Initialize SignUpService with UserRepository
	userService := service.NewSignUpService(userRepository)
	// SignUp route
	// @Summary Register a new user
	// @Description Register a new user with the provided details
	// @Tags authentication
	// @Accept json
	// @Produce json
	// @Param body body models.SignUpRequest true "User registration details"
	// @Success 200 {string} string "User signed up successfully"
	// @Failure 400 {object} ErrorResponse "Bad request"
	// @Failure 500 {object} ErrorResponse "Internal server error"
	// @Router /signup [post]

	r.POST(constants.SignUpRoute, handlers.SignUpHandler(userService))

	// Initialize SignInService with UserRepository
	userAuthService := service.NewSignInService(userRepository)

	// SignIn route
	// @Summary Authenticate user
	// @Description Authenticate a user with the provided credentials
	// @Tags authentication
	// @Accept json
	// @Produce json
	// @Param body body models.SignInRequest true "User credentials"
	// @Success 200 {string} string "User authenticated successfully"
	// @Failure 400 {object} ErrorResponse "Bad request"
	// @Failure 401 {object} ErrorResponse "Unauthorized"
	// @Failure 500 {object} ErrorResponse "Internal server error"
	// @Router /signin [post]
	r.POST(constants.SignInRoute, handlers.SignInHandler(userAuthService))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
