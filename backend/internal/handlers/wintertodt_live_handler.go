package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/wintertodt"
	"osrs-xp-kits/internal/services"
)

// WintertodtLiveHandler handles Wintertodt calculations with live price support
type WintertodtLiveHandler struct {
	cacheManager *services.CacheManager
}

// NewWintertodtLiveHandler creates a new Wintertodt handler with live price support
func NewWintertodtLiveHandler(cacheManager *services.CacheManager) *WintertodtLiveHandler {
	return &WintertodtLiveHandler{
		cacheManager: cacheManager,
	}
}

// WintertodtLiveInput extends the basic input with live price options
type WintertodtLiveInput struct {
	CurrentLevel          int                    `json:"current_level"`
	TargetLevel           int                    `json:"target_level"`
	Strategy              string                 `json:"strategy"`
	CustomPointsPerRound  *int                   `json:"custom_points_per_round,omitempty"`
	CustomMinutesPerRound *float64               `json:"custom_minutes_per_round,omitempty"`
	SkillLevels           wintertodt.SkillLevels `json:"skill_levels"`
	UseLivePrices         bool                   `json:"use_live_prices,omitempty"`
	Username              string                 `json:"username,omitempty"` // Optional: auto-populate skill levels
}

// WintertodtLiveResponse extends the basic response with price information
type WintertodtLiveResponse struct {
	wintertodt.WintertodtResult
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
}

// PriceInfo contains information about the prices used in calculation
type PriceInfo struct {
	Source      string         `json:"source"` // "live" or "static"
	LastUpdated string         `json:"last_updated,omitempty"`
	PricesUsed  map[string]int `json:"prices_used,omitempty"`
}

// Calculate handles POST /api/wintertodt/live
func (h *WintertodtLiveHandler) Calculate(w http.ResponseWriter, r *http.Request) {
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

	var input WintertodtLiveInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Auto-populate skill levels if username is provided
	skillLevels := input.SkillLevels
	if input.Username != "" {
		osrsAPI := services.NewOSRSAPIService()
		playerStats, err := osrsAPI.GetPlayerStats(input.Username)
		if err != nil {
			// Don't fail, just use provided skill levels and add warning
			w.Header().Set("X-Player-Stats-Warning", "Could not fetch player stats: "+err.Error())
		} else {
			// Convert player stats to skill levels
			skillLevels = ConvertPlayerStatsToSkillLevels(playerStats)
		}
	}

	// Convert strategy string to Strategy type
	strategy := wintertodt.Strategy(input.Strategy)

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

	// Calculate Wintertodt data
	result, err := wintertodt.CalculateWintertodtDataWithPrices(
		input.CurrentLevel,
		input.TargetLevel,
		strategy,
		input.CustomPointsPerRound,
		input.CustomMinutesPerRound,
		skillLevels,
		livePrices,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create enhanced response
	response := WintertodtLiveResponse{
		WintertodtResult: result,
		PriceInfo:        priceInfo,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
