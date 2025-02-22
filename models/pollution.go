package models

type PollutionLevel struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Pollutant string `json:"pollutant"`
	Level     int    `json:"level"`
}
