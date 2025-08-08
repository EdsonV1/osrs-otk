package wintertodt

import (
	"fmt"
	"testing"
)

func TestCalculateWintertodtData(t *testing.T) {
	tests := []struct {
		name              string
		firemakingLevel   int
		roundsPerHour     float64
		totalRounds       int
		expectError       bool
		expectedMinXPHour float64
		expectedMaxXPHour float64
		expectedMinValue  int
		expectedMaxValue  int
	}{
		{
			name:              "Level 50 baseline",
			firemakingLevel:   50,
			roundsPerHour:     4.0,
			totalRounds:       100,
			expectError:       false,
			expectedMinXPHour: 150000,
			expectedMaxXPHour: 200000,
			expectedMinValue:  500000,
			expectedMaxValue:  1500000,
		},
		{
			name:              "Level 75 mid-level",
			firemakingLevel:   75,
			roundsPerHour:     5.0,
			totalRounds:       200,
			expectError:       false,
			expectedMinXPHour: 200000,
			expectedMaxXPHour: 280000,
			expectedMinValue:  1000000,
			expectedMaxValue:  2500000,
		},
		{
			name:              "Level 99 maximum",
			firemakingLevel:   99,
			roundsPerHour:     6.0,
			totalRounds:       500,
			expectError:       false,
			expectedMinXPHour: 250000,
			expectedMaxXPHour: 350000,
			expectedMinValue:  2500000,
			expectedMaxValue:  8000000,
		},
		{
			name:            "Invalid: level too low",
			firemakingLevel: 40,
			roundsPerHour:   4.0,
			totalRounds:     100,
			expectError:     true,
		},
		{
			name:            "Invalid: negative rounds per hour",
			firemakingLevel: 70,
			roundsPerHour:   -1.0,
			totalRounds:     100,
			expectError:     true,
		},
		{
			name:            "Invalid: zero total rounds",
			firemakingLevel: 70,
			roundsPerHour:   4.0,
			totalRounds:     0,
			expectError:     true,
		},
		{
			name:            "Invalid: negative total rounds",
			firemakingLevel: 70,
			roundsPerHour:   4.0,
			totalRounds:     -50,
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateWintertodtData(tt.firemakingLevel, tt.roundsPerHour, tt.totalRounds)

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

			// Test XP per hour calculation
			if result.AverageExpHour < tt.expectedMinXPHour || result.AverageExpHour > tt.expectedMaxXPHour {
				t.Errorf("XP per hour out of range: got %f, want between %f and %f",
					result.AverageExpHour, tt.expectedMinXPHour, tt.expectedMaxXPHour)
			}

			// Test total value range
			if result.TotalValue < tt.expectedMinValue || result.TotalValue > tt.expectedMaxValue {
				t.Errorf("Total value out of range: got %d, want between %d and %d",
					result.TotalValue, tt.expectedMinValue, tt.expectedMaxValue)
			}

			// Test experience scaling with level
			if result.TotalExperience <= 0 {
				t.Errorf("Total experience should be positive, got %d", result.TotalExperience)
			}

			// Test pet chance is reasonable (0-100%)
			if result.PetChance < 0 || result.PetChance > 100 {
				t.Errorf("Pet chance out of range: got %f, want between 0 and 100", result.PetChance)
			}

			// Test that higher rounds give higher pet chance
			if tt.totalRounds > 100 && result.PetChance < 1.0 {
				t.Errorf("Pet chance should be reasonable with %d rounds: got %f", tt.totalRounds, result.PetChance)
			}

			// Test total time calculation
			expectedTime := float64(tt.totalRounds) / tt.roundsPerHour
			if result.TotalTime != expectedTime {
				t.Errorf("Total time incorrect: got %f, want %f", result.TotalTime, expectedTime)
			}

			// Test estimated loot is not empty
			if len(result.EstimatedLoot) == 0 {
				t.Errorf("Estimated loot should not be empty")
			}

			// Test loot has reasonable items
			foundLogs := false
			for itemName := range result.EstimatedLoot {
				if itemName == "yew_logs" || itemName == "magic_logs" {
					foundLogs = true
					break
				}
			}
			if !foundLogs {
				t.Errorf("Should have some log drops in estimated loot")
			}
		})
	}
}

func TestSimulateLoot(t *testing.T) {
	tests := []struct {
		name        string
		totalRounds int
		minValue    int
		maxValue    int
		minItems    int
	}{
		{"Small run", 10, 50000, 200000, 3},
		{"Medium run", 100, 500000, 2000000, 5},
		{"Large run", 1000, 5000000, 20000000, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loot, totalValue := SimulateLoot(tt.totalRounds)

			if totalValue < tt.minValue || totalValue > tt.maxValue {
				t.Errorf("Total value out of range: got %d, want between %d and %d",
					totalValue, tt.minValue, tt.maxValue)
			}

			if len(loot) < tt.minItems {
				t.Errorf("Should have at least %d different items, got %d", tt.minItems, len(loot))
			}

			// Test that all quantities are positive
			for itemName, quantity := range loot {
				if q, ok := quantity.(int); ok && q <= 0 {
					t.Errorf("Item %s should have positive quantity, got %d", itemName, q)
				}
			}

			// Test consistency - more rounds should generally mean more value
			if tt.totalRounds > 100 {
				smallLoot, smallValue := SimulateLoot(10)
				if totalValue <= smallValue {
					t.Errorf("More rounds should generally give more value: %d rounds gave %d, 10 rounds gave %d",
						tt.totalRounds, totalValue, smallValue)
				}
				if len(loot) < len(smallLoot) {
					t.Errorf("More rounds should generally give more item types: %d rounds gave %d types, 10 rounds gave %d types",
						tt.totalRounds, len(loot), len(smallLoot))
				}
			}
		})
	}
}

// Test that pet chance calculation is accurate
func TestPetChanceCalculation(t *testing.T) {
	tests := []struct {
		rounds    int
		minChance float64
		maxChance float64
	}{
		{1, 0.01, 0.03},    // Single round
		{100, 1.8, 2.2},    // Medium rounds
		{1000, 18.0, 22.0}, // Many rounds
		{5000, 63.0, 67.0}, // Very many rounds (approaching but not exceeding ~95%)
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Rounds_%d", tt.rounds), func(t *testing.T) {
			result, err := CalculateWintertodtData(75, 5.0, tt.rounds)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result.PetChance < tt.minChance || result.PetChance > tt.maxChance {
				t.Errorf("Pet chance for %d rounds out of range: got %f, want between %f and %f",
					tt.rounds, result.PetChance, tt.minChance, tt.maxChance)
			}
		})
	}
}

// Benchmark tests
func BenchmarkCalculateWintertodtData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateWintertodtData(75, 5.0, 100)
	}
}

func BenchmarkSimulateLoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimulateLoot(100)
	}
}
