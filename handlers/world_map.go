// handlers/world_map.go
package handlers

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	mapUsersCache    []models.MapUser
	mapUsersCacheExp time.Time
	cacheMutex       sync.RWMutex
	cacheTimeout     = 5 * time.Minute
)

func MapRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/users", GetMapUsers)
	return router
}

// GetMapUsers handles the /api/map/users route
func GetMapUsers(w http.ResponseWriter, r *http.Request) {
	cacheMutex.RLock()
	if mapUsersCache != nil && time.Now().Before(mapUsersCacheExp) {
		cacheMutex.RUnlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mapUsersCache)
		return
	}
	cacheMutex.RUnlock()

	var mapUsers []models.MapUser
	result := db.GetDB().Find(&mapUsers)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	cacheMutex.Lock()
	mapUsersCache = mapUsers
	mapUsersCacheExp = time.Now().Add(cacheTimeout)
	cacheMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mapUsers)
}
