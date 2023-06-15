package app

import (
	"net/http"

	"github.com/MicaTechnology/escrow_api/controllers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "*")
	w.Header().Set("Allow-Control-Allow-Headers", "*")
	w.Header().Set("Allow-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

func (app *Config) routes() http.Handler {
	mux := mux.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           3000,
	}))

	mux.Use(middleware.Heartbeat("/health"))

	mux.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	mux.HandleFunc("/friend-bot/add-funds", controllers.FriendBotController.AddFunds).Methods(http.MethodPut)
	mux.HandleFunc("/api/v1/escrow", controllers.EscrowController.Create).Methods(http.MethodPost)
	mux.HandleFunc("/api/v1/escrow/{id}", controllers.EscrowController.Get).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/escrow/{id}/claim", controllers.EscrowController.Claim).Methods(http.MethodPut)

	// // Options
	mux.HandleFunc("/api/v1/escrow", optionsHandler).Methods(http.MethodOptions)
	mux.HandleFunc("/api/v1/escrow/{id}", optionsHandler).Methods(http.MethodOptions)
	mux.HandleFunc("/api/v1/escrow/{id}/claim", optionsHandler).Methods(http.MethodOptions)

	return mux
}
