package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/herbiboar"
	"osrs-xp-kits/internal/services"
)

// HerbiboarLiveHandler handles herbiboar calculations with live price support
type HerbiboarLiveHandler struct {
	cacheManager *services.CacheManager
}

// NewHerbiboarLiveHandler creates a new herbiboar handler with live price support
func NewHerbiboarLiveHandler(cacheManager *services.CacheManager) *HerbiboarLiveHandler {
	return &HerbiboarLiveHandler{
		cacheManager: cacheManager,
	}
}

// HerbiboarLiveInput extends the basic input with live price options
type HerbiboarLiveInput struct {
	HunterLevel     int    `json:"hunter_level"`
	HerbloreLevel   int    `json:"herblore_level"`
	MagicSecateurs  bool   `json:"magic_secateurs"`
	CalculationType string `json:"calculation_type"` // "target" or "number"
	TargetLevel     *int   `json:"target_level,omitempty"`
	NumberToCatch   *int   `json:"number_to_catch,omitempty"`
	UseLivePrices   bool   `json:"use_live_prices,omitempty"`
	Username        string `json:"username,omitempty"` // Optional: auto-populate skill levels
}

// HerbiboarLiveResponse extends the basic response with price information
type HerbiboarLiveResponse struct {
	herbiboar.HerbiboarResult
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
}

// Calculate handles POST /api/herbiboar/live
func (h *HerbiboarLiveHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var input HerbiboarLiveInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Auto-populate skill levels if username is provided
	hunterLevel := input.HunterLevel
	herbloreLevel := input.HerbloreLevel

	if input.Username != "" {
		osrsAPI := services.NewOSRSAPIService()
		playerStats, err := osrsAPI.GetPlayerStats(input.Username)
		if err != nil {
			// Don't fail, just use provided levels and add warning
			w.Header().Set("X-Player-Stats-Warning", "Could not fetch player stats: "+err.Error())
		} else {
			// Use player stats if they're higher than provided levels
			if playerStats.Hunter > hunterLevel {
				hunterLevel = playerStats.Hunter
			}
			if playerStats.Herblore > herbloreLevel {
				herbloreLevel = playerStats.Herblore
			}
		}
	}

	// Get live prices if requested
	var livePrices map[string]int
	var priceInfo *PriceInfo

	if input.UseLivePrices {
		prices, err := h.cacheManager.GetPrices()
		if err != nil {
			http.Error(w, "Failed to fetch live prices: "+err.Error(), http.StatusInternalServerError)
			return
		}

		livePrices = prices
		cacheStatus := h.cacheManager.GetCacheStatus()

		priceInfo = &PriceInfo{
			Source:     "live",
			PricesUsed: prices,
		}

		if lastUpdated, ok := cacheStatus["last_updated"].(string); ok {
			priceInfo.LastUpdated = lastUpdated
		}
	} else {
		priceInfo = &PriceInfo{
			Source: "static",
		}
	}

	// Convert to domain input
	domainInput := herbiboar.HerbiboarInput{
		HunterLevel:     hunterLevel,
		HerbloreLevel:   herbloreLevel,
		MagicSecateurs:  input.MagicSecateurs,
		CalculationType: input.CalculationType,
		TargetLevel:     input.TargetLevel,
		NumberToCatch:   input.NumberToCatch,
	}

	// Calculate herbiboar data
	result, err := herbiboar.CalculateHerbiboarDataWithPrices(domainInput, livePrices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create enhanced response
	response := HerbiboarLiveResponse{
		HerbiboarResult: result,
		PriceInfo:       priceInfo,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
