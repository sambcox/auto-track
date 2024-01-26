package models

import "time"

type Dealership struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Location        string    `json:"location"`
	EstablishedDate string    `json:"established_date"`
	CreatedAt       time.Time `json:"created_at"`
}
