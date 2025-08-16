package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/birdhouses"
)

type BirdhouseInput struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func BirdhouseCalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var input BirdhouseInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := birdhouses.CalculateBirdhouseData(input.Type, input.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// BirdhouseProTipsHandler provides detailed calculation methodology and tips
func BirdhouseProTipsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	tips := birdhouses.GetCalculationProTips()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tips)
}
