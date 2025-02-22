package tests

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"hyper-api/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB
var testServer *httptest.Server

func setupTestDB(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetDB(testDB)
}

func TestMain(m *testing.M) {
	// Set up the test server
	router := server.NewRouter()
	testServer = httptest.NewServer(router)
	defer testServer.Close()

	os.Exit(m.Run())
}

func TestGetDashboard(t *testing.T) {
	setupTestDB(t)

	resp, err := http.Get(testServer.URL + "/api/dashboard")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var dashboardData models.DashboardResponse
	err = json.NewDecoder(resp.Body).Decode(&dashboardData)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Validate dashboard data
	assert.NotEmpty(t, dashboardData.CarbonFootprint)
	assert.NotEmpty(t, dashboardData.PollutionLevels)
}

func TestGetMapUsers(t *testing.T) {
	setupTestDB(t)

	resp, err := http.Get(testServer.URL + "/api/map/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var mapUsers []models.MapUser
	err = json.NewDecoder(resp.Body).Decode(&mapUsers)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Validate map users data
	assert.NotEmpty(t, mapUsers)
	for _, user := range mapUsers {
		assert.NotEmpty(t, user.Name)
		assert.NotEmpty(t, user.Avatar)
		assert.NotZero(t, user.Latitude)
		assert.NotZero(t, user.Longitude)
	}
}

func TestGetUsers(t *testing.T) {
	setupTestDB(t)

	resp, err := http.Get(testServer.URL + "/api/users")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Validate users data
	assert.NotEmpty(t, users)
	assert.LessOrEqual(t, len(users), 5) // Check limit of 5 users
}

func TestGetUserByEmail(t *testing.T) {
	setupTestDB(t)

	// First get a list of users to get a valid email
	resp, err := http.Get(testServer.URL + "/api/users")
	assert.NoError(t, err)

	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(t, err)
	resp.Body.Close()

	if len(users) > 0 {
		// Test with existing email
		resp, err = http.Get(testServer.URL + "/api/users/email/" + users[0].Email)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var user models.User
		err = json.NewDecoder(resp.Body).Decode(&user)
		assert.NoError(t, err)
		assert.Equal(t, users[0].Email, user.Email)
		resp.Body.Close()
	}

	// Test with non-existent email
	resp, err = http.Get(testServer.URL + "/api/users/email/nonexistent@email.com")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
