package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sambcox/auto-track/internal/app/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/dealerships", handlers.GetAllDealerships).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/dealerships", handlers.AddDealership).Methods(http.MethodPost)

	log.Println("Time to show off some metal!")
	http.ListenAndServe(":4000", router)
}
