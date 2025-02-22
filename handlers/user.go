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
	router.Get("/email/{email}", GetUserByEmail)
	return router
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.GetDB().Order("id desc").Limit(5).Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	var user models.User

	result := db.GetDB().Where("email = ?", email).First(&user)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
