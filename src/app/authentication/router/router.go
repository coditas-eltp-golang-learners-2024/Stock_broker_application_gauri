// route/router.go

package route

import (
    "database/sql" // Import the sql package
    "Stock_broker_application/handlers"
    "Stock_broker_application/repo"
    "Stock_broker_application/service"
    "github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter(db *sql.DB) *gin.Engine {
    // Initialize Gin router
    r := gin.Default()

    // Initialize UserRepository
    userRepository := repo.NewUserRepositoryImpl(db)

    // Initialize SignUpService with UserRepository
    userService := service.NewSignUpService(userRepository)

    // Define routes
    r.POST("/signup", handlers.SignUpHandler(userService))

    return r
}
