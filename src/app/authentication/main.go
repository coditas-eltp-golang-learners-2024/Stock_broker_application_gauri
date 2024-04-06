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
    // Your application logic here...
    log.Println("Database connection initialized. Starting application...")
    r := route.SetupRouter()
    r.Run("localhost:8080")
}
