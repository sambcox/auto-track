package models

import "time"

type Car struct {
	ID           int       `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Price        float64   `json:"price"`
	Color        string    `json:"color"`
	Status       string    `json:"status"`
	DealershipID int       `json:"dealership_id"`
	CreatedAt    time.Time `json:"created_at"`
}
