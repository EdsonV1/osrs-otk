package wintertodt

import (
	"math/rand"
	"time"
)

// SimulateLoot simulates Wintertodt loot for the given number of rounds
func SimulateLoot(rounds int) (map[string]any, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	lootCounts := make(map[string]int)
	totalValue := 0

	// Always get burnt page and supply crate per round
	for _, item := range CommonLoot {
		lootCounts[item.Name] = rounds * item.Quantity
		totalValue += lootCounts[item.Name] * item.Value
	}

	// Simulate supply crate contents (average 3-4 items per crate)
	for range rounds {
		itemsInCrate := 3 + r.Intn(2) // 3-4 items

		for range itemsInCrate {
			for _, item := range SupplyLoot {
				if r.Float64() < item.Rate {
					lootCounts[item.Name] += item.Quantity
					totalValue += item.Quantity * item.Value
					break // Only one item per slot
				}
			}
		}

		// Rare drops (independent rolls)
		for _, item := range RareLoot {
			if r.Float64() < item.Rate {
				lootCounts[item.Name] += item.Quantity
				if item.Name != "Phoenix" {
					totalValue += item.Quantity * item.Value
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
