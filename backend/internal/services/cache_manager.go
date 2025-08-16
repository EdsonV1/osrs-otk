package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// CacheManager handles persistent caching of API data
type CacheManager struct {
	cacheDir      string
	osrsAPI       *OSRSAPIService
	mu            sync.RWMutex
	refreshTicker *time.Ticker
	stopChan      chan struct{}
}

// CacheData represents the structure of cached data
type CacheData struct {
	Prices      map[string]int `json:"prices"`
	LastUpdated time.Time      `json:"last_updated"`
	Version     string         `json:"version"`
}

const (
	CacheVersion     = "1.0"
	PriceCacheFile   = "price_cache.json"
	CacheRefreshHour = 6 // Refresh at 6 AM daily
)

// NewCacheManager creates a new cache manager
func NewCacheManager(cacheDir string, osrsAPI *OSRSAPIService) *CacheManager {
	cm := &CacheManager{
		cacheDir: cacheDir,
		osrsAPI:  osrsAPI,
		stopChan: make(chan struct{}),
	}

	// Ensure cache directory exists
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		fmt.Printf("Warning: Failed to create cache directory: %v\n", err)
	}

	// Load existing cache
	cm.loadPriceCache()

	return cm
}

// StartDailyRefresh starts the daily cache refresh routine
func (cm *CacheManager) StartDailyRefresh() {
	// Calculate time until next refresh (6 AM)
	now := time.Now()
	nextRefresh := time.Date(now.Year(), now.Month(), now.Day(), CacheRefreshHour, 0, 0, 0, now.Location())

	// If it's already past 6 AM today, schedule for tomorrow
	if now.After(nextRefresh) {
		nextRefresh = nextRefresh.Add(24 * time.Hour)
	}

	// Start ticker for daily refresh
	duration := nextRefresh.Sub(now)

	// Initial wait until first refresh time
	go func() {
		select {
		case <-time.After(duration):
			cm.refreshPrices()
		case <-cm.stopChan:
			return
		}

		// Then refresh every 24 hours
		cm.refreshTicker = time.NewTicker(24 * time.Hour)
		for {
			select {
			case <-cm.refreshTicker.C:
				cm.refreshPrices()
			case <-cm.stopChan:
				if cm.refreshTicker != nil {
					cm.refreshTicker.Stop()
				}
				return
			}
		}
	}()

	fmt.Printf("Daily price refresh scheduled for %s\n", nextRefresh.Format("2006-01-02 15:04:05"))
}

// Stop stops the daily refresh routine
func (cm *CacheManager) Stop() {
	close(cm.stopChan)
	if cm.refreshTicker != nil {
		cm.refreshTicker.Stop()
	}
}

// refreshPrices refreshes the price cache and saves to disk
func (cm *CacheManager) refreshPrices() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	fmt.Println("Refreshing price cache...")

	if err := cm.osrsAPI.RefreshPrices(); err != nil {
		fmt.Printf("Error refreshing prices: %v\n", err)
		return
	}

	if err := cm.savePriceCache(); err != nil {
		fmt.Printf("Error saving price cache: %v\n", err)
		return
	}

	fmt.Println("Price cache refreshed successfully")
}

// loadPriceCache loads the price cache from disk
func (cm *CacheManager) loadPriceCache() {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cacheFile := filepath.Join(cm.cacheDir, PriceCacheFile)

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Printf("Warning: Failed to read price cache: %v\n", err)
		}
		return
	}

	var cacheData CacheData
	if err := json.Unmarshal(data, &cacheData); err != nil {
		fmt.Printf("Warning: Failed to parse price cache: %v\n", err)
		return
	}

	// Check if cache is valid version and not too old
	if cacheData.Version != CacheVersion {
		fmt.Println("Cache version mismatch, will refresh")
		return
	}

	if time.Since(cacheData.LastUpdated) > 25*time.Hour { // Allow 1 hour buffer
		fmt.Println("Cache is stale, will refresh")
		return
	}

	// Load into OSRS API service
	cm.osrsAPI.priceCache.Prices = cacheData.Prices
	cm.osrsAPI.priceCache.LastUpdated = cacheData.LastUpdated

	fmt.Printf("Loaded price cache with %d items (updated %s)\n",
		len(cacheData.Prices), cacheData.LastUpdated.Format("2006-01-02 15:04:05"))
}

// savePriceCache saves the current price cache to disk
func (cm *CacheManager) savePriceCache() error {
	cacheFile := filepath.Join(cm.cacheDir, PriceCacheFile)

	cacheData := CacheData{
		Prices:      cm.osrsAPI.priceCache.Prices,
		LastUpdated: cm.osrsAPI.priceCache.LastUpdated,
		Version:     CacheVersion,
	}

	data, err := json.MarshalIndent(cacheData, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling cache data: %w", err)
	}

	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		return fmt.Errorf("writing cache file: %w", err)
	}

	return nil
}

// GetPrices returns cached prices, refreshing if necessary
func (cm *CacheManager) GetPrices() (map[string]int, error) {
	cm.mu.RLock()

	// Check if cache needs refresh
	needsRefresh := time.Since(cm.osrsAPI.priceCache.LastUpdated) > 24*time.Hour ||
		len(cm.osrsAPI.priceCache.Prices) == 0

	cm.mu.RUnlock()

	if needsRefresh {
		cm.mu.Lock()
		// Double-check after acquiring write lock
		if time.Since(cm.osrsAPI.priceCache.LastUpdated) > 24*time.Hour ||
			len(cm.osrsAPI.priceCache.Prices) == 0 {

			if err := cm.osrsAPI.RefreshPrices(); err != nil {
				cm.mu.Unlock()
				return nil, fmt.Errorf("refreshing prices: %w", err)
			}

			if err := cm.savePriceCache(); err != nil {
				fmt.Printf("Warning: Failed to save price cache: %v\n", err)
			}
		}
		cm.mu.Unlock()
	}

	return cm.osrsAPI.GetCurrentPrices()
}

// ForceRefresh forces an immediate refresh of all cached data
func (cm *CacheManager) ForceRefresh() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if err := cm.osrsAPI.RefreshPrices(); err != nil {
		return fmt.Errorf("refreshing prices: %w", err)
	}

	if err := cm.savePriceCache(); err != nil {
		return fmt.Errorf("saving cache: %w", err)
	}

	return nil
}

// GetCacheStatus returns detailed cache status information
func (cm *CacheManager) GetCacheStatus() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	status := cm.osrsAPI.GetCacheStatus()

	// Add cache manager specific information
	status["cache_file"] = filepath.Join(cm.cacheDir, PriceCacheFile)
	status["version"] = CacheVersion
	status["refresh_hour"] = CacheRefreshHour

	// Check if cache file exists
	cacheFile := filepath.Join(cm.cacheDir, PriceCacheFile)
	if _, err := os.Stat(cacheFile); err == nil {
		status["cache_file_exists"] = true
		if info, err := os.Stat(cacheFile); err == nil {
			status["cache_file_size"] = info.Size()
			status["cache_file_modified"] = info.ModTime()
		}
	} else {
		status["cache_file_exists"] = false
	}

	return status
}
