package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sambcox/auto-track/internal/app/mocks"
)

func DeleteDealership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, dealership := range mocks.Dealerships {
		if dealership.ID == id {
			mocks.Dealerships = append(mocks.Dealerships[:index], mocks.Dealerships[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dealership)
			break
		}
	}
}
