package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sambcox/auto-track/internal/app/mocks"
	"github.com/sambcox/auto-track/internal/app/models"
)

func UpdateDealership(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedDealership models.Dealership
	json.Unmarshal(body, &updatedDealership)

	for index, dealership := range mocks.Dealerships {
		if dealership.ID == id {
			dealership.Name = updatedDealership.Name
			dealership.Location = updatedDealership.Location
			dealership.EstablishedDate = updatedDealership.EstablishedDate

			mocks.Dealerships[index] = dealership
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(dealership)
			break
		}
	}
}
