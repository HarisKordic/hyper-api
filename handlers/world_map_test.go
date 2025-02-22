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

// setupWorldMapTestDB initializes and returns a connection to the test database
func setupWorldMapTestDB(t *testing.T) *gorm.DB {
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

// TestGetMapUsers tests the GetMapUsers handler
func TestGetMapUsers(t *testing.T) {
	// Set up the test database
	setupWorldMapTestDB(t)

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	GetMapUsers(rr, req)

	// Print the response body for debugging
	t.Logf("Response Body: %s", rr.Body.String())

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status code")

	// Decode the response body
	var mapUsers []models.MapUser
	err = json.NewDecoder(rr.Body).Decode(&mapUsers)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Print the decoded response for debugging
	t.Logf("Decoded Response: %+v", mapUsers)

	// Assert that the response contains more than 0 users
	assert.Greater(t, len(mapUsers), 0, "Expected more than 0 users, got %d", len(mapUsers))

	// Optionally, check for specific fields in the response
	if len(mapUsers) > 0 {
		for i, user := range mapUsers {
			// Check if the Name field is a string
			assert.IsType(t, "", user.Name, "Name field in user %d is not a string", i)
		}
	}
	// Check for unique user IDs
	userIDs := make(map[uint]bool)
	for _, user := range mapUsers {
		assert.False(t, userIDs[user.ID], "Duplicate user ID found: %d", user.ID)
		userIDs[user.ID] = true
	}
}
