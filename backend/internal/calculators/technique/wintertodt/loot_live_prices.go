package wintertodt

import (
	"math/rand"
)

// SimulateLootWithLivePricesAndSeed simulates Wintertodt loot with live prices and specific seed
func SimulateLootWithLivePricesAndSeed(rounds int, pointsPerRound int, skillLevels SkillLevels, livePrices map[string]int, seed int64) (map[string]any, int) {
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
					// Use live price if available, otherwise use default
					price := unique.Value
					if livePrice, exists := livePrices[unique.Name]; exists {
						price = livePrice
					}
					totalValue += unique.Quantity * price
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
		// Use live price for magic seeds
		magicSeedPrice := 104000 // Default price
		if livePrice, exists := livePrices["Magic seeds"]; exists {
			magicSeedPrice = livePrice
		}
		totalValue += magicSeedsToAdd * magicSeedPrice
	}

	// If 3+ bruma torches rolled, convert some to torstol seeds
	if brumaTorchCount >= 3 {
		torstolSeedsToAdd := brumaTorchCount / 3
		lootCounts["Torstol seeds"] += torstolSeedsToAdd
		// Use live price for torstol seeds
		torstolSeedPrice := 58000 // Default price
		if livePrice, exists := livePrices["Torstol seeds"]; exists {
			torstolSeedPrice = livePrice
		}
		totalValue += torstolSeedsToAdd * torstolSeedPrice
	}

	// Enhanced supply drops based on skill levels with live prices
	enhancedSupplyDrops := getEnhancedSupplyDropsWithLivePrices(skillLevels, livePrices)

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

// getEnhancedSupplyDropsWithLivePrices returns supply drops with live prices
func getEnhancedSupplyDropsWithLivePrices(skillLevels SkillLevels, livePrices map[string]int) []LootItem {
	enhanced := make([]LootItem, len(SupplyDrops))
	copy(enhanced, SupplyDrops)

	for i, item := range enhanced {
		// Update with live prices first
		if livePrice, exists := livePrices[item.Name]; exists {
			enhanced[i].Value = livePrice
		}

		// Then enhance based on relevant skill levels
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
