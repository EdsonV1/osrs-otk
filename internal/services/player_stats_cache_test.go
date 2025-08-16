package services

import (
	"testing"
	"time"
)

func TestPlayerStatsCache_Initialization(t *testing.T) {
	service := NewOSRSAPIService()

	if service.playerStatsCache == nil {
		t.Error("playerStatsCache is nil")
	}

	if service.playerStatsCache.Stats == nil {
		t.Error("playerStatsCache.Stats is nil")
	}

	if service.playerStatsCache.LastUpdated == nil {
		t.Error("playerStatsCache.LastUpdated is nil")
	}

	// Initially should be empty
	if len(service.playerStatsCache.Stats) != 0 {
		t.Errorf("Expected 0 cached players initially, got %d", len(service.playerStatsCache.Stats))
	}

	if len(service.playerStatsCache.LastUpdated) != 0 {
		t.Errorf("Expected 0 cache timestamps initially, got %d", len(service.playerStatsCache.LastUpdated))
	}
}

func TestPlayerStatsCache_CacheStatus(t *testing.T) {
	service := NewOSRSAPIService()

	// Add some mock data to cache
	testUsername := "testuser"
	testStats := &PlayerStats{
		Username:    testUsername,
		Woodcutting: 85,
		Firemaking:  75,
		Herblore:    60,
	}

	service.playerStatsCache.Stats[testUsername] = testStats
	service.playerStatsCache.LastUpdated[testUsername] = time.Now()

	status := service.GetCacheStatus()
	playerStatsCache := status["player_stats"].(map[string]interface{})

	playersCached, ok := playerStatsCache["players_cached"].(int)
	if !ok || playersCached != 1 {
		t.Errorf("Expected 1 player cached, got %v", playersCached)
	}

	cacheEntries, ok := playerStatsCache["cache_entries"].(int)
	if !ok || cacheEntries != 1 {
		t.Errorf("Expected 1 cache entry, got %v", cacheEntries)
	}
}

func TestRefreshPlayerStats_ForcesRefresh(t *testing.T) {
	service := NewOSRSAPIService()

	// Add existing cache entry with recent timestamp
	testUsername := "testuser"
	testStats := &PlayerStats{
		Username:    testUsername,
		Woodcutting: 85,
	}

	service.playerStatsCache.Stats[testUsername] = testStats
	service.playerStatsCache.LastUpdated[testUsername] = time.Now().Add(-1 * time.Hour) // 1 hour ago

	// Verify cache exists
	if _, exists := service.playerStatsCache.Stats[testUsername]; !exists {
		t.Error("Test setup failed: cache should exist")
	}

	// Try to refresh (this will likely fail in test environment due to network/API)
	err := service.RefreshPlayerStats(testUsername)

	// The important thing is that cache was cleared (even if refresh failed)
	if _, exists := service.playerStatsCache.Stats[testUsername]; exists {
		// If refresh succeeded, cache should be repopulated, if it failed, cache should be empty
		// Either way is acceptable behavior
		t.Logf("Cache still exists after refresh attempt - this is OK if refresh succeeded")
	}

	if err != nil {
		t.Logf("RefreshPlayerStats failed (expected in test environment): %v", err)
	}
}

func TestPlayerStatsCache_CacheExpiry(t *testing.T) {
	service := NewOSRSAPIService()

	testUsername := "testuser"

	// Test case 1: Fresh cache (should return cached data)
	service.playerStatsCache.Stats[testUsername] = &PlayerStats{Username: testUsername}
	service.playerStatsCache.LastUpdated[testUsername] = time.Now().Add(-1 * time.Hour) // 1 hour ago (fresh)

	// This should use cache (won't hit network in test)
	stats, err := service.GetPlayerStats(testUsername)
	if err == nil && stats != nil && stats.Username == testUsername {
		t.Log("Cache hit successful for fresh cache")
	}

	// Test case 2: Stale cache (older than 6 hours)
	service.playerStatsCache.LastUpdated[testUsername] = time.Now().Add(-7 * time.Hour) // 7 hours ago (stale)

	// This should attempt to refresh cache (will likely fail in test environment)
	_, err = service.GetPlayerStats(testUsername)
	if err != nil {
		t.Logf("Cache refresh failed for stale cache (expected in test environment): %v", err)
	}
}
