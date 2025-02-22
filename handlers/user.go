package handlers

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", GetUsers)
	return router
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.GetDB().Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
