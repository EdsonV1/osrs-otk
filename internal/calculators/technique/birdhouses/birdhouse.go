package birdhouses

import (
	"fmt"
	"log"
	"math"
)

type BirdhouseResult struct {
	EstimatedNests float64                   `json:"estimated_nests"`
	HunterXP       int                       `json:"hunter_xp"`
	CraftingXP     int                       `json:"crafting_xp"`
	DaysLowEff     int                       `json:"days_low_efficiency"`
	DaysMedEff     int                       `json:"days_medium_efficiency"`
	DaysHighEff    int                       `json:"days_high_efficiency"`
	SeedDrops      map[string]map[string]int `json:"seed_drops"`
	TotalLoot      int                       `json:"total_loot"`
}

var avgNests = map[string]float64{
	"regular": 0.5, "oak": 0.75, "willow": 1.0,
	"teak": 1.25, "maple": 1.5, "mahogany": 1.75,
	"yew": 2.0, "magic": 2.25, "redwood": 2.5,
}

var hunterXPPerBirdhouse = map[string]int{
	"regular":  280,
	"oak":      420,
	"willow":   560,
	"teak":     700,
	"maple":    820,
	"mahogany": 960,
	"yew":      1020,
	"magic":    1140,
	"redwood":  1200,
}

var craftingXPPerBirdhouse = map[string]int{
	"regular":  15,
	"oak":      20,
	"willow":   25,
	"teak":     30,
	"maple":    35,
	"mahogany": 40,
	"yew":      45,
	"magic":    50,
	"redwood":  55,
}

func CalculateBirdhouseData(typ string, quantity int) (BirdhouseResult, error) {
	if quantity <= 0 {
		return BirdhouseResult{}, fmt.Errorf("quantity must be positive, got %d", quantity)
	}

	rate, ok := avgNests[typ]
	if !ok {
		return BirdhouseResult{}, fmt.Errorf("unknown birdhouse type: %s", typ)
	}

	perHunterXP, ok := hunterXPPerBirdhouse[typ]
	if !ok {
		return BirdhouseResult{}, fmt.Errorf("unknown hunter XP for birdhouse type: %s", typ)
	}

	perCraftingXP, ok := craftingXPPerBirdhouse[typ]
	if !ok {
		return BirdhouseResult{}, fmt.Errorf("unknown crafting XP for birdhouse type: %s", typ)
	}

	nests := float64(quantity) * rate
	totalHunterXP := perHunterXP * quantity
	totalCraftingXP := perCraftingXP * quantity

	runsFloat := math.Ceil(float64(quantity) / 4.0)
	//runs := int(runsFloat)

	seedDrops, totalLoot, err := SimulateNestLoot(int(math.Round(nests)))
	if err != nil {
		log.Fatalf("Simulation error: %v", err)
	}

	totalNestLoot := nests * 7107

	return BirdhouseResult{
		EstimatedNests: nests,
		HunterXP:       totalHunterXP,
		CraftingXP:     totalCraftingXP,
		DaysLowEff:     int(math.Ceil(runsFloat / 2.0)),  // 2 runs/day
		DaysMedEff:     int(math.Ceil(runsFloat / 7.0)),  // 7 runs/day
		DaysHighEff:    int(math.Ceil(runsFloat / 14.0)), // 14 runs/day
		SeedDrops:      seedDrops,
		TotalLoot:      totalLoot + int(totalNestLoot),
	}, nil
}
