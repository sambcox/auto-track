package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sambcox/auto-track/internal/app/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dealerships", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World")
	})

	log.Println("Time to show off some metal!")
	http.ListenAndServe(":4000", router)
}
