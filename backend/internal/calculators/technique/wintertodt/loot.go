package wintertodt

import (
	"math"
	"math/rand"
	"time"
)

// SimulateLoot simulates Wintertodt loot for the given number of rounds
func SimulateLoot(rounds int) (map[string]any, int) {
	return SimulateLootWithSeed(rounds, time.Now().UnixNano())
}

// SimulateLootWithSeed simulates Wintertodt loot with a specific seed for deterministic testing
func SimulateLootWithSeed(rounds int, seed int64) (map[string]any, int) {
	r := rand.New(rand.NewSource(seed))

	lootCounts := make(map[string]int)
	totalValue := 0

	// Always get supply crate per round
	lootCounts["Supply crate"] = rounds

	// Simulate unique rolls per round
	for range rounds {
		for _, item := range UniqueRolls {
			if r.Float64() < item.Rate {
				lootCounts[item.Name] += item.Quantity
				if item.Name != "Phoenix" && item.Name != "Pyromancer outfit" {
					totalValue += item.Quantity * item.Value
				}
			}
		}
	}

	// Simulate supply drops (average 3 rolls per round)
	for range rounds {
		rolls := 3 // Simplified for legacy function

		for range rolls {
			for _, item := range SupplyDrops {
				if r.Float64() < item.Rate {
					lootCounts[item.Name] += item.Quantity
					totalValue += item.Quantity * item.Value
					break // Only one item per roll
				}
			}
		}
	}

	// Format loot for response
	loot := make(map[string]any)
	for name, quantity := range lootCounts {
		if quantity > 0 {
			loot[name] = quantity
		}
	}

	return loot, totalValue
}

// SimulateLootWithSkills simulates Wintertodt loot taking skill levels into account for drop quality
func SimulateLootWithSkills(rounds int, skillLevels SkillLevels) (map[string]any, int) {
	return SimulateLootWithSkillsAndSeed(rounds, skillLevels, time.Now().UnixNano())
}

// SimulateLootWithSkillsAndPoints simulates Wintertodt loot with points and skill levels
func SimulateLootWithSkillsAndPoints(rounds int, pointsPerRound int, skillLevels SkillLevels) (map[string]any, int) {
	return SimulateLootWithSkillsAndPointsAndSeed(rounds, pointsPerRound, skillLevels, time.Now().UnixNano())
}

// SimulateLootWithLivePrices simulates Wintertodt loot with live price data
func SimulateLootWithLivePrices(rounds int, pointsPerRound int, skillLevels SkillLevels, livePrices map[string]int) (map[string]any, int) {
	return SimulateLootWithLivePricesAndSeed(rounds, pointsPerRound, skillLevels, livePrices, time.Now().UnixNano())
}

// CalculateRolls calculates the number of reward rolls based on points
func CalculateRolls(points int) int {
	if points < 500 {
		return 0 // No rewards if less than 500 points
	}

	baseRolls := 2 // Base 2 rolls at 500 points
	extraPoints := points - 500

	// Every 5 points past 500 gives 1% chance of extra roll
	// For simplicity, we'll use the expected value calculation
	extraRollChance := float64(extraPoints) / 5.0 / 100.0 // Convert to percentage
	expectedExtraRolls := extraRollChance

	// Use expected value for deterministic calculation
	totalRolls := baseRolls + int(math.Floor(expectedExtraRolls))

	// Add probabilistic extra roll based on remaining chance
	fractionalChance := expectedExtraRolls - math.Floor(expectedExtraRolls)
	if rand.Float64() < fractionalChance {
		totalRolls++
	}

	return totalRolls
}

// SimulateLootWithSkillsAndSeed simulates Wintertodt loot with proper reward cart mechanics
func SimulateLootWithSkillsAndSeed(rounds int, skillLevels SkillLevels, seed int64) (map[string]any, int) {
	r := rand.New(rand.NewSource(seed))

	lootCounts := make(map[string]int)
	totalValue := 0

	// Track special conditions for unique mechanics
	warmGlovesCount := 0
	brumaTorchCount := 0

	// Always get supply crate per round
	lootCounts["Supply crate"] = rounds

	for range rounds {
		// Each round gets unique rolls (done once per round regardless of points)
		for _, unique := range UniqueRolls {
			if r.Float64() < unique.Rate {
				lootCounts[unique.Name] += unique.Quantity
				if unique.Name != "Phoenix" && unique.Name != "Pyromancer outfit" {
					totalValue += unique.Quantity * unique.Value
				}

				// Track special items for conversion mechanics
				if unique.Name == "Warm gloves" {
					warmGlovesCount++
				}
				if unique.Name == "Bruma torch" {
					brumaTorchCount++
				}
			}
		}
	}

	// Apply special conversion mechanics
	// If 3+ warm gloves rolled, convert some to magic seeds
	if warmGlovesCount >= 3 {
		magicSeedsToAdd := warmGlovesCount / 3
		lootCounts["Magic seeds"] += magicSeedsToAdd
		totalValue += magicSeedsToAdd * 104000 // Magic seed value
	}

	// If 3+ bruma torches rolled, convert some to torstol seeds
	if brumaTorchCount >= 3 {
		torstolSeedsToAdd := brumaTorchCount / 3
		lootCounts["Torstol seeds"] += torstolSeedsToAdd
		totalValue += torstolSeedsToAdd * 58000 // Torstol seed value
	}

	// Enhanced supply drops based on skill levels
	enhancedSupplyDrops := getEnhancedSupplyDrops(skillLevels)

	// Simulate reward rolls for each round (based on points)
	for range rounds {
		// Use provided points per round for accurate roll calculation
		rolls := CalculateRolls(600) // Default baseline for compatibility

		for range rolls {
			for _, item := range enhancedSupplyDrops {
				if r.Float64() < item.Rate {
					lootCounts[item.Name] += item.Quantity
					totalValue += item.Quantity * item.Value
					break // Only one item per roll
				}
			}
		}
	}

	// Format loot for response
	loot := make(map[string]any)
	for name, quantity := range lootCounts {
		if quantity > 0 {
			loot[name] = quantity
		}
	}

	return loot, totalValue
}

