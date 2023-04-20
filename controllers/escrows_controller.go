package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/MicaTechnology/escrow_api/domains/escrows"
	"github.com/MicaTechnology/escrow_api/services"
	"github.com/MicaTechnology/escrow_api/utils/http_utils"
	"github.com/MicaTechnology/escrow_api/utils/rest_errors"
	"github.com/gorilla/mux"
)

var (
	EscrowController escrowControllerInterface = &escrowController{}
)

type escrowControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type escrowController struct {
}

func (c *escrowController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}
	defer r.Body.Close()

	var escrow escrows.Escrow
	if err := json.Unmarshal(requestBody, &escrow); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid escrow json body")
		http_utils.ResponseJsonError(w, respErr)
		return
	}

	result, createErr := services.EscrowsService.Create(escrow)
	if createErr != nil {
		http_utils.ResponseJsonError(w, createErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *escrowController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	escrowId := strings.TrimSpace(vars["id"])

	escrow, getErr := services.EscrowsService.Get(escrowId)
	if getErr != nil {
		http_utils.ResponseJsonError(w, getErr)
		return
	}
	http_utils.ResponseJson(w, http.StatusOK, escrow)
}
