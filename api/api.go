package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db/model"
	"simple-microservice-backend/pkg/response"
	"strconv"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type APIHandler struct {
	opts APIHandlerOptions
}

type APIHandlerOptions struct {
	serverOptions *config.ServerOptions
}

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type structuredError struct {
	Code int    `json:"data,omitempty"`
	Msg  string `json:"message,omitempty"`
}

type structuredResponse struct {
	Data   interface{}       `json:"data"`
	Errors []structuredError `json:"errors"`
}

type ApiResponse struct {
	Status    status
	Data      interface{}
	ErrorType response.ErrorType
	Error     string
}

func NewAPIHandler() (*APIHandler, error) {
	opts, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("Error initialising API Handler :", err)
		return nil, err
	}

	return &APIHandler{opts: APIHandlerOptions{serverOptions: opts}}, nil
}

func newApiErrorBadData(err error) *response.ApiError {
	return &response.ApiError{Typ: response.ErrorBadData, Err: err}
}

func newApiErrorInvalidData(err error) *response.ApiError {
	return &response.ApiError{Typ: response.ErrorInternal, Err: err}
}

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-type", "application/json")
	return w
}

func (aH *APIHandler) writeJSON(w http.ResponseWriter, r *http.Request, response interface{}) {
	marshall := json.Marshal
	if pretty := r.FormValue("pretty"); pretty != "" && pretty != "false" {
		marshall = func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "    ")
		}
	}
	resp, _ := marshall(response)
	setHeader(w)
	w.Write(resp)

}

func RespondError(w http.ResponseWriter, apiErr response.BaseApiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	var code int
	switch apiErr.Type() {
	case response.ErrorBadData:
		code = http.StatusBadRequest
	case response.ErrorInternal:
		code = http.StatusInternalServerError
	case response.ErrorExec:
		code = 422
	case response.ErrorCanceled, response.ErrorTimeout:
		code = http.StatusServiceUnavailable
	case response.ErrorForbidden, response.ErrorUnauthorized:
		code = http.StatusForbidden
	default:
		code = http.StatusInternalServerError
	}

	setHeader(w)
	w.WriteHeader(code)

	b, err := json.Marshal(&ApiResponse{Status: statusError, Data: data, ErrorType: apiErr.Type(), Error: apiErr.Error()})
	if err != nil {
		http.Error(w, "Error Marshalling the response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Fatalf("Error writing response :%s\n", err)
	}
}

func writeHttpResponse(w http.ResponseWriter, Data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&ApiResponse{Status: statusSuccess, Data: Data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	setHeader(w)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		log.Fatalf("Error writing response :%s\n", err)
	}

}

func (aH *APIHandler) Respond(w http.ResponseWriter, data interface{}) {
	writeHttpResponse(w, data)

}

func (aH *APIHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	writeHttpResponse(w, "Welcome to Home page")
}

func (aH *APIHandler) GetAccountByCRN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crn, err := strconv.ParseUint(vars["crn"], 10, 32)
	if err != nil {
		RespondError(w, newApiErrorBadData(errors.New("invalid crn")), nil)
	}

	dBInstance, ok := r.Context().Value(dbContextKey).(*gorm.DB)
	if !ok || dBInstance == nil {
		RespondError(w, newApiErrorBadData(errors.New("database connection not available")), nil)
		return
	}

	var accountMaster model.AccountMaster

	if err := dBInstance.Preload("AccOwner").Where("acc_owner_id=(select id from owners where cr_number=?)", crn).First(&accountMaster).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RespondError(w, newApiErrorBadData(errors.New("account not found")), nil)
		} else if errors.Is(err, context.DeadlineExceeded) {
			RespondError(w, newApiErrorBadData(errors.New("request timed out")), nil)
		} else {
			RespondError(w, newApiErrorBadData(errors.New("internal server error")), nil)
		}
		return
	}

	aH.Respond(w, accountMaster)
}
