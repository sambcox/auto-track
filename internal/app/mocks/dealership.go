package mocks

import (
	"github.com/sambcox/auto-track/internal/app/models"
	"time"
)

var Dealerships = []models.Dealership{
	{
		ID:             1,
		Name:           "ABC Motors",
		Location:       "123 Main Street, City, Country",
		EstablishedDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:      time.Now(),
	},
}
