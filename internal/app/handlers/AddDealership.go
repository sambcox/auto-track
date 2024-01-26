package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/sambcox/auto-track/internal/app/mocks"
	"github.com/sambcox/auto-track/internal/app/models"
)

func AddDealership(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var dealership models.Dealership
	json.Unmarshal(body, &dealership)

	dealership.ID = rand.Intn(100)
	mocks.Dealerships = append(mocks.Dealerships, dealership)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dealership)
}
