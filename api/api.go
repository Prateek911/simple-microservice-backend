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
	entitybuilder "simple-microservice-backend/pkg/service/entityBuilder"
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
	Code int    `json:"code,omitempty"`
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
	case response.ErrorCanceled, response.ErrorTimeout, response.ErrorUnavailable:
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

func respondStructuredError(w http.ResponseWriter, statusCode int, data interface{}, errors []structuredError) {
	response := structuredResponse{
		Data:   data,
		Errors: errors,
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	r, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, "Error marshalling respone", http.StatusInternalServerError)
		return
	}

	setHeader(w)
	w.WriteHeader(statusCode)

	if _, err := w.Write(r); err != nil {
		log.Printf("Error writing response :%s\n", err)
	}
}

func (aH *APIHandler) RespondStructuredError(w http.ResponseWriter, statusCode int, data interface{}, errors []structuredError) {
	respondStructuredError(w, statusCode, data, errors)
}

func (aH *APIHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	writeHttpResponse(w, "Welcome to Home page")
}

func (aH *APIHandler) GetAccountByCRN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crn, err := strconv.ParseUint(vars["crn"], 10, 32)
	if err != nil {
		RespondError(w, &response.ApiError{Typ: response.ErrorBadData, Err: err}, nil)
	}

	dBInstance, ok := r.Context().Value(dbContextKey).(*gorm.DB)
	if !ok || dBInstance == nil {
		http.Error(w, "Database instance not found", http.StatusInternalServerError)
		return
	}

	var accountMaster model.AccountMaster

	if err := dBInstance.Preload("AccOwner").Where("acc_owner_id=(select id from owners where cr_number=?)", crn).First(&accountMaster).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			RespondError(w, response.NotFoundError(err), nil)
		} else if errors.Is(err, context.DeadlineExceeded) {
			RespondError(w, response.RequestTimeOut(err), nil)
		} else {
			RespondError(w, response.BadRequest(err), nil)
		}
		return
	}

	aH.Respond(w, accountMaster)
}

func (aH *APIHandler) CreateOwner(w http.ResponseWriter, r *http.Request) {
	var request response.OwnerCreate
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		RespondError(w, response.BadRequest(err), "Invalid Payload")
		return
	}

	owner := entitybuilder.CreateOwner(request)

	dBInstance, ok := r.Context().Value(dbContextKey).(*gorm.DB)
	if !ok || dBInstance == nil {
		http.Error(w, "Database instance not found", http.StatusInternalServerError)
		return
	}

	if err := dBInstance.Create(owner).Error; err != nil {
		respondStructuredError(w, http.StatusInternalServerError, nil, []structuredError{
			{Code: http.StatusInternalServerError, Msg: err.Error()},
		})
		return
	}

	aH.Respond(w, owner)
}
