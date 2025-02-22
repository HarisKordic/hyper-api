package models

type DashboardResponse struct {
	CarbonFootprint []CarbonFootprint `json:"carbonFootprint"`
	PollutionLevels []PollutionLevel  `json:"pollutionLevels"`
}
