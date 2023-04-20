package app

import (
	"net/http"

	"github.com/MicaTechnology/escrow_api/controllers"
	"github.com/gorilla/handlers"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

func MapUrls() {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.Use(handlers.CORS(headers, methods, origins))

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/escrow", controllers.EscrowController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/escrow/{id}", controllers.EscrowController.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/escrow/{id}/claim", controllers.EscrowController.Claim).Methods(http.MethodPut)

	// Options
	router.HandleFunc("/api/v1/escrow", optionsHandler).Methods(http.MethodOptions)
	router.HandleFunc("/api/v1/escrow/{id}", optionsHandler).Methods(http.MethodOptions)
	router.HandleFunc("/api/v1/escrow/{id}/claim", optionsHandler).Methods(http.MethodOptions)
}
