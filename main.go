package main

import (
	"github.com/MicaTechnology/escrow_api/app"
	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	app.StartApplication()
}
