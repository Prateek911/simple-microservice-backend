package api

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-microservice-backend/config"
	"simple-microservice-backend/internal/response"

	jsoniter "github.com/json-iterator/go"
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

func setHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-type", "application/json")
	return w
}

func (aH *APIHandler) writeJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	marshall := json.Marshal
	if pretty := r.FormValue("pretty"); pretty != "" && pretty != "false" {
		marshall = func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "    ")
		}
	}
	resp, _ := marshall(r)
	setHeader(w)
	w.Write(resp)

}

func RespondError(w http.ResponseWriter, apiErr response.BaseApiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&ApiResponse{Status: statusError, Data: data, ErrorType: apiErr.Type(), Error: apiErr.Error()})
	if err != nil {
		http.Error(w, "Error Marshalling the response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Fatalf("Error writing response :%s\n", err)
	}
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
