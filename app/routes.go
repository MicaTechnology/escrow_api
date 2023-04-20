package app

import (
	"net/http"

	"github.com/MicaTechnology/escrow_api/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

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

	mux.Get("/ping", controllers.PingController.Ping)
	mux.Post("/api/v1/escrow", controllers.EscrowController.Create)
	mux.Get("/api/v1/escrow/{id}", controllers.EscrowController.Get)
	mux.Put("/api/v1/escrow/{id}/claim", controllers.EscrowController.Claim)

	// // Options
	mux.Options("/api/v1/escrow", optionsHandler)
	mux.Options("/api/v1/escrow/{id}", optionsHandler)
	mux.Options("/api/v1/escrow/{id}/claim", optionsHandler)

	return mux
}