// getEnhancedSupplyDrops returns supply drops with quantities and rates modified by skill levels
func getEnhancedSupplyDrops(skillLevels SkillLevels) []LootItem {
	enhanced := make([]LootItem, len(SupplyDrops))
	copy(enhanced, SupplyDrops)

	for i, item := range enhanced {
		// Enhance based on relevant skill levels
		switch item.Name {
		case "Grimy ranarr weed", "Grimy snapdragon", "Grimy torstol":
			// Herblore affects herb quantity and quality
			if skillLevels.Herblore >= 70 {
				enhanced[i].Quantity += 1
			}
			if skillLevels.Herblore >= 90 {
				enhanced[i].Rate *= 1.2
			}
		case "Uncut diamond":
			// Mining affects gem drops
			if skillLevels.Mining >= 75 {
				enhanced[i].Rate *= 1.3
			}
		case "Raw shark":
			// Fishing affects fish quality and quantity
			if skillLevels.Fishing >= 76 {
				enhanced[i].Quantity += 1
			}
		case "Yew logs", "Magic logs":
			// Woodcutting affects log drops
			if skillLevels.Woodcutting >= 75 {
				enhanced[i].Quantity += 2
			}
			if skillLevels.Woodcutting >= 90 {
				enhanced[i].Rate *= 1.1
			}
		case "Pure essence":
			// Mining affects essence quantity
			if skillLevels.Mining >= 60 {
				enhanced[i].Quantity += 25
			}
		}
	}

	return enhanced
}

// SimulateLootWithSkillsAndPointsAndSeed simulates Wintertodt loot with proper reward cart mechanics and points
func SimulateLootWithSkillsAndPointsAndSeed(rounds int, pointsPerRound int, skillLevels SkillLevels, seed int64) (map[string]any, int) {
	r := rand.New(rand.NewSource(seed))

	lootCounts := make(map[string]int)
	totalValue := 0

	// Track special conditions for unique mechanics
	warmGlovesCount := 0
	brumaTorchCount := 0

	// Always get supply crate per round
	lootCounts["Supply crate"] = rounds

	for range rounds {
		// Each round gets unique rolls (done once per round regardless of points)
		for _, unique := range UniqueRolls {
			if r.Float64() < unique.Rate {
				lootCounts[unique.Name] += unique.Quantity
				if unique.Name != "Phoenix" && unique.Name != "Pyromancer outfit" {
					totalValue += unique.Quantity * unique.Value
				}

				// Track special items for conversion mechanics
				if unique.Name == "Warm gloves" {
					warmGlovesCount++
				}
				if unique.Name == "Bruma torch" {
					brumaTorchCount++
				}
			}
		}
	}

	// Apply special conversion mechanics
	// If 3+ warm gloves rolled, convert some to magic seeds
	if warmGlovesCount >= 3 {
		magicSeedsToAdd := warmGlovesCount / 3
		lootCounts["Magic seeds"] += magicSeedsToAdd
		totalValue += magicSeedsToAdd * 104000 // Magic seed value
	}

	// If 3+ bruma torches rolled, convert some to torstol seeds
	if brumaTorchCount >= 3 {
		torstolSeedsToAdd := brumaTorchCount / 3
		lootCounts["Torstol seeds"] += torstolSeedsToAdd
		totalValue += torstolSeedsToAdd * 58000 // Torstol seed value
	}

	// Enhanced supply drops based on skill levels
	enhancedSupplyDrops := getEnhancedSupplyDrops(skillLevels)

	// Simulate reward rolls for each round (based on actual points)
	for range rounds {
		// Use actual points per round for accurate roll calculation
		rolls := CalculateRolls(pointsPerRound)

		for range rolls {
			for _, item := range enhancedSupplyDrops {
				if r.Float64() < item.Rate {
					lootCounts[item.Name] += item.Quantity
					totalValue += item.Quantity * item.Value
					break // Only one item per roll
				}
			}
		}
	}

	// Format loot for response
	loot := make(map[string]any)
	for name, quantity := range lootCounts {
		if quantity > 0 {
			loot[name] = quantity
		}
	}

	return loot, totalValue
}
