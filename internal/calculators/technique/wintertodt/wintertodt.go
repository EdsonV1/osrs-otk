package wintertodt

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type WintertodtResult struct {
	TotalExperience int                    `json:"total_experience"`
	AverageExpHour  float64                `json:"average_exp_hour"`
	PetChance       float64                `json:"pet_chance"`
	EstimatedLoot   map[string]any `json:"estimated_loot"`
	TotalValue      int                    `json:"total_value"`
	TotalTime       float64                `json:"total_time"`
}

type LootItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Value    int     `json:"value"`
	Rate     float64 `json:"rate"`
}

var commonLoot = []LootItem{
	{"Burnt page", 1, 750, 1.0},
	{"Supply crate", 1, 0, 1.0},
}

var rareLoot = []LootItem{
	{"Torstol seeds", 1, 58000, 0.02},
	{"Magic seeds", 1, 104000, 0.015},
	{"Palm tree seeds", 1, 37000, 0.025},
	{"Yew seeds", 1, 67000, 0.02},
	{"Dragon axe", 1, 8500000, 0.0001},
	{"Phoenix", 1, 0, 0.0002},
}

var supplyLoot = []LootItem{
	{"Grimy ranarr weed", 2, 7000, 0.1},
	{"Grimy snapdragon", 2, 11000, 0.08},
	{"Grimy torstol", 1, 25000, 0.05},
	{"Uncut diamond", 1, 2800, 0.12},
	{"Pure essence", 50, 4, 0.2},
	{"Raw shark", 3, 800, 0.15},
}

func CalculateWintertodtData(firemakingLevel int, roundsPerHour float64, totalRounds int) (WintertodtResult, error) {
	if firemakingLevel < 50 {
		return WintertodtResult{}, fmt.Errorf("firemaking level must be at least 50")
	}

	if roundsPerHour <= 0 || totalRounds <= 0 {
		return WintertodtResult{}, fmt.Errorf("rounds per hour and total rounds must be positive")
	}

	// Experience calculation (base + bonuses)
	baseExp := 740
	bonusExp := int(math.Floor(float64(firemakingLevel) * 13.6))
	expPerRound := baseExp + bonusExp
	totalExp := expPerRound * totalRounds

	// Time calculation
	totalTime := float64(totalRounds) / roundsPerHour
	avgExpHour := float64(totalExp) / totalTime

	// Pet chance calculation (1/5000 per supply crate)
	petChance := 1.0 - math.Pow(4999.0/5000.0, float64(totalRounds))

	// Loot simulation
	estimatedLoot, totalValue := simulateWintertodtLoot(totalRounds)

	return WintertodtResult{
		TotalExperience: totalExp,
		AverageExpHour:  avgExpHour,
		PetChance:       petChance * 100, // Convert to percentage
		EstimatedLoot:   estimatedLoot,
		TotalValue:      totalValue,
		TotalTime:       totalTime,
	}, nil
}

func simulateWintertodtLoot(rounds int) (map[string]any, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	lootCounts := make(map[string]int)
	totalValue := 0

	// Always get burnt page and supply crate per round
	for _, item := range commonLoot {
		lootCounts[item.Name] = rounds * item.Quantity
		totalValue += lootCounts[item.Name] * item.Value
	}

	// Simulate supply crate contents (average 3-4 items per crate)
	for range rounds {
		itemsInCrate := 3 + r.Intn(2) // 3-4 items

		for range itemsInCrate {
			for _, item := range supplyLoot {
				if r.Float64() < item.Rate {
					lootCounts[item.Name] += item.Quantity
					totalValue += item.Quantity * item.Value
					break // Only one item per slot
				}
			}
		}

		// Rare drops (independent rolls)
		for _, item := range rareLoot {
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