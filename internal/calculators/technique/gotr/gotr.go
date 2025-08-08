package gotr

import (
	"fmt"
	"math"
	"osrs-xp-kits/internal/calculators"
)

// GOTRResult represents the calculated results for GOTR training
type GOTRResult struct {
	CurrentLevel        int      `json:"current_level"`
	TargetLevel         int      `json:"target_level"`
	XPNeeded            int      `json:"xp_needed"`
	GamesNeeded         int      `json:"games_needed"`
	HoursNeeded         float64  `json:"hours_needed"`
	AverageXPPerGame    float64  `json:"average_xp_per_game"`
	AverageXPPerHour    float64  `json:"average_xp_per_hour"`
	TotalRewardRolls    int      `json:"total_reward_rolls"`
	PetChancePercentage float64  `json:"pet_chance_percentage"`
	EstimatedRewards    []Reward `json:"estimated_rewards"`
	TotalRewardValue    int      `json:"total_reward_value"`
	GPPerHour           float64  `json:"gp_per_hour"`
}

// CalculateGOTRData performs the main GOTR calculation
func CalculateGOTRData(currentLevel, targetLevel int) (GOTRResult, error) {
	// Input validation
	if currentLevel < 27 || currentLevel > 126 {
		return GOTRResult{}, fmt.Errorf("current level must be between 27 and 126 (minimum level to access GOTR)")
	}

	if targetLevel < 27 || targetLevel > 126 {
		return GOTRResult{}, fmt.Errorf("target level must be between 27 and 126")
	}

	if targetLevel <= currentLevel {
		return GOTRResult{}, fmt.Errorf("target level must be higher than current level")
	}

	// Calculate XP needed
	xpNeeded, err := calculators.XPRequired(currentLevel, targetLevel)
	if err != nil {
		return GOTRResult{}, fmt.Errorf("error calculating XP required: %v", err)
	}

	// Calculate average XP per game based on runecrafting level
	// GOTR XP scales with RC level, with optimal rates around level 77+
	avgXPPerGame := calculateXPPerGame(currentLevel, targetLevel)
	avgXPPerHour := avgXPPerGame * GamesPerHour

	// Calculate games and time needed
	gamesNeeded := int(math.Ceil(xpNeeded / avgXPPerGame))
	hoursNeeded := float64(gamesNeeded) / GamesPerHour

	// Calculate reward searches (approximately 18 searches per game on average)
	totalSearches := int(float64(gamesNeeded) * AverageSearchesPerGame)

	// Calculate pet chance (Abyssal Protector - 1/4000 per search)
	petChance := 1.0 - math.Pow(1.0-PetRatePerSearch, float64(totalSearches))
	petChancePercentage := petChance * 100

	// Simulate rewards
	rewards, totalValue := SimulateAverageRewards(totalSearches)
	gpPerHour := float64(totalValue) / hoursNeeded

	return GOTRResult{
		CurrentLevel:        currentLevel,
		TargetLevel:         targetLevel,
		XPNeeded:            int(xpNeeded),
		GamesNeeded:         gamesNeeded,
		HoursNeeded:         hoursNeeded,
		AverageXPPerGame:    avgXPPerGame,
		AverageXPPerHour:    avgXPPerHour,
		TotalRewardRolls:    totalSearches,
		PetChancePercentage: petChancePercentage,
		EstimatedRewards:    rewards,
		TotalRewardValue:    totalValue,
		GPPerHour:           gpPerHour,
	}, nil
}

// calculateXPPerGame calculates the average XP per GOTR game based on RC level
func calculateXPPerGame(currentLevel, targetLevel int) float64 {
	// Use the average level for calculation to account for leveling up during training
	avgLevel := float64(currentLevel+targetLevel) / 2.0

	// Base XP calculation - GOTR XP scales significantly with RC level
	// Real GOTR rates: ~180k+ XP/hr at level 77, up to ~220k+ at 99

	var xpPerGame float64

	if avgLevel < 77 {
		// Before level 77, GOTR is less efficient (use ZMI instead)
		// Scale from ~120k/hr to ~160k/hr
		xpPerHourBase := 120000 + (avgLevel-50)*1500
		xpPerGame = xpPerHourBase / GamesPerHour
	} else {
		// At and after level 77, GOTR becomes very efficient
		// Base rate at 77: ~180k XP/hr, scaling to ~220k+ at 99
		xpPerHourBase := 180000 + (avgLevel-77)*1800
		xpPerGame = xpPerHourBase / GamesPerHour

		// Additional scaling for very high levels
		if avgLevel >= 90 {
			xpPerGame *= 1.08 // 8% bonus for mastery
		}

		if avgLevel >= 95 {
			xpPerGame *= 1.05 // Additional 5% for near-max efficiency
		}
	}

	return xpPerGame
}

// CalculateOptimalStrategy suggests the best approach based on current level
func CalculateOptimalStrategy(currentLevel int) string {
	switch {
	case currentLevel < 50:
		return "Early GOTR access - consider mixing with other RC methods for better XP rates until level 50+"
	case currentLevel < 77:
		return "Good GOTR training range - consider training to level 77 for maximum efficiency, or continue with GOTR for convenience"
	case currentLevel < 85:
		return "Optimal GOTR efficiency reached - focus on consistent games with good portal management"
	case currentLevel < 95:
		return "High efficiency phase - maximize searches per game and consider RuneCrafting outfit for bonus XP"
	default:
		return "Maximum efficiency - perfect for final push to 99 with excellent profit potential"
	}
}

// EstimateTimeToLevel provides a more detailed breakdown of time estimates
func EstimateTimeToLevel(currentLevel, targetLevel int) (map[string]any, error) {
	result, err := CalculateGOTRData(currentLevel, targetLevel)
	if err != nil {
		return nil, err
	}

	strategy := CalculateOptimalStrategy(currentLevel)

	breakdown := map[string]any{
		"total_hours":      result.HoursNeeded,
		"total_games":      result.GamesNeeded,
		"daily_hours_1h":   math.Ceil(result.HoursNeeded),
		"daily_hours_2h":   math.Ceil(result.HoursNeeded / 2),
		"daily_hours_3h":   math.Ceil(result.HoursNeeded / 3),
		"optimal_strategy": strategy,
		"xp_per_hour":      result.AverageXPPerHour,
		"profit_potential": result.GPPerHour,
		"pet_chance":       result.PetChancePercentage,
	}

	return breakdown, nil
}
