// src/app/authentication/main.go

package main

import (
    "Stock_broker_application/utils/db"
    "Stock_broker_application/router"
    "log"
)

func main() {
    // Initialize database
    db.InitDB()
    defer db.DB.Close() // Ensure database connection is closed when the main function exits

    // Your application logic here...
    log.Println("Database connection initialized. Starting application...")
    r := route.SetupRouter(db.DB) // Pass the database connection to the router setup function
    r.Run(":8080") // Corrected the address to listen on all interfaces
}
