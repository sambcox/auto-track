package models

import "time"

type CarImage struct {
	ID        int       `json:"id"`
	CarID     int       `json:"car_id"`
	ImagePath string    `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
}
