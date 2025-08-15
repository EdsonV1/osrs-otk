package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"osrs-xp-kits/internal/services"
)

func TestGetPlayerStats_URLDecoding(t *testing.T) {
	// Create test handler
	osrsAPI := services.NewOSRSAPIService()
	cacheManager := services.NewCacheManager("./test-cache", osrsAPI)
	handlers := NewAPIHandlers(osrsAPI, cacheManager)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		shouldContain  string
	}{
		{
			name:           "Normal username without spaces",
			path:           "/api/player-stats/Zezima",
			expectedStatus: 404, // Player might not exist, but URL parsing should work
			shouldContain:  "not found",
		},
		{
			name:           "Username with encoded space",
			path:           "/api/player-stats/Lynx%20Titan",
			expectedStatus: 404, // Player might not exist, but URL parsing should work
			shouldContain:  "not found",
		},
		{
			name:           "Empty username",
			path:           "/api/player-stats/",
			expectedStatus: 400,
			shouldContain:  "Username is required",
		},
		{
			name:           "Username with special characters",
			path:           "/api/player-stats/Test%21%40%23",
			expectedStatus: 404, // Should decode Test!@# and try to look it up
			shouldContain:  "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handlers.GetPlayerStats(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, status)
			}

			var response PlayerStatsResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Errorf("Failed to parse response: %v", err)
			}

			if response.Success {
				t.Errorf("Expected success to be false for test cases")
			}

			// For URL decoding test, we mainly care that it doesn't crash
			// and returns a proper error format
			if response.Error == "" {
				t.Errorf("Expected error message in response")
			}
		})
	}
}

func TestGetPlayerStats_LongUsername(t *testing.T) {
	osrsAPI := services.NewOSRSAPIService()
	cacheManager := services.NewCacheManager("./test-cache", osrsAPI)
	handlers := NewAPIHandlers(osrsAPI, cacheManager)

	// Test with very long username (OSRS usernames are max 12 chars)
	longUsername := "VeryLongUsernameExceedingLimit"
	req, err := http.NewRequest("GET", "/api/player-stats/"+longUsername, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlers.GetPlayerStats(rr, req)

	// Should still process but likely fail at API level
	if status := rr.Code; status != http.StatusBadRequest && status != http.StatusNotFound {
		t.Errorf("Expected status 400 or 404 for long username, got %d", status)
	}

	var response PlayerStatsResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response: %v", err)
	}

	if response.Success {
		t.Errorf("Expected success to be false for long username")
	}
}