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
	totalSearches := int(float64(gamesNeeded) * AverageRewardSearches)

	// Calculate pet chance (Note: Rift guardian cannot be obtained during minigame per wiki)
	// Using generic pet rate for any other possible pets
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
// Based on real player data: 20k-50k XP/hour depending on level
func calculateXPPerGame(currentLevel, targetLevel int) float64 {
	// Use the average level for calculation to account for leveling up during training
	avgLevel := float64(currentLevel+targetLevel) / 2.0

	// Calculate XP per hour based on realistic rates from player guides
	var xpPerHour float64

	// Linear interpolation based on known data points:
	// Level 27: 20,000 XP/hr
	// Level 80: 45,000 XP/hr
	// Level 90: 50,000 XP/hr

	if avgLevel <= 27 {
		xpPerHour = 20000
	} else if avgLevel <= 80 {
		// Linear interpolation between level 27 (20k) and level 80 (45k)
		// Rate increases by 25k over 53 levels = ~471 XP/hr per level
		xpPerHour = 20000 + (avgLevel-27)*471.7
	} else if avgLevel <= 90 {
		// Linear interpolation between level 80 (45k) and level 90 (48k)
		// Rate increases by 3k over 10 levels = 300 XP/hr per level
		xpPerHour = 45000 + (avgLevel-80)*300
	} else {
		// Cap at level 90+ rates
		xpPerHour = 48000
	}

	// Convert XP per hour to XP per game
	// Using 6 games per hour (10 minute games)
	xpPerGame := xpPerHour / GamesPerHour

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

// GetCalculationProTips provides detailed information about how GOTR calculations work
func GetCalculationProTips() map[string]any {
	return map[string]any{
		"calculation_methodology": map[string]any{
			"xp_rates_source": "Based on real player data from community guides",
			"base_formula":    "Linear interpolation between known XP/hour data points",
			"data_points": []map[string]any{
				{"level": 27, "xp_per_hour": 20000, "note": "Minimum access level"},
				{"level": 50, "xp_per_hour": 30000, "note": "Mid-level efficiency"},
				{"level": 80, "xp_per_hour": 45000, "note": "High efficiency threshold"},
				{"level": 90, "xp_per_hour": 48000, "note": "Near-maximum rates"},
			},
		},
		"game_mechanics": map[string]any{
			"game_duration":  "10 minutes per game (experienced players)",
			"games_per_hour": 6,
			"xp_sources": []string{
				"Essence crafting (primary source)",
				"Barrier repairs",
				"Guardian creation",
				"Game completion bonus",
			},
		},
		"factors_considered": []string{
			"Runecrafting level (major impact on XP rates)",
			"Player efficiency and experience",
			"Portal availability and management",
			"Consistent gameplay over time",
			"Average performance across multiple games",
		},
		"accuracy_notes": map[string]any{
			"rates_vary": "Individual XP rates can vary ±10% based on:",
			"variance_factors": []string{
				"Number of portals available",
				"Team coordination in mass worlds",
				"Personal skill and familiarity",
				"RNG in portal spawns and timing",
			},
			"calculation_basis": "Rates assume consistent, moderately efficient gameplay",
		},
		"pro_tips": []map[string]string{
			{
				"tip":         "Level Progression",
				"description": "GOTR becomes significantly more efficient after level 77-80",
			},
			{
				"tip":         "Game Strategy",
				"description": "Focus on portal management and essence mining for maximum XP",
			},
			{
				"tip":         "Time Investment",
				"description": "GOTR is moderate XP but offers excellent rewards - factor in GP/hour value",
			},
			{
				"tip":         "Realistic Expectations",
				"description": "These rates assume consistent play - expect some variation in practice",
			},
			{
				"tip":         "Alternative Methods",
				"description": "Consider ZMI or Lavas for pure XP, GOTR for balanced XP + profit",
			},
		},
		"reward_calculation": map[string]any{
			"searches_per_game": "18 reward searches on average (good performance)",
			"search_variation":  "12-24 searches depending on efficiency and RNG",
			"loot_simulation":   "Based on drop table probabilities and average quantities",
			"gp_per_hour":       "Calculated from total reward value divided by training time",
		},
	}
}
