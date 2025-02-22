package handlers

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MapRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/users", GetMapUsers)
	return router
}

// GetMapUsers handles the /api/map/users route
func GetMapUsers(w http.ResponseWriter, r *http.Request) {
	var mapUsers []models.MapUser

	result := db.GetDB().Find(&mapUsers)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mapUsers)
}
