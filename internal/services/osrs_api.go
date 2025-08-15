package services

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// OSRSAPIService handles interactions with OSRS external APIs
type OSRSAPIService struct {
	httpClient       *http.Client
	priceCache       *PriceCache
	playerStatsCache *PlayerStatsCache
}

// PriceCache stores cached price data with timestamp
type PriceCache struct {
	Prices      map[string]int `json:"prices"`
	LastUpdated time.Time      `json:"last_updated"`
}

// PlayerStatsCache stores cached player stats with timestamp
type PlayerStatsCache struct {
	Stats       map[string]*PlayerStats `json:"stats"`
	LastUpdated map[string]time.Time    `json:"last_updated"`
}

// WikiPriceResponse represents the OSRS Wiki price API response
type WikiPriceResponse struct {
	Data map[string]WikiPriceData `json:"data"`
}

type WikiPriceData struct {
	High     *int       `json:"high"`
	HighTime *time.Time `json:"highTime"`
	Low      *int       `json:"low"`
	LowTime  *time.Time `json:"lowTime"`
}

// PlayerStats represents a player's skill levels from hiscores
type PlayerStats struct {
	Username     string `json:"username"`
	Overall      int    `json:"overall"`
	Attack       int    `json:"attack"`
	Defence      int    `json:"defence"`
	Strength     int    `json:"strength"`
	Hitpoints    int    `json:"hitpoints"`
	Ranged       int    `json:"ranged"`
	Prayer       int    `json:"prayer"`
	Magic        int    `json:"magic"`
	Cooking      int    `json:"cooking"`
	Woodcutting  int    `json:"woodcutting"`
	Fletching    int    `json:"fletching"`
	Fishing      int    `json:"fishing"`
	Firemaking   int    `json:"firemaking"`
	Crafting     int    `json:"crafting"`
	Smithing     int    `json:"smithing"`
	Mining       int    `json:"mining"`
	Herblore     int    `json:"herblore"`
	Agility      int    `json:"agility"`
	Thieving     int    `json:"thieving"`
	Slayer       int    `json:"slayer"`
	Farming      int    `json:"farming"`
	Runecrafting int    `json:"runecrafting"`
	Hunter       int    `json:"hunter"`
	Construction int    `json:"construction"`
}

// Item ID mappings for OSRS Wiki API
var itemIDMap = map[string]int{
	"Grimy ranarr weed": 207,
	"Grimy snapdragon":  3051,
	"Grimy torstol":     219,
	"Uncut diamond":     1617,
	"Pure essence":      7936,
	"Raw shark":         383,
	"Yew logs":          1515,
	"Magic logs":        1513,
	"Dragon axe":        6739,
	"Tome of fire":      20714,
	"Warm gloves":       10071,
	"Bruma torch":       20730,
	"Burnt page":        20718,
	"Magic seeds":       5316,
	"Torstol seeds":     5304,
}

// NewOSRSAPIService creates a new OSRS API service
func NewOSRSAPIService() *OSRSAPIService {
	return &OSRSAPIService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		priceCache: &PriceCache{
			Prices: make(map[string]int),
		},
		playerStatsCache: &PlayerStatsCache{
			Stats:       make(map[string]*PlayerStats),
			LastUpdated: make(map[string]time.Time),
		},
	}
}

// GetCurrentPrices fetches current item prices from OSRS Wiki API
func (s *OSRSAPIService) GetCurrentPrices() (map[string]int, error) {
	// Check if cache is still valid (less than 24 hours old)
	if time.Since(s.priceCache.LastUpdated) < 24*time.Hour && len(s.priceCache.Prices) > 0 {
		return s.priceCache.Prices, nil
	}

	// Build item ID list for API request
	var itemIDs []string
	for _, id := range itemIDMap {
		itemIDs = append(itemIDs, strconv.Itoa(id))
	}

	url := fmt.Sprintf("https://prices.runescape.wiki/api/v1/osrs/latest?id=%s", strings.Join(itemIDs, ","))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Set descriptive User-Agent header as recommended by OSRS Wiki API
	req.Header.Set("User-Agent", "OSRS-OTK Calculator v1.0")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching prices: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var priceResponse WikiPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return nil, fmt.Errorf("decoding price response: %w", err)
	}

	// Convert response to our format
	prices := make(map[string]int)
	for itemName, itemID := range itemIDMap {
		idStr := strconv.Itoa(itemID)
		if data, exists := priceResponse.Data[idStr]; exists {
			// Use high price if available, otherwise low price
			if data.High != nil {
				prices[itemName] = *data.High
			} else if data.Low != nil {
				prices[itemName] = *data.Low
			}
		}
	}

	// Update cache
	s.priceCache.Prices = prices
	s.priceCache.LastUpdated = time.Now()

	return prices, nil
}

