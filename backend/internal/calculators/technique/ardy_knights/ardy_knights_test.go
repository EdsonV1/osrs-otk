package ardyknights

import (
	"fmt"
	"testing"
)

func TestCalculateArdyKnightStats(t *testing.T) {
	tests := []struct {
		name              string
		currentThievingXP int
		targetThievingXP  int
		hasArdyMed        bool
		hasThievingCape   bool
		hasRoguesOutfit   bool
		hasShadowVeil     bool
		hourlyPickpockets int
		foodHealAmount    int
		foodCost          int
		expectError       bool
		expectedMinXPHour int
		expectedMaxXPHour int
		expectedMinGPHour int
		expectedMaxGPHour int
	}{
		{
			name:              "Level 55 baseline",
			currentThievingXP: 166636, // Level 55
			targetThievingXP:  200000,
			hasArdyMed:        false,
			hasThievingCape:   false,
			hasRoguesOutfit:   false,
			hasShadowVeil:     false,
			hourlyPickpockets: 3000,
			foodHealAmount:    20,
			foodCost:          500,
			expectError:       false,
			expectedMinXPHour: 160000,
			expectedMaxXPHour: 170000,
			expectedMinGPHour: 45000,
			expectedMaxGPHour: 55000,
		},
		{
			name:              "Level 99 with all bonuses",
			currentThievingXP: 13034431, // Level 99
			targetThievingXP:  14000000,
			hasArdyMed:        true,
			hasThievingCape:   true,
			hasRoguesOutfit:   true,
			hasShadowVeil:     true,
			hourlyPickpockets: 5000,
			foodHealAmount:    20,
			foodCost:          500,
			expectError:       false,
			expectedMinXPHour: 410000,
			expectedMaxXPHour: 420000,
			expectedMinGPHour: 250000,
			expectedMaxGPHour: 260000,
		},
		{
			name:              "Level 70 with some bonuses",
			currentThievingXP: 737627, // Level 70
			targetThievingXP:  1200000,
			hasArdyMed:        true,
			hasThievingCape:   false,
			hasRoguesOutfit:   true,
			hasShadowVeil:     false,
			hourlyPickpockets: 4000,
			foodHealAmount:    20,
			foodCost:          400,
			expectError:       false,
			expectedMinXPHour: 300000,
			expectedMaxXPHour: 310000,
			expectedMinGPHour: 180000,
			expectedMaxGPHour: 190000,
		},
		{
			name:              "Invalid: level too low",
			currentThievingXP: 100000, // Level 50
			targetThievingXP:  200000,
			hasArdyMed:        false,
			hasThievingCape:   false,
			hasRoguesOutfit:   false,
			hasShadowVeil:     false,
			hourlyPickpockets: 3000,
			foodHealAmount:    20,
			foodCost:          500,
			expectError:       true,
		},
		{
			name:              "Invalid: target XP lower than current",
			currentThievingXP: 200000,
			targetThievingXP:  150000,
			hasArdyMed:        false,
			hasThievingCape:   false,
			hasRoguesOutfit:   false,
			hasShadowVeil:     false,
			hourlyPickpockets: 3000,
			foodHealAmount:    20,
			foodCost:          500,
			expectError:       true,
		},
		{
			name:              "Invalid: zero pickpockets per hour",
			currentThievingXP: 155000,
			targetThievingXP:  200000,
			hasArdyMed:        false,
			hasThievingCape:   false,
			hasRoguesOutfit:   false,
			hasShadowVeil:     false,
			hourlyPickpockets: 0,
			foodHealAmount:    20,
			foodCost:          500,
			expectError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateArdyKnightStats(
				tt.currentThievingXP,
				tt.targetThievingXP,
				tt.hasArdyMed,
				tt.hasThievingCape,
				tt.hasRoguesOutfit,
				tt.hasShadowVeil,
				tt.hourlyPickpockets,
				tt.foodHealAmount,
				tt.foodCost,
			)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Test XP per hour ranges
			if result.XPHour < tt.expectedMinXPHour || result.XPHour > tt.expectedMaxXPHour {
				t.Errorf("XP per hour out of range: got %d, want between %d and %d",
					result.XPHour, tt.expectedMinXPHour, tt.expectedMaxXPHour)
			}

			// Test GP per hour ranges
			if result.GPHour < tt.expectedMinGPHour || result.GPHour > tt.expectedMaxGPHour {
				t.Errorf("GP per hour out of range: got %d, want between %d and %d",
					result.GPHour, tt.expectedMinGPHour, tt.expectedMaxGPHour)
			}

			// Test success rate is reasonable
			if result.CalculatedSuccessRate < 0 || result.CalculatedSuccessRate > 1 {
				t.Errorf("Success rate out of range: got %f, want between 0 and 1",
					result.CalculatedSuccessRate)
			}

			// Test that bonuses increase success rate
			if tt.hasArdyMed || tt.hasThievingCape || tt.hasShadowVeil {
				if result.CalculatedSuccessRate < 0.5 {
					t.Errorf("Success rate too low with bonuses: got %f", result.CalculatedSuccessRate)
				}
			}

			// Test that rogues outfit doubles coin value effectively
			if tt.hasRoguesOutfit && result.EffectiveGPPerAttempt < 40 {
				t.Errorf("GP per attempt should be higher with rogues outfit: got %f", result.EffectiveGPPerAttempt)
			}

			// Test level calculations
			if result.CurrentThievingLevel < 55 {
				t.Errorf("Current level should be at least 55, got %d", result.CurrentThievingLevel)
			}

			// Test progress calculations
			if result.XPToTarget < 0 {
				t.Errorf("XP to target should not be negative: got %d", result.XPToTarget)
			}

			if result.HoursToTarget < 0 {
				t.Errorf("Hours to target should not be negative: got %f", result.HoursToTarget)
			}
		})
	}
}

