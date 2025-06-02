package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	ardyknights "osrs-xp-kits/internal/calculators/technique/ardy_knights"
)

type ArdyKnightInput struct {
	CurrentThievingXP    *int `json:"current_thieving_xp,omitempty"`
	CurrentThievingLevel *int `json:"current_thieving_level,omitempty"`

	TargetThievingXP    *int `json:"target_thieving_xp,omitempty"`
	TargetThievingLevel *int `json:"target_thieving_level,omitempty"`

	HasArdyMed        bool `json:"has_ardy_med"`
	HasThievingCape   bool `json:"has_thieving_cape"`
	HasRoguesOutfit   bool `json:"has_rogues_outfit"`
	HasShadowVeil     bool `json:"has_shadow_veil"`
	HourlyPickpockets int  `json:"hourly_pickpockets"`
	FoodHealAmount    int  `json:"food_heal_amount"`
	FoodCost          int  `json:"food_cost"`
}

func ArdyKnightCalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var input ArdyKnightInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	var actualCurrentXP int

	if input.CurrentThievingXP != nil {
		actualCurrentXP = *input.CurrentThievingXP
		if actualCurrentXP < 0 {
			http.Error(w, "Current Thieving XP cannot be negative.", http.StatusBadRequest)
			return
		}
	} else if input.CurrentThievingLevel != nil {
		lvl := *input.CurrentThievingLevel
		if lvl < 1 || lvl > 99 {
			http.Error(w, "Current Thieving Level must be between 1 and 99.", http.StatusBadRequest)
			return
		}
		actualCurrentXP = ardyknights.GetTotalXPForLevel(lvl)
	} else {
		http.Error(w, "Either current_thieving_xp or current_thieving_level must be provided.", http.StatusBadRequest)
		return
	}

	var actualTargetXP int

	if input.TargetThievingXP != nil {
		actualTargetXP = *input.TargetThievingXP
		if actualTargetXP < 0 {
			http.Error(w, "Target Thieving XP cannot be negative.", http.StatusBadRequest)
			return
		}

	} else if input.TargetThievingLevel != nil {
		lvl := *input.TargetThievingLevel
		if lvl < 1 || lvl > 99 { // Max OSRS level
			http.Error(w, "Target Thieving Level must be between 1 and 99.", http.StatusBadRequest)
			return
		}
		actualTargetXP = ardyknights.GetTotalXPForLevel(lvl)
	} else {
		http.Error(w, "Either target_thieving_xp or target_thieving_level must be provided.", http.StatusBadRequest)
		return
	}

	if actualTargetXP <= actualCurrentXP {
		http.Error(w, fmt.Sprintf("Target progress (XP: %d) must be greater than current progress (XP: %d).", actualTargetXP, actualCurrentXP), http.StatusBadRequest)
		return
	}

	fmt.Println(actualCurrentXP,
		actualTargetXP,
		input.HasArdyMed,
		input.HasThievingCape,
		input.HasRoguesOutfit,
		input.HasShadowVeil,
		input.HourlyPickpockets,
		input.FoodHealAmount,
		input.FoodCost)

	result, err := ardyknights.CalculateArdyKnightStats(
		actualCurrentXP,
		actualTargetXP,
		input.HasArdyMed,
		input.HasThievingCape,
		input.HasRoguesOutfit,
		input.HasShadowVeil,
		input.HourlyPickpockets,
		input.FoodHealAmount,
		input.FoodCost,
	)
	if err != nil {
		http.Error(w, "Calculation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
