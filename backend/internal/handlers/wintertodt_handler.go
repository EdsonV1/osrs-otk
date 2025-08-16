package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/wintertodt"
)

type WintertodtInput struct {
	CurrentLevel          int                    `json:"current_level"`
	TargetLevel           int                    `json:"target_level"`
	Strategy              string                 `json:"strategy"`
	CustomPointsPerRound  *int                   `json:"custom_points_per_round,omitempty"`
	CustomMinutesPerRound *float64               `json:"custom_minutes_per_round,omitempty"`
	SkillLevels           wintertodt.SkillLevels `json:"skill_levels"`
	UseLivePrices         bool                   `json:"use_live_prices,omitempty"`
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

	// Convert strategy string to Strategy type
	strategy := wintertodt.Strategy(input.Strategy)

	result, err := wintertodt.CalculateWintertodtData(
		input.CurrentLevel,
		input.TargetLevel,
		strategy,
		input.CustomPointsPerRound,
		input.CustomMinutesPerRound,
		input.SkillLevels,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// WintertodtProTipsHandler provides detailed calculation methodology and tips
func WintertodtProTipsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	tips := wintertodt.GetCalculationProTips()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tips)
}
