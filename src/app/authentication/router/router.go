// route/router.go

package route

import (
	"Stock_broker_application/constants"
	"Stock_broker_application/handlers"
	"Stock_broker_application/repo"
	"Stock_broker_application/service"
	"database/sql" 

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
    r.POST(constants.SignUpRoute, handlers.SignUpHandler(userService))

    return r
}
