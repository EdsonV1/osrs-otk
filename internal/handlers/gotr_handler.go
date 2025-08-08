package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/gotr"
)

// GOTRInput represents the input structure for GOTR calculations
type GOTRInput struct {
	CurrentLevel int `json:"current_level"`
	TargetLevel  int `json:"target_level"`
}

// GOTRCalcHandler handles HTTP requests for GOTR calculations
func GOTRCalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var input GOTRInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if input.CurrentLevel < 27 || input.CurrentLevel > 126 {
		http.Error(w, "Current level must be between 27 and 126 (minimum level to access GOTR)", http.StatusBadRequest)
		return
	}

	if input.TargetLevel < 27 || input.TargetLevel > 126 {
		http.Error(w, "Target level must be between 27 and 126", http.StatusBadRequest)
		return
	}

	if input.TargetLevel <= input.CurrentLevel {
		http.Error(w, "Target level must be higher than current level", http.StatusBadRequest)
		return
	}

	// Calculate GOTR data
	result, err := gotr.CalculateGOTRData(input.CurrentLevel, input.TargetLevel)
	if err != nil {
		http.Error(w, "Calculation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode and send response
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// GOTRStrategyHandler provides strategic advice for GOTR training
func GOTRStrategyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		CurrentLevel int `json:"current_level"`
		TargetLevel  int `json:"target_level"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get detailed time breakdown and strategy
	breakdown, err := gotr.EstimateTimeToLevel(input.CurrentLevel, input.TargetLevel)
	if err != nil {
		http.Error(w, "Strategy calculation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breakdown)
}
