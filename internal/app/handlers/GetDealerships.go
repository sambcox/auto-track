package handlers

import (
	"encoding/json"
	"github.com/sambcox/auto-track/internal/app/mocks"
	"net/http"
)

func GetAllDealerships(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Dealerships)
}
