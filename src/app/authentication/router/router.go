package router

import (
	"Stock_broker_application/constants"
	_ "Stock_broker_application/docs"
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
// @BasePath /
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Initialize UserRepository
	userRepository := repo.NewUserRepositoryImpl(db)

	// Initialize SignUpService with UserRepository
	userService := service.NewSignUpService(userRepository)

	// Initialize SignInService with UserRepository
	userAuthService := service.NewSignInService(userRepository)

	r.POST(constants.SignUpRoute, handlers.SignUpHandler(userService))
	r.POST(constants.SignInRoute, handlers.SignInHandler(userAuthService))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
