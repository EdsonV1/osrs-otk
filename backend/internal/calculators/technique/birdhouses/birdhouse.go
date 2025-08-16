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
	return CalculateBirdhouseDataWithPrices(typ, quantity, nil)
}

func CalculateBirdhouseDataWithPrices(typ string, quantity int, livePrices map[string]int) (BirdhouseResult, error) {
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

	seedDrops, totalLoot, err := SimulateNestLootWithPrices(int(math.Round(nests)), livePrices)
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

// GetCalculationProTips provides detailed information about how Birdhouse calculations work
func GetCalculationProTips() map[string]any {
	return map[string]any{
		"calculation_methodology": map[string]any{
			"xp_rates_source": "Based on official birdhouse mechanics and Hunter XP rates",
			"base_formula":    "XP per birdhouse × Number of houses × Runs completed",
			"data_points": []map[string]any{
				{"log_type": "Regular", "hunter_xp": 280, "note": "Basic birdhouse"},
				{"log_type": "Oak", "hunter_xp": 420, "note": "Early upgrade"},
				{"log_type": "Yew", "hunter_xp": 1020, "note": "High-level option"},
				{"log_type": "Redwood", "hunter_xp": 1200, "note": "Maximum XP per house"},
			},
		},
		"game_mechanics": map[string]any{
			"run_duration": "50 minutes per full cycle (4 birdhouses)",
			"setup_time":   "2-3 minutes to place and fill all birdhouses",
			"xp_sources": []string{
				"Hunter XP from emptying full birdhouses",
				"Crafting XP from creating birdhouses",
				"No bonus XP modifiers available",
			},
		},
		"factors_considered": []string{
			"Log type (determines XP and nest rates)",
			"Number of birdhouses (max 4 with achievement diary)",
			"Completion frequency (affects daily XP rates)",
			"Bird nest drop rates (varies by log type)",
			"Seed market values for profit calculation",
		},
		"accuracy_notes": map[string]any{
			"rates_vary": "Individual results can vary ±15% based on:",
			"variance_factors": []string{
				"RNG in bird nest drops",
				"Seed types from nests",
				"Market price fluctuations",
				"Consistency in run timing",
			},
			"calculation_basis": "Rates assume optimal timing and consistent completion",
		},
		"pro_tips": []map[string]string{
			{
				"tip":         "Log Selection",
				"description": "Higher-tier logs give more XP and bird nests but cost more upfront",
			},
			{
				"tip":         "Timing Strategy",
				"description": "Check birdhouses every 50 minutes for maximum efficiency",
			},
			{
				"tip":         "Achievement Diary",
				"description": "Complete Western Provinces diary for access to all 4 birdhouse locations",
			},
			{
				"tip":         "Profit Optimization",
				"description": "Bird nests contain valuable tree seeds - check current market prices",
			},
			{
				"tip":         "Low Maintenance",
				"description": "Excellent passive training method - combine with other activities",
			},
		},
		"reward_calculation": map[string]any{
			"nest_rates":     "0.5-2.5 bird nests per birdhouse depending on log type",
			"seed_variety":   "Tree seeds, fruit tree seeds, and regular seeds from nests",
			"profit_factors": "Nest drop rates, seed values, log costs",
			"gp_calculation": "Based on average seed values and nest contents",
		},
	}
}
