package utils

import (
    "log"
    "github.com/spf13/viper"
    "Stock_broker_application/models"
)

// LoadConfig loads the database configuration from the application.yml file
func LoadConfig() models.SQLConfig {
    // Set the path to the configuration file
    viper.SetConfigFile("resources/application.yml")

    // Read the configuration file into Viper
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    // Unmarshal the configuration into a struct
    var config models.SQLConfig
    if err := viper.UnmarshalKey("database", &config); err != nil {
        log.Fatalf("Unable to decode config into struct: %s", err)
    }

    // Return the configuration
    return config
}
