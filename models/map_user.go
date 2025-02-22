package models

import "time"

type MapUser struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Avatar         string    `json:"avatar"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Activity       string    `json:"activity"`
	Timestamp      time.Time `json:"timestamp"`
	IconType       string    `json:"iconType"`
	EmissionAmount float64   `json:"emission_amount"`
}
