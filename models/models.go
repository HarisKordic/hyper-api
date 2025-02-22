package models

import (
	"time"
)

type CarbonFootprint struct {
	ID       uint      `json:"-" gorm:"primaryKey"`
	Month    time.Time `json:"-"`
	MonthStr string    `json:"month" gorm:"-"`
	Amount   float64   `json:"amount"`
}

func (CarbonFootprint) TableName() string {
	return "carbon_footprint_history"
}

type CarbonFootprintGraph struct {
	MonthStr string  `json:"month"`
	Amount   float64 `json:"amount"`
}

type PollutionLevel struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	Pollutant string `json:"pollutant"`
	Level     int    `json:"level"`
}

type DashboardResponse struct {
	CarbonFootprint []CarbonFootprint `json:"carbonFootprint"`
	PollutionLevels []PollutionLevel  `json:"pollutionLevels"`
}
