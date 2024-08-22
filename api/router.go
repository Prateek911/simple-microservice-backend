package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() *mux.Router {
	aH, err := NewAPIHandler()
	if err != nil {
		log.Fatal("Error setting up routing :", err)
		return mux.NewRouter()
	}

	r := mux.NewRouter()
	r.Use(ApplicationContext)

	r.HandleFunc("/", aH.HomeHandler).Methods(http.MethodGet).Name("Home")
	r.HandleFunc("/account/{crn}", aH.GetAccountByCRN).Methods(http.MethodGet).Name("GetAccountByCRN")
	r.HandleFunc("/owner", aH.CreateOwner).Methods(http.MethodPost)
	r.HandleFunc("/owners/{id}", aH.GetOwnerByID).Methods(http.MethodGet)

	return r

}
