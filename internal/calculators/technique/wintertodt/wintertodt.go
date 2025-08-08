package wintertodt

import (
	"fmt"
	"math"
)

type WintertodtResult struct {
	TotalExperience int            `json:"total_experience"`
	AverageExpHour  float64        `json:"average_exp_hour"`
	PetChance       float64        `json:"pet_chance"`
	EstimatedLoot   map[string]any `json:"estimated_loot"`
	TotalValue      int            `json:"total_value"`
	TotalTime       float64        `json:"total_time"`
}

func CalculateWintertodtData(firemakingLevel int, roundsPerHour float64, totalRounds int) (WintertodtResult, error) {
	if firemakingLevel < 50 {
		return WintertodtResult{}, fmt.Errorf("firemaking level must be at least 50")
	}

	if roundsPerHour <= 0 || totalRounds <= 0 {
		return WintertodtResult{}, fmt.Errorf("rounds per hour and total rounds must be positive")
	}

	// Experience calculation (base + bonuses)
	baseExp := BaseExpPerRound
	bonusExp := int(math.Floor(float64(firemakingLevel) * ExpMultiplier))
	expPerRound := baseExp + bonusExp
	totalExp := expPerRound * totalRounds

	// Time calculation
	totalTime := float64(totalRounds) / roundsPerHour
	avgExpHour := float64(totalExp) / totalTime

	// Pet chance calculation (1/5000 per supply crate)
	petChance := 1.0 - math.Pow(1.0-PetRatePerCrate, float64(totalRounds))

	// Loot simulation
	estimatedLoot, totalValue := SimulateLoot(totalRounds)

	return WintertodtResult{
		TotalExperience: totalExp,
		AverageExpHour:  avgExpHour,
		PetChance:       petChance * 100, // Convert to percentage
		EstimatedLoot:   estimatedLoot,
		TotalValue:      totalValue,
		TotalTime:       totalTime,
	}, nil
}
