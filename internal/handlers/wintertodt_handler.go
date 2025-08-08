package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/wintertodt"
)

type WintertodtInput struct {
	FiremakingLevel int     `json:"firemaking_level"`
	RoundsPerHour   float64 `json:"rounds_per_hour"`
	TotalRounds     int     `json:"total_rounds"`
}

func WintertodtCalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input WintertodtInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := wintertodt.CalculateWintertodtData(input.FiremakingLevel, input.RoundsPerHour, input.TotalRounds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
