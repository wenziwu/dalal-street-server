package utils

import (
	"encoding/json"
	"os"
	"log"
)

// Configuration contains all the configuration options
var Configuration = struct {
	// Pragyan API related options

	// EventId is the Id of DalalStreet event
	EventId string
	// EventSecret is the secret string for DalalStreet event for Pragyan API
	EventSecret string

	// Logging related options

	// LogFileName is the name of the log file name
	LogFileName string
	// LogMaxSize is the maximum size(MB) of a log file before it gets rotated
	LogMaxSize int
	// LogLevel determines the log level.
	// Can be one of "debug", "info", "warn", "error"
	LogLevel string

	// Database related options

	// DbUser is the name of the database user
	DbUser string
	// DbPassword is the password of the database user
	DbPassword string
	// DbHost is the host name of the database server
	DbHost string
	// DbName is the name of the database
	DbName string
}{}

// InitConfiguration reads the config.json file and loads the
// config options into Configuration
func InitConfiguration() {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Failed to open config.json. Cannot proceed")
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Configuration)

	if err != nil {
		log.Fatal("Failed to load configuration. Cannot proceed. Error: ", err)
	}

	log.Printf("Loaded configuration from config.json: %+v\n", Configuration)
}
