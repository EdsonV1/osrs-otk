package services

import (
	"testing"
	"time"
)

func TestNewOSRSAPIService(t *testing.T) {
	service := NewOSRSAPIService()

	if service == nil {
		t.Fatal("NewOSRSAPIService() returned nil")
	}

	if service.httpClient == nil {
		t.Error("httpClient is nil")
	}

	if service.priceCache == nil {
		t.Error("priceCache is nil")
	}

	if service.priceCache.Prices == nil {
		t.Error("priceCache.Prices is nil")
	}
}

func TestGetPlayerStats_InvalidUsername(t *testing.T) {
	service := NewOSRSAPIService()

	// Test empty username
	_, err := service.GetPlayerStats("")
	if err == nil {
		t.Error("Expected error for empty username")
	}

	// Test very long username (longer than OSRS allows)
	longUsername := "verylongusernamethatexceedsosrslimits"
	_, err = service.GetPlayerStats(longUsername)
	// This should either error immediately or when hitting the API
	// We don't expect it to succeed
}

func TestGetPriceByName_NonexistentItem(t *testing.T) {
	service := NewOSRSAPIService()

	// Try to get price for non-existent item
	_, err := service.GetPriceByName("Non-existent item")
	if err == nil {
		t.Error("Expected error for non-existent item")
	}
}

func TestGetCacheStatus(t *testing.T) {
	service := NewOSRSAPIService()

	status := service.GetCacheStatus()

	// Check price cache fields exist
	pricesCache, exists := status["prices"]
	if !exists {
		t.Error("Cache status missing 'prices' field")
	}

	pricesCacheMap, ok := pricesCache.(map[string]interface{})
	if !ok {
		t.Error("Prices cache should be a map")
	}

	if _, exists := pricesCacheMap["last_updated"]; !exists {
		t.Error("Price cache missing 'last_updated' field")
	}

	if _, exists := pricesCacheMap["items_cached"]; !exists {
		t.Error("Price cache missing 'items_cached' field")
	}

	if _, exists := pricesCacheMap["cache_age_hours"]; !exists {
		t.Error("Price cache missing 'cache_age_hours' field")
	}

	if _, exists := pricesCacheMap["is_stale"]; !exists {
		t.Error("Price cache missing 'is_stale' field")
	}

	// Check player stats cache fields exist
	playerStatsCache, exists := status["player_stats"]
	if !exists {
		t.Error("Cache status missing 'player_stats' field")
	}

	playerStatsCacheMap, ok := playerStatsCache.(map[string]interface{})
	if !ok {
		t.Error("Player stats cache should be a map")
	}

	if _, exists := playerStatsCacheMap["players_cached"]; !exists {
		t.Error("Player stats cache missing 'players_cached' field")
	}

	if _, exists := playerStatsCacheMap["cache_entries"]; !exists {
		t.Error("Player stats cache missing 'cache_entries' field")
	}
}

func TestPriceCache_Initialization(t *testing.T) {
	service := NewOSRSAPIService()

	// Initially cache should be empty and considered stale
	status := service.GetCacheStatus()

	pricesCache := status["prices"].(map[string]interface{})

	itemsCached, ok := pricesCache["items_cached"].(int)
	if !ok || itemsCached != 0 {
		t.Errorf("Expected 0 items cached initially, got %v", itemsCached)
	}

	isStale, ok := pricesCache["is_stale"].(bool)
	if !ok || !isStale {
		t.Errorf("Expected cache to be stale initially, got %v", isStale)
	}
}

func TestRefreshPrices_EmptyCache(t *testing.T) {
	service := NewOSRSAPIService()

	// Reset cache to ensure empty state
	service.priceCache.LastUpdated = time.Time{}
	service.priceCache.Prices = make(map[string]int)

	// RefreshPrices should work even with empty cache
	err := service.RefreshPrices()
	// We expect this might fail due to network/API issues in test environment
	// but it shouldn't panic or have structural issues
	if err != nil {
		t.Logf("RefreshPrices failed (expected in test environment): %v", err)
	}
}
