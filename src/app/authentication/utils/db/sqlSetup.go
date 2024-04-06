// src/app/authentication/utils/db/sqlSetup.go

package db

import (
    "database/sql"
    "fmt"
    "log"
    "Stock_broker_application/utils"
   "Stock_broker_application/constants"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

var DB *sql.DB

func InitDB() {
    // Load configuration
    sqlConfig := utils.LoadConfig()

    // Database connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        sqlConfig.Username,
        sqlConfig.Password,
        sqlConfig.Host,
        sqlConfig.Port,
        sqlConfig.DBName,
    )

    var err error
    DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatalf("Error connecting to database: %v",constants.ErrDatabaseConnection)
    }

    // Test the connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v",constants.ErrDatabasePing)
    }
    log.Println("Connected to database")
}
