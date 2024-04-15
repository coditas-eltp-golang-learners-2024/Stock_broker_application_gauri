package main

import (
    "Stock_broker_application/router"
    "Stock_broker_application/utils/db"
    "log"
)

// @title Stock Broker Application API
// @version 1.0
// @description API endpoints for a stock broker application
// @host localhost:8080
// @BasePath /
func main() {
    // Initialize database
    db.InitDB()

    // Setup router with the database connection
    r := router.SetupRouter(db.DB)

    // Run the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}