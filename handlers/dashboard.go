package handlers

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// DashboardRoutes creates a router for the dashboard routes
func DashboardRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", GetDashboardData)
	return router
}

// GetDashboardData handles the /api/dashboard route
func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	var carbonFootprintResults []models.CarbonFootprintGraph
	var pollutionLevels []models.PollutionLevel

	// Get carbon footprint data
	result := db.GetDB().Table((models.CarbonFootprint{}).TableName()).
		Select("TO_CHAR(month, 'Mon') as month_str, amount").
		Order("month").
		Find(&carbonFootprintResults)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Map the results to the CarbonFootprint struct
	var carbonFootprint []models.CarbonFootprint
	for _, res := range carbonFootprintResults {
		carbonFootprint = append(carbonFootprint, models.CarbonFootprint{
			MonthStr: res.MonthStr,
			Amount:   res.Amount,
		})
	}

	// Get pollution levels
	if err := db.GetDB().Find(&pollutionLevels).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.DashboardResponse{
		CarbonFootprint: carbonFootprint,
		PollutionLevels: pollutionLevels,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
