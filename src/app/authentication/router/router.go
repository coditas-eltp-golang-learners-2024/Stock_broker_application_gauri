package route

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

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Serve Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize UserRepository
	userRepository := repo.NewUserRepositoryImpl(db)

	// Initialize SignUpService with UserRepository
	userService := service.NewSignUpService(userRepository)

	// SignUp route
	r.POST(constants.SignUpRoute, handlers.SignUpHandler(userService))

	// Initialize SignInService with UserRepository
	userAuthService := service.NewSignInService(userRepository)

	// SignIn route
	r.POST(constants.SignInRoute, handlers.SignInHandler(userAuthService))

	return r
}
