package app

import (
	"net/http"

	"github.com/MicaTechnology/escrow_api/controllers"
)

func MapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