// GetPlayerStats fetches player stats from OSRS hiscores with caching
func (s *OSRSAPIService) GetPlayerStats(username string) (*PlayerStats, error) {
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	// Check cache first (valid for 6 hours)
	if lastUpdated, exists := s.playerStatsCache.LastUpdated[username]; exists {
		if time.Since(lastUpdated) < 6*time.Hour {
			if stats, hasStats := s.playerStatsCache.Stats[username]; hasStats && stats != nil {
				return stats, nil
			}
		}
	}

	url := fmt.Sprintf("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=%s", username)

	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching player stats: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("player '%s' not found on hiscores", username)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hiscores API returned status %d", resp.StatusCode)
	}

	// Parse CSV response with robust error handling
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	reader.TrimLeadingSpace = true

	var records [][]string
	lineNum := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			// Skip malformed lines and continue
			lineNum++
			if lineNum > 50 { // Prevent infinite loop
				break
			}
			continue
		}
		records = append(records, record)
		lineNum++
	}

	if len(records) < 23 { // Need at least 23 skills (some players might not have all skills ranked)
		return nil, fmt.Errorf("invalid hiscores response: only found %d skill records, expected at least 23", len(records))
	}

	stats := &PlayerStats{Username: username}

	// Parse each skill (format: rank,level,xp)
	skills := []*int{
		&stats.Overall, &stats.Attack, &stats.Defence, &stats.Strength, &stats.Hitpoints,
		&stats.Ranged, &stats.Prayer, &stats.Magic, &stats.Cooking,
		&stats.Woodcutting, &stats.Fletching, &stats.Fishing, &stats.Firemaking,
		&stats.Crafting, &stats.Smithing, &stats.Mining, &stats.Herblore,
		&stats.Agility, &stats.Thieving, &stats.Slayer, &stats.Farming,
		&stats.Runecrafting, &stats.Hunter, &stats.Construction,
	}

	for i, skill := range skills {
		if i >= len(records) {
			break
		}

		record := records[i]

		// Skip records that don't have at least 2 fields (rank, level)
		if len(record) < 2 {
			*skill = 1 // Default to level 1
			continue
		}

		// Handle cases where level field might be "-1" or empty
		levelStr := strings.TrimSpace(record[1])
		if levelStr == "" || levelStr == "-1" {
			*skill = 1 // Default to level 1 for unranked skills
			continue
		}

		level, err := strconv.Atoi(levelStr)
		if err != nil || level < 1 {
			// If parsing fails or invalid level, default to level 1
			*skill = 1
		} else {
			*skill = level
		}
	}

	// Cache the stats
	s.playerStatsCache.Stats[username] = stats
	s.playerStatsCache.LastUpdated[username] = time.Now()

	return stats, nil
}

// GetPriceByName returns the current price for a specific item
func (s *OSRSAPIService) GetPriceByName(itemName string) (int, error) {
	prices, err := s.GetCurrentPrices()
	if err != nil {
		return 0, err
	}

	price, exists := prices[itemName]
	if !exists {
		return 0, fmt.Errorf("price not found for item: %s", itemName)
	}

	return price, nil
}

// RefreshPrices forces a refresh of the price cache
func (s *OSRSAPIService) RefreshPrices() error {
	s.priceCache.LastUpdated = time.Time{} // Reset cache timestamp
	_, err := s.GetCurrentPrices()
	return err
}

// RefreshPlayerStats forces a refresh of player stats for a specific username
func (s *OSRSAPIService) RefreshPlayerStats(username string) error {
	// Remove from cache to force refresh
	delete(s.playerStatsCache.Stats, username)
	delete(s.playerStatsCache.LastUpdated, username)
	_, err := s.GetPlayerStats(username)
	return err
}

// GetCacheStatus returns information about the current cache status
func (s *OSRSAPIService) GetCacheStatus() map[string]interface{} {
	return map[string]interface{}{
		"prices": map[string]interface{}{
			"last_updated":    s.priceCache.LastUpdated,
			"items_cached":    len(s.priceCache.Prices),
			"cache_age_hours": time.Since(s.priceCache.LastUpdated).Hours(),
			"is_stale":        time.Since(s.priceCache.LastUpdated) > 24*time.Hour,
		},
		"player_stats": map[string]interface{}{
			"players_cached": len(s.playerStatsCache.Stats),
			"cache_entries":  len(s.playerStatsCache.LastUpdated),
		},
	}
}
