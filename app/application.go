package app

import (
	"net/http"
	"time"

	"github.com/MicaTechnology/escrow_api/utils/logger"
)

type Config struct {
}

func StartApplication() {
	config := Config{}
	StartRepositories()

	srv := &http.Server{
		Handler:      config.routes(),
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Info("Server starting on port 8888...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
