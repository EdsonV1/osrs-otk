package handlers

import (
	"encoding/json"
	"net/http"
	"osrs-xp-kits/internal/calculators/technique/birdhouses"
	"osrs-xp-kits/internal/services"
)

// BirdhouseLiveHandler handles birdhouse calculations with live price support
type BirdhouseLiveHandler struct {
	cacheManager *services.CacheManager
}

// NewBirdhouseLiveHandler creates a new birdhouse handler with live price support
func NewBirdhouseLiveHandler(cacheManager *services.CacheManager) *BirdhouseLiveHandler {
	return &BirdhouseLiveHandler{
		cacheManager: cacheManager,
	}
}

// BirdhouseLiveInput extends the basic input with live price options
type BirdhouseLiveInput struct {
	Type          string `json:"type"`
	Quantity      int    `json:"quantity"`
	UseLivePrices bool   `json:"use_live_prices,omitempty"`
}

// BirdhouseLiveResponse extends the basic response with price information
type BirdhouseLiveResponse struct {
	birdhouses.BirdhouseResult
	PriceInfo *PriceInfo `json:"price_info,omitempty"`
}

// Calculate handles POST /api/birdhouse/live
func (h *BirdhouseLiveHandler) Calculate(w http.ResponseWriter, r *http.Request) {
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

	var input BirdhouseLiveInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
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

	// Calculate birdhouse data
	result, err := birdhouses.CalculateBirdhouseDataWithPrices(
		input.Type,
		input.Quantity,
		livePrices,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create enhanced response
	response := BirdhouseLiveResponse{
		BirdhouseResult: result,
		PriceInfo:       priceInfo,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}