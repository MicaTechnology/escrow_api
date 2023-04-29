package controllers

import (
	"net/http"

	"github.com/MicaTechnology/escrow_api/utils/logger"
)

const (
	pong = "{\"message\": \"pong\"}"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct {
}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	logger.Info("Ping")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "*")
	w.Header().Set("Allow-Control-Allow-Headers", "*")
	w.Header().Set("Allow-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
