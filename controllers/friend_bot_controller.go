package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/MicaTechnology/escrow_api/domains/friend_bot"
	"github.com/MicaTechnology/escrow_api/services"
	"github.com/MicaTechnology/escrow_api/utils/http_utils"
	"github.com/MicaTechnology/escrow_api/utils/logger"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
)

var (
	FriendBotController friendBotControllerInterface = &friendBotController{}
)

const (
	success = "{\"message\": \"OK\"}"
)

type friendBotControllerInterface interface {
	AddFunds(w http.ResponseWriter, r *http.Request)
}

type friendBotController struct {
}

func (c *friendBotController) AddFunds(w http.ResponseWriter, r *http.Request) {
	logger.Info("Add Funds")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "*")
	w.Header().Set("Allow-Control-Allow-Headers", "*")
	w.Header().Set("Allow-Control-Allow-Credentials", "true")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}
	defer r.Body.Close()

	var fundRequest friend_bot.FundRequest
	if err := json.Unmarshal(requestBody, &fundRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid fund request json body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}

	fundErr := services.FriendBotService.AddFunds(fundRequest)
	if fundErr != nil {
		http_utils.ResponseJsonError(w, fundErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(success))
}
