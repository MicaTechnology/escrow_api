package app

import (
	"fmt"
	"net/http"
	"os"
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
		Addr:         fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.GetLogger().Printf("Server starting on port %s", os.Getenv("PORT"))
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
