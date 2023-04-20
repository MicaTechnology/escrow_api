package app

import (
	"net/http"
	"time"

	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	MapUrls()
	StartRepositories()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	logger.Info("Server starting on port 8888...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
