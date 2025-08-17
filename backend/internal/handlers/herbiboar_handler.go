package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/herbiboar"
)

type HerbiboarInput struct {
	HunterLevel     int    `json:"hunter_level"`
	HerbloreLevel   int    `json:"herblore_level"`
	MagicSecateurs  bool   `json:"magic_secateurs"`
	CalculationType string `json:"calculation_type"` // "target" or "number"
	TargetLevel     *int   `json:"target_level,omitempty"`
	NumberToCatch   *int   `json:"number_to_catch,omitempty"`
}

func HerbiboarCalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input HerbiboarInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Convert to domain input
	domainInput := herbiboar.HerbiboarInput{
		HunterLevel:     input.HunterLevel,
		HerbloreLevel:   input.HerbloreLevel,
		MagicSecateurs:  input.MagicSecateurs,
		CalculationType: input.CalculationType,
		TargetLevel:     input.TargetLevel,
		NumberToCatch:   input.NumberToCatch,
	}

	result, err := herbiboar.CalculateHerbiboarData(domainInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// HerbiboarProTipsHandler provides detailed calculation methodology and tips
func HerbiboarProTipsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	tips := herbiboar.GetCalculationProTips()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tips)
}
