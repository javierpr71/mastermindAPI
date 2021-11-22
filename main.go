package main

import (
	"fmt"
	"os"

	"github.com/javierpr71/mastermind/controllers"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {

	logger := log.New()

	// Initialize loggging
	logger.SetOutput(os.Stdout)
	logger.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("🚀 MasterMind-API starting...")

	// Load config settings
	var err error
	err = godotenv.Load()
	if err != nil {
		logger.Infof("Error getting env, %v", err)
	} else {
		logger.Info("values from env readed")
	}

	// Create New instance o
	server := controllers.NewServer(logger)
	// Initialize and run Server
	if err := server.Initialize(); err != nil {
		logger.Errorf("Error initializing service: %v", err)
		return
	}
	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
