// src/app/authentication/main.go

package main

import (
    "Stock_broker_application/utils/db"
    router "Stock_broker_application/router"
    "log"
)
func main() {
      // Initialize database
      db.InitDB()
     //defer db.DB.Close() // Defer closing the database connection until the function returns
    
      // Setup router with the database connection
      r := router.SetupRouter(db.DB)
      
      // Run the server
      if err := r.Run(":8080"); err != nil {
          log.Fatalf("Failed to start server: %v", err)
      }
}