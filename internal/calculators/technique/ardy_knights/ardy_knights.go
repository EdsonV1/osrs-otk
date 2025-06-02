package ardyknights

import (
	"fmt"
	"math"
	"sort"
)

type ArdyKnightResult struct {
	CalculatedSuccessRate float64 `json:"calculated_success_rate"`
	EffectiveXPPerAttempt float64 `json:"effective_xp_per_attempt"`
	EffectiveGPPerAttempt float64 `json:"effective_gp_per_attempt"`
	XPHour                int     `json:"xp_hour"`
	GPHour                int     `json:"gp_hour"`
	DamagePerHour         int     `json:"damage_per_hour"`
	FoodNeededPerHour     int     `json:"food_needed_per_hour"`
	ProfitPerHour         int     `json:"profit_per_hour"`

	CurrentThievingLevel int     `json:"current_thieving_level"`
	TargetThievingLevel  int     `json:"target_thieving_level"`
	CurrentTotalXP       int     `json:"current_total_xp"`
	TargetTotalXP        int     `json:"target_total_xp"`
	XPToTarget           int     `json:"xp_to_target"`
	HoursToTarget        float64 `json:"hours_to_target"`
	PickpocketsToTarget  int     `json:"pickpockets_to_target"`
}

func getArdyKnightBaseSuccessChance(level int) float64 {
	if level < 55 {
		return 0.0
	}

	if chance, ok := BaseSuccessChance[level]; ok {
		return chance
	}
	if level >= 99 {
		return BaseSuccessChance[99]
	}

	var sortedLevels []int
	for l := range BaseSuccessChance {
		sortedLevels = append(sortedLevels, l)
	}
	sort.Ints(sortedLevels)

	var lowerLevelKey int = -1
	var upperLevelKey int = -1

	for _, lKey := range sortedLevels {
		if lKey <= level {
			lowerLevelKey = lKey
		} else {
			upperLevelKey = lKey
			break
		}
	}

	if lowerLevelKey == -1 {

		if len(sortedLevels) > 0 && level < sortedLevels[0] {
			return 0.0
		}
		return 0.0
	}

	if upperLevelKey == -1 {

		return BaseSuccessChance[lowerLevelKey]
	}

	if lowerLevelKey == upperLevelKey {
		return BaseSuccessChance[lowerLevelKey]
	}

	lowerProb := BaseSuccessChance[lowerLevelKey]
	upperProb := BaseSuccessChance[upperLevelKey]

	// Linear interpolation: y = y1 + (y2-y1) * (x-x1) / (x2-x1)
	// x = level, x1 = lowerLevelKey, x2 = upperLevelKey
	// y1 = lowerProb, y2 = upperProb
	interpolatedChance := lowerProb + (upperProb-lowerProb)*(float64(level-lowerLevelKey))/float64(upperLevelKey-lowerLevelKey)

	return interpolatedChance
}

// CalculateArdyKnightStats calculates XP, GP, and other stats for pickpocketing Ardougne Knights.
func CalculateArdyKnightStats(
	currentThievingXP int,
	targetThievingXP int,
	HasArdyMed bool,
	hasThievingCape bool,
	hasRoguesOutfit bool,
	hasShadowVeil bool,
	hourlyPickpockets int,
	foodHealAmount int,
	foodCost int,
) (ArdyKnightResult, error) {

	currentLevel := GetLevelForXP(currentThievingXP)
	derivedTargetLevel := GetLevelForXP(targetThievingXP)

	if currentLevel < 55 {
		return ArdyKnightResult{}, fmt.Errorf("current thieving level (%d, from XP %d) must be at least 55 to pickpocket Ardougne Knights effectively", currentLevel, currentThievingXP)
	}

	if targetThievingXP <= currentThievingXP {
		return ArdyKnightResult{}, fmt.Errorf("target thieving XP (%d) must be greater than current thieving XP (%d)", targetThievingXP, currentThievingXP)
	}

	if hourlyPickpockets <= 0 {
		return ArdyKnightResult{}, fmt.Errorf("hourly pickpockets must be greater than 0")
	}

	baseChance := getArdyKnightBaseSuccessChance(currentLevel)
	totalSuccessChance := baseChance

	if HasArdyMed {
		totalSuccessChance += ArdyHardBoost
	}
	if hasThievingCape {
		totalSuccessChance += ThievingCapeBoost
	}
	if hasShadowVeil {
		totalSuccessChance += ShadowVeilBoost
	}
	if currentLevel >= 99 && HasArdyMed && hasThievingCape && hasShadowVeil {
		totalSuccessChance = math.Min(totalSuccessChance, 0.995)
	} else {
		totalSuccessChance = math.Min(totalSuccessChance, 1.0)
	}

	failureRate := 1.0 - totalSuccessChance
	effectiveXPPerAttempt := BaseXPPerPickpocket * totalSuccessChance
	avgCoinsPerSuccess := float64(MinCoinDrop+MaxCoinDrop) / 2.0
	if hasRoguesOutfit {
		avgCoinsPerSuccess *= 2.0
	}
	effectiveGPPerAttempt := avgCoinsPerSuccess * totalSuccessChance

	xpPerHour := int(math.Round(effectiveXPPerAttempt * float64(hourlyPickpockets)))
	gpPerHour := int(math.Round(effectiveGPPerAttempt * float64(hourlyPickpockets)))
	damagePerHour := int(math.Round(failureRate * float64(StunDamage) * float64(hourlyPickpockets)))
	foodNeededPerHour := 0
	if foodHealAmount > 0 {
		foodNeededPerHour = int(math.Ceil(float64(damagePerHour) / float64(foodHealAmount)))
	}
	foodCostPerHour := foodNeededPerHour * foodCost
	profitPerHour := gpPerHour - foodCostPerHour

	xpToTarget := targetThievingXP - currentThievingXP

	hoursToTarget := 0.0
	pickpocketsToTarget := 0
	if xpToTarget > 0 && xpPerHour > 0 {
		hoursToTarget = float64(xpToTarget) / float64(xpPerHour)
		if effectiveXPPerAttempt > 0 {
			pickpocketsToTarget = int(math.Ceil(float64(xpToTarget) / effectiveXPPerAttempt))
		} else if xpToTarget > 0 {

		}
	} else if xpToTarget <= 0 {
		xpToTarget = 0
		hoursToTarget = 0
		pickpocketsToTarget = 0
	}

	return ArdyKnightResult{
		CalculatedSuccessRate: totalSuccessChance,
		EffectiveXPPerAttempt: effectiveXPPerAttempt,
		EffectiveGPPerAttempt: effectiveGPPerAttempt,
		XPHour:                xpPerHour,
		GPHour:                gpPerHour,
		DamagePerHour:         damagePerHour,
		FoodNeededPerHour:     foodNeededPerHour,
		ProfitPerHour:         profitPerHour,

		CurrentThievingLevel: currentLevel,
		TargetThievingLevel:  derivedTargetLevel,
		CurrentTotalXP:       currentThievingXP,
		TargetTotalXP:        targetThievingXP,
		XPToTarget:           xpToTarget,
		HoursToTarget:        hoursToTarget,
		PickpocketsToTarget:  pickpocketsToTarget,
	}, nil
}
