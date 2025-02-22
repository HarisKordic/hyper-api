package handlers

import (
	"encoding/json"
	"hyper-api/db"
	"hyper-api/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupTestDB initializes and returns a connection to the test database
func setupTestDB(t *testing.T) *gorm.DB {
	// Load the .env file
	err := godotenv.Load("../.env") // Adjust the path to your .env file if needed
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	// Load the DATABASE_URL from the environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatal("DATABASE_URL environment variable is not set")
	}

	// Log the DATABASE_URL (for debugging)
	t.Logf("Connecting to database with DATABASE_URL: %s", dsn)

	// Connect to the database
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Log a success message if the connection is established
	t.Log("Successfully connected to the database")

	// Set the global DB instance
	db.SetDB(gormDB)

	return gormDB
}

// TestDatabaseConnectivity tests and logs database connectivity
func TestDatabaseConnectivity(t *testing.T) {
	// Set up the test database
	db := setupTestDB(t)

	// Test if the database is reachable by running a simple query
	var result int
	err := db.Raw("SELECT 1").Scan(&result).Error
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}

	// Log the result of the query
	t.Logf("Database connectivity test result: SELECT 1 => %d", result)

	// Assert that the query returned the expected result
	assert.Equal(t, 1, result, "Database connectivity test failed: expected 1, got %d", result)
}

// TestGetDashboardData tests the GetDashboardData handler
func TestGetDashboardData(t *testing.T) {
	// Set up the test database
	setupTestDB(t)

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	GetDashboardData(rr, req)

	// Print the response body for debugging
	t.Logf("Response Body: %s", rr.Body.String())

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")

	// Decode the response body
	var response models.DashboardResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Print the decoded response for debugging
	t.Logf("Decoded Response: %+v", response)

	// Assert that the response contains the expected data
	assert.NotEmpty(t, response.CarbonFootprint, "CarbonFootprint data should not be empty")
	assert.NotEmpty(t, response.PollutionLevels, "PollutionLevels data should not be empty")
}
