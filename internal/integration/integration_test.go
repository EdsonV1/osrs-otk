package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"osrs-xp-kits/internal/config"
	"osrs-xp-kits/internal/server"
)

// TestServer holds the test server instance
var testServer *httptest.Server

// Setup test server before running integration tests
func TestMain(m *testing.M) {
	// Create test configuration
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: "8080",
		},
		CORS: config.CORSConfig{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		},
		Assets: config.AssetsConfig{
			SkillDataPath: "../../assets/data/skills",
		},
	}

	// Create server
	srv := server.New(cfg)

	// Create test server using the CORS middleware handler
	handler := srv.GetHandler() // We need to add this method
	testServer = httptest.NewServer(handler)
	defer testServer.Close()

	// Run tests
	m.Run()
}

// TestHealthEndpoint tests if the server is responding
func TestHealthEndpoint(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/api/skill-data/hunter")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

// TestGOTREndpoint tests the GOTR calculator endpoint
func TestGOTREndpoint(t *testing.T) {
	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
		shouldHaveXP   bool
	}{
		{
			name: "Valid GOTR calculation",
			payload: map[string]interface{}{
				"current_level": 77,
				"target_level":  99,
			},
			expectedStatus: http.StatusOK,
			shouldHaveXP:   true,
		},
		{
			name: "Invalid level too low",
			payload: map[string]interface{}{
				"current_level": 20,
				"target_level":  99,
			},
			expectedStatus: http.StatusBadRequest,
			shouldHaveXP:   false,
		},
		{
			name: "Invalid target lower than current",
			payload: map[string]interface{}{
				"current_level": 90,
				"target_level":  80,
			},
			expectedStatus: http.StatusBadRequest,
			shouldHaveXP:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal payload to JSON
			jsonPayload, err := json.Marshal(tt.payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}

			// Make request
			resp, err := http.Post(
				testServer.URL+"/api/tools/gotr",
				"application/json",
				bytes.NewBuffer(jsonPayload),
			)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			// Check status code
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			// For successful requests, check response structure
			if tt.shouldHaveXP && resp.StatusCode == http.StatusOK {
				var result map[string]interface{}
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}

				// Check required fields
				requiredFields := []string{
					"current_level", "target_level", "xp_needed",
					"games_needed", "hours_needed", "average_xp_per_hour",
					"pet_chance_percentage", "estimated_rewards", "total_reward_value",
				}

				for _, field := range requiredFields {
					if _, exists := result[field]; !exists {
						t.Errorf("Response missing required field: %s", field)
					}
				}

				// Validate data types and ranges
				if xpNeeded, ok := result["xp_needed"].(float64); ok {
					if xpNeeded <= 0 {
						t.Errorf("XP needed should be positive, got %f", xpNeeded)
					}
				}

				if petChance, ok := result["pet_chance_percentage"].(float64); ok {
					if petChance < 0 || petChance > 100 {
						t.Errorf("Pet chance should be 0-100%%, got %f", petChance)
					}
				}
			}
		})
	}
}

// TestWintertodtEndpoint tests the Wintertodt calculator endpoint
func TestWintertodtEndpoint(t *testing.T) {
	payload := map[string]interface{}{
		"firemaking_level": 75,
		"rounds_per_hour":  5.0,
		"total_rounds":     100,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	resp, err := http.Post(
		testServer.URL+"/api/wintertodt",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check required fields
	requiredFields := []string{
		"total_experience", "average_exp_hour", "pet_chance",
		"estimated_loot", "total_value", "total_time",
	}

	for _, field := range requiredFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Response missing required field: %s", field)
		}
	}
}

// TestArdyKnightsEndpoint tests the Ardy Knights calculator endpoint
func TestArdyKnightsEndpoint(t *testing.T) {
	payload := map[string]interface{}{
		"current_thieving_xp":  500000,
		"target_thieving_xp":   1000000,
		"has_ardy_med":         true,
		"has_thieving_cape":    false,
		"has_rogues_outfit":    true,
		"has_shadow_veil":      false,
		"hourly_pickpockets":   4000,
		"food_heal_amount":     20,
		"food_cost":            400,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	resp, err := http.Post(
		testServer.URL+"/api/ardyknights",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

// TestBirdhouseEndpoint tests the Birdhouse calculator endpoint
func TestBirdhouseEndpoint(t *testing.T) {
	payload := map[string]interface{}{
		"type":     "yew",
		"quantity": 100,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	resp, err := http.Post(
		testServer.URL+"/api/birdhouse",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Check required fields
	requiredFields := []string{
		"estimated_nests", "hunter_xp", "crafting_xp",
		"seed_drops", "total_loot",
	}

	for _, field := range requiredFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Response missing required field: %s", field)
		}
	}
}

// TestCORSHeaders tests that CORS headers are properly set
func TestCORSHeaders(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", testServer.URL+"/api/tools/gotr", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Origin", "http://localhost:3000")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check CORS headers
	allowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	if allowOrigin == "" {
		t.Error("Missing Access-Control-Allow-Origin header")
	}

	allowMethods := resp.Header.Get("Access-Control-Allow-Methods")
	if allowMethods == "" {
		t.Error("Missing Access-Control-Allow-Methods header")
	}
}

// TestConcurrentRequests tests the server under concurrent load
func TestConcurrentRequests(t *testing.T) {
	const numRequests = 50
	const concurrency = 10

	payload := map[string]interface{}{
		"current_level": 77,
		"target_level":  99,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	// Channel to collect results
	results := make(chan error, numRequests)

	// Worker function
	worker := func() {
		for i := 0; i < numRequests/concurrency; i++ {
			resp, err := http.Post(
				testServer.URL+"/api/tools/gotr",
				"application/json",
				bytes.NewBuffer(jsonPayload),
			)
			if err != nil {
				results <- fmt.Errorf("request failed: %v", err)
				continue
			}
			resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- fmt.Errorf("expected status 200, got %d", resp.StatusCode)
				continue
			}

			results <- nil
		}
	}

	// Start concurrent workers
	start := time.Now()
	for i := 0; i < concurrency; i++ {
		go worker()
	}

	// Collect results
	var errors []error
	for i := 0; i < numRequests; i++ {
		if err := <-results; err != nil {
			errors = append(errors, err)
		}
	}
	elapsed := time.Since(start)

	// Report results
	if len(errors) > 0 {
		t.Errorf("Concurrent test had %d errors out of %d requests:", len(errors), numRequests)
		for i, err := range errors {
			if i < 5 { // Show first 5 errors
				t.Errorf("  Error %d: %v", i+1, err)
			}
		}
	}

	successRate := float64(numRequests-len(errors)) / float64(numRequests) * 100
	t.Logf("Concurrent test completed: %d requests in %v (%.1f%% success rate)",
		numRequests, elapsed, successRate)

	if successRate < 95 {
		t.Errorf("Success rate too low: %.1f%% (expected >= 95%%)", successRate)
	}
}