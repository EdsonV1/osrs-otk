package gotr

import (
	"math"
	"math/rand"
	"time"
)

// LootSimulator handles GOTR reward calculations
type LootSimulator struct {
	rand *rand.Rand
}

// NewLootSimulator creates a new loot simulator
func NewLootSimulator() *LootSimulator {
	return &LootSimulator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SimulateRewards calculates expected rewards based on the number of searches
func (ls *LootSimulator) SimulateRewards(totalSearches int) ([]Reward, int) {
	rewards := make(map[string]*Reward)
	totalValue := 0

	// Calculate total weight for drop rate calculations
	totalWeight := 0
	for _, item := range RewardTable {
		totalWeight += item.Weight
	}

	// Simulate each search
	for search := 0; search < totalSearches; search++ {
		// Determine which item(s) are rewarded this search
		// In GOTR, you typically get multiple items per search
		itemsThisSearch := ls.getItemsForSearch(totalWeight)

		for _, itemKey := range itemsThisSearch {
			item := RewardTable[itemKey]
			quantity := ls.calculateQuantity(item)

			if existing, exists := rewards[item.Name]; exists {
				existing.Quantity += quantity
			} else {
				rewards[item.Name] = &Reward{
					Name:     item.Name,
					Quantity: quantity,
					Value:    item.Value,
					DropRate: ls.calculateDropRate(item.Weight, totalWeight),
				}
			}

			totalValue += quantity * item.Value
		}
	}

	// Convert map to slice
	rewardSlice := make([]Reward, 0, len(rewards))
	for _, reward := range rewards {
		rewardSlice = append(rewardSlice, *reward)
	}

	return rewardSlice, totalValue
}

// getItemsForSearch determines which items are rewarded in a single search
func (ls *LootSimulator) getItemsForSearch(totalWeight int) []string {
	items := make([]string, 0, 3) // typically 2-3 items per search

	// Guaranteed common drops (essence and stones)
	commonItems := []string{"guardian_essence", "catalytic_guardian_stone", "elemental_guardian_stone"}
	for _, item := range commonItems {
		if ls.rand.Float64() < 0.8 { // 80% chance for each common item
			items = append(items, item)
		}
	}

	// Chance for rune rewards
	runeItems := []string{"nature_rune", "death_rune", "blood_rune", "soul_rune"}
	for _, item := range runeItems {
		itemData := RewardTable[item]
		chance := float64(itemData.Weight) / float64(totalWeight) * 2.5 // boost rune drop rates
		if ls.rand.Float64() < chance {
			items = append(items, item)
		}
	}

	// Rare drops (very low chance)
	rareItems := []string{"abyssal_needle", "abyssal_lantern", "raiments_of_the_eye_top", "raiments_of_the_eye_bottom", "hat_of_the_eye"}
	for _, item := range rareItems {
		itemData := RewardTable[item]
		chance := float64(itemData.Weight) / float64(totalWeight) * 0.1 // very rare
		if ls.rand.Float64() < chance {
			items = append(items, item)
		}
	}

	// Special items
	specialItems := []string{"intrinsic_catalyst", "lantern_lens"}
	for _, item := range specialItems {
		itemData := RewardTable[item]
		chance := float64(itemData.Weight) / float64(totalWeight) * 1.2
		if ls.rand.Float64() < chance {
			items = append(items, item)
		}
	}

	// Ensure at least one item is always dropped
	if len(items) == 0 {
		items = append(items, "guardian_essence")
	}

	return items
}

// calculateQuantity determines the quantity of an item based on its variance
func (ls *LootSimulator) calculateQuantity(item RewardItem) int {
	if item.VarianceMin == item.VarianceMax {
		return item.BaseQuantity
	}

	// Random quantity within the variance range
	variance := item.VarianceMax - item.VarianceMin
	return item.VarianceMin + ls.rand.Intn(variance+1)
}

// calculateDropRate formats the drop rate for display
func (ls *LootSimulator) calculateDropRate(weight, totalWeight int) string {
	if weight >= 100 {
		return "Common"
	} else if weight >= 30 {
		return "Uncommon"
	} else if weight >= 10 {
		return "Rare"
	} else {
		// Calculate approximate rate
		rate := int(math.Round(float64(totalWeight) / float64(weight)))
		return "~1/" + string(rune(rate))
	}
}

// SimulateAverageRewards calculates expected rewards using statistical averages instead of RNG
func SimulateAverageRewards(totalSearches int) ([]Reward, int) {
	rewards := make([]Reward, 0)
	totalValue := 0

	// Calculate total weight
	totalWeight := 0
	for _, item := range RewardTable {
		totalWeight += item.Weight
	}

	// Calculate expected drops based on drop rates and weights
	for _, item := range RewardTable {
		// Calculate expected quantity based on search count and drop rate
		dropChance := float64(item.Weight) / float64(totalWeight)

		// Adjust drop chances based on item category
		switch item.Category {
		case "essence":
			dropChance *= 0.9 // 90% chance per search
		case "catalysts":
			dropChance *= 0.7 // 70% chance per search
		case "runes":
			dropChance *= 0.4 // 40% chance per search
		case "tools":
			dropChance *= 0.15 // 15% chance per search
		case "outfit":
			dropChance *= 0.005 // 0.5% chance per search
		case "rare":
			dropChance *= 0.001 // 0.1% chance per search
		}

		expectedDrops := float64(totalSearches) * dropChance
		if expectedDrops < 0.1 && item.Category != "rare" {
			continue // Skip items with very low expected quantities
		}

		avgQuantity := (item.VarianceMin + item.VarianceMax) / 2
		totalQuantity := int(math.Round(expectedDrops * float64(avgQuantity)))

		if totalQuantity > 0 {
			reward := Reward{
				Name:     item.Name,
				Quantity: totalQuantity,
				Value:    item.Value,
				DropRate: calculateStaticDropRate(item.Weight, totalWeight),
			}
			rewards = append(rewards, reward)
			totalValue += totalQuantity * item.Value
		}
	}

	return rewards, totalValue
}

// calculateStaticDropRate formats the drop rate for static calculations
func calculateStaticDropRate(weight, totalWeight int) string {
	if weight >= 100 {
		return "Common"
	} else if weight >= 30 {
		return "Uncommon"
	} else if weight >= 10 {
		return "Rare"
	} else {
		rate := int(math.Round(float64(totalWeight) / float64(weight)))
		return "~1/" + string(rune(rate))
	}
}
