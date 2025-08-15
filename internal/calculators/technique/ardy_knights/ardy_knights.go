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

// GetCalculationProTips provides detailed information about how Ardougne Knights calculations work
func GetCalculationProTips() map[string]any {
	return map[string]any{
		"calculation_methodology": map[string]any{
			"xp_rates_source": "Based on official success rates and community testing",
			"base_formula":    "Success rate × Base XP (84.3) × Attempts per hour",
			"data_points": []map[string]any{
				{"level": 55, "xp_per_hour": 60000, "note": "Minimum access level"},
				{"level": 70, "xp_per_hour": 90000, "note": "Improved success rate"},
				{"level": 85, "xp_per_hour": 120000, "note": "High efficiency"},
				{"level": 99, "xp_per_hour": 150000, "note": "Maximum rates with full setup"},
			},
		},
		"game_mechanics": map[string]any{
			"pickpocket_speed": "2.4 second intervals (1500 attempts/hour max)",
			"success_formula":  "Base rate + Equipment bonuses + Level scaling",
			"xp_sources": []string{
				"Successful pickpockets (84.3 XP each)",
				"No bonus XP sources",
			},
		},
		"factors_considered": []string{
			"Thieving level (affects base success rate)",
			"Ardougne Medium diary (10% success boost)",
			"Thieving cape (10% success boost)",
			"Shadow veil spell (15% success boost)",
			"Rogue's outfit (doubles coin rewards)",
			"Player click consistency and timing",
		},
		"accuracy_notes": map[string]any{
			"rates_vary": "Individual XP rates can vary ±20% based on:",
			"variance_factors": []string{
				"Click timing and consistency",
				"Lag and connection stability",
				"Interruptions from failures",
				"Food consumption timing",
			},
			"calculation_basis": "Rates assume optimal clicking and minimal interruptions",
		},
		"pro_tips": []map[string]string{
			{
				"tip":         "Equipment Setup",
				"description": "Ardougne diary, thieving cape, and shadow veil dramatically improve rates",
			},
			{
				"tip":         "Positioning",
				"description": "Position knight near a bank for easy food access",
			},
			{
				"tip":         "Food Strategy",
				"description": "Use cheap, stackable food like wines or cakes for efficiency",
			},
			{
				"tip":         "Timing Optimization",
				"description": "Consistent 2.4s intervals maximize attempts per hour",
			},
			{
				"tip":         "Profit Focus",
				"description": "Rogue's outfit doubles coin rewards - essential for profitable training",
			},
		},
		"reward_calculation": map[string]any{
			"base_coins":     "50-100 coins per successful pickpocket",
			"rogue_bonus":    "Doubles coin rewards when wearing full Rogue's outfit",
			"profit_factors": "Success rate, coin rewards, food costs",
			"gp_per_hour":    "Calculated from successful attempts minus food expenses",
		},
	}
}
