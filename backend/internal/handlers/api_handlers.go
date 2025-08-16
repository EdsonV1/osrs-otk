package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"osrs-xp-kits/internal/calculators/technique/wintertodt"
	"osrs-xp-kits/internal/services"
)

// APIHandlers handles external API endpoints
type APIHandlers struct {
	osrsAPI      *services.OSRSAPIService
	cacheManager *services.CacheManager
}

// NewAPIHandlers creates a new API handlers instance
func NewAPIHandlers(osrsAPI *services.OSRSAPIService, cacheManager *services.CacheManager) *APIHandlers {
	return &APIHandlers{
		osrsAPI:      osrsAPI,
		cacheManager: cacheManager,
	}
}

// PlayerStatsRequest represents the request for player stats
type PlayerStatsRequest struct {
	Username string `json:"username"`
}

// PlayerStatsResponse represents the response for player stats
type PlayerStatsResponse struct {
	Success bool                  `json:"success"`
	Data    *services.PlayerStats `json:"data,omitempty"`
	Error   string                `json:"error,omitempty"`
}

// PricesResponse represents the response for current prices
type PricesResponse struct {
	Success     bool           `json:"success"`
	Data        map[string]int `json:"data,omitempty"`
	CacheStatus map[string]any `json:"cache_status,omitempty"`
	Error       string         `json:"error,omitempty"`
}

// CacheStatusResponse represents the response for cache status
type CacheStatusResponse struct {
	Success bool           `json:"success"`
	Data    map[string]any `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
}

// GetPlayerStats handles GET /api/player-stats/{username}
func (h *APIHandlers) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extract username from URL path (Go HTTP already handles URL decoding)
	path := strings.TrimPrefix(r.URL.Path, "/api/player-stats/")
	username := strings.TrimSpace(path)

	// Check for force refresh parameter
	forceRefresh := r.URL.Query().Get("refresh") == "true"

	if username == "" {
		response := PlayerStatsResponse{
			Success: false,
			Error:   "Username is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate username (basic validation)
	if len(username) > 12 || len(username) < 1 {
		response := PlayerStatsResponse{
			Success: false,
			Error:   "Username must be between 1 and 12 characters",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Force refresh if requested
	if forceRefresh {
		err := h.osrsAPI.RefreshPlayerStats(username)
		if err != nil {
			response := PlayerStatsResponse{
				Success: false,
				Error:   fmt.Sprintf("Failed to refresh player stats: %v", err),
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Fetch player stats
	stats, err := h.osrsAPI.GetPlayerStats(username)
	if err != nil {
		response := PlayerStatsResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to fetch player stats: %v", err),
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := PlayerStatsResponse{
		Success: true,
		Data:    stats,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetCurrentPrices handles GET /api/prices
func (h *APIHandlers) GetCurrentPrices(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Get prices from cache manager (handles caching automatically)
	prices, err := h.cacheManager.GetPrices()
	if err != nil {
		response := PricesResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to fetch prices: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get cache status for response
	cacheStatus := h.cacheManager.GetCacheStatus()

	response := PricesResponse{
		Success:     true,
		Data:        prices,
		CacheStatus: cacheStatus,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// RefreshPrices handles POST /api/prices/refresh
func (h *APIHandlers) RefreshPrices(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
		return
	}

	// Force refresh the cache
	err := h.cacheManager.ForceRefresh()
	if err != nil {
		response := PricesResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to refresh prices: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get updated prices and cache status
	prices, _ := h.cacheManager.GetPrices()
	cacheStatus := h.cacheManager.GetCacheStatus()

	response := PricesResponse{
		Success:     true,
		Data:        prices,
		CacheStatus: cacheStatus,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetCacheStatus handles GET /api/cache-status
func (h *APIHandlers) GetCacheStatus(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cacheStatus := h.cacheManager.GetCacheStatus()

	response := CacheStatusResponse{
		Success: true,
		Data:    cacheStatus,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ConvertPlayerStatsToSkillLevels converts PlayerStats to Wintertodt SkillLevels
func ConvertPlayerStatsToSkillLevels(stats *services.PlayerStats) wintertodt.SkillLevels {
	return wintertodt.SkillLevels{
		Herblore:    stats.Herblore,
		Mining:      stats.Mining,
		Fishing:     stats.Fishing,
		Crafting:    stats.Crafting,
		Farming:     stats.Farming,
		Woodcutting: stats.Woodcutting,
	}
}
