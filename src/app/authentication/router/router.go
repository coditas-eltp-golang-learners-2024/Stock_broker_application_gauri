package router

import (
    "Stock_broker_application/constants"
    "Stock_broker_application/docs"
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

    // Initialize UserRepository
    userRepository := repo.NewUserRepositoryImpl(db)

    // Initialize SignUpService with UserRepository
    userService := service.NewSignUpService(userRepository)

    // Initialize SignInService with UserRepository
    userAuthService := service.NewSignInService(userRepository)

    // Initialize OTPService
    otpService := service.NewOTPService(userRepository)

    r.POST(constants.SignUpRoute, handlers.SignUpHandler(userService))
    r.POST(constants.SignInRoute, handlers.SignInHandler(userAuthService, otpService))
    r.POST(constants.ValidateOtpRoute, handlers.ValidateOTPHandler(otpService))
    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Swagger documentation
    docs.SwaggerInfo.Title = "Stock Broker Application API"
    docs.SwaggerInfo.Description = "API endpoints for a stock broker application"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