func TestGetArdyKnightBaseSuccessChance(t *testing.T) {
	tests := []struct {
		level    int
		expected float64
		minRate  float64
		maxRate  float64
	}{
		{54, 0.0, 0.0, 0.0},     // Below minimum level
		{55, 0.65, 0.6, 0.7},    // Minimum level
		{70, 0.80, 0.75, 0.85},  // Mid-level
		{99, 0.97, 0.95, 0.98},  // Maximum level
		{120, 0.97, 0.95, 0.98}, // Above maximum (should cap at 99)
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Level_%d", tt.level), func(t *testing.T) {
			result := getArdyKnightBaseSuccessChance(tt.level)

			if tt.expected > 0 {
				if result != tt.expected {
					t.Errorf("Expected exact value %f for level %d, got %f",
						tt.expected, tt.level, result)
				}
			} else {
				if result < tt.minRate || result > tt.maxRate {
					t.Errorf("Success chance for level %d out of range: got %f, want between %f and %f",
						tt.level, result, tt.minRate, tt.maxRate)
				}
			}

			// Test that higher levels generally have higher success rates
			if tt.level >= 55 && tt.level < 99 {
				lowerResult := getArdyKnightBaseSuccessChance(tt.level - 5)
				if result < lowerResult && tt.level > 60 {
					t.Errorf("Success chance should increase with level: level %d (%f) should be >= level %d (%f)",
						tt.level, result, tt.level-5, lowerResult)
				}
			}
		})
	}
}

func TestGetLevelForXP(t *testing.T) {
	tests := []struct {
		xp            int
		expectedLevel int
	}{
		{0, 1},
		{83, 2},
		{13034431, 99}, // Exactly level 99
		{13034430, 98}, // Just below level 99
		{50000000, 99}, // Way above level 99 (should cap at 99)
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("XP_%d", tt.xp), func(t *testing.T) {
			result := GetLevelForXP(tt.xp)
			if result != tt.expectedLevel {
				t.Errorf("GetLevelForXP(%d) = %d, want %d", tt.xp, result, tt.expectedLevel)
			}
		})
	}
}

// Benchmark tests
func BenchmarkCalculateArdyKnightStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateArdyKnightStats(
			737627,  // Level 70
			1200000, // Target XP
			true,    // Has Ardy Med
			false,   // No Thieving Cape
			true,    // Has Rogues Outfit
			false,   // No Shadow Veil
			4000,    // Hourly pickpockets
			20,      // Food heal amount
			400,     // Food cost
		)
	}
}

func BenchmarkGetArdyKnightBaseSuccessChance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getArdyKnightBaseSuccessChance(75)
	}
}
