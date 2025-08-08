package gotr

import (
	"testing"
)

func TestCalculateGOTRData(t *testing.T) {
	tests := []struct {
		name         string
		currentLevel int
		targetLevel  int
		expectError  bool
		minXPPerHour float64
		maxXPPerHour float64
	}{
		{
			name:         "Level 77 to 99",
			currentLevel: 77,
			targetLevel:  99,
			expectError:  false,
			minXPPerHour: 180000, // Minimum expected XP/hour at level 77
			maxXPPerHour: 240000, // Maximum expected XP/hour
		},
		{
			name:         "Level 50 to 77",
			currentLevel: 50,
			targetLevel:  77,
			expectError:  false,
			minXPPerHour: 120000, // Lower rates for pre-77
			maxXPPerHour: 165000,
		},
		{
			name:         "Level 27 to 99",
			currentLevel: 27,
			targetLevel:  99,
			expectError:  false,
			minXPPerHour: 130000, // Lower because average level is 63 (pre-77)
			maxXPPerHour: 150000,
		},
		{
			name:         "Invalid: same level",
			currentLevel: 77,
			targetLevel:  77,
			expectError:  true,
		},
		{
			name:         "Invalid: target lower than current",
			currentLevel: 90,
			targetLevel:  80,
			expectError:  true,
		},
		{
			name:         "Invalid: level too low",
			currentLevel: 20,
			targetLevel:  99,
			expectError:  true,
		},
		{
			name:         "Invalid: level too high",
			currentLevel: 27,
			targetLevel:  150,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateGOTRData(tt.currentLevel, tt.targetLevel)

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

			// Validate basic result structure
			if result.CurrentLevel != tt.currentLevel {
				t.Errorf("Current level mismatch: got %d, want %d", result.CurrentLevel, tt.currentLevel)
			}

			if result.TargetLevel != tt.targetLevel {
				t.Errorf("Target level mismatch: got %d, want %d", result.TargetLevel, tt.targetLevel)
			}

			if result.XPNeeded <= 0 {
				t.Errorf("XP needed should be positive, got %d", result.XPNeeded)
			}

			if result.GamesNeeded <= 0 {
				t.Errorf("Games needed should be positive, got %d", result.GamesNeeded)
			}

			if result.HoursNeeded <= 0 {
				t.Errorf("Hours needed should be positive, got %f", result.HoursNeeded)
			}

			// Check XP per hour is within reasonable bounds
			if result.AverageXPPerHour < tt.minXPPerHour || result.AverageXPPerHour > tt.maxXPPerHour {
				t.Errorf("XP per hour out of range: got %f, want between %f and %f",
					result.AverageXPPerHour, tt.minXPPerHour, tt.maxXPPerHour)
			}

			// Pet chance should be between 0 and 100
			if result.PetChancePercentage < 0 || result.PetChancePercentage > 100 {
				t.Errorf("Pet chance percentage out of range: got %f", result.PetChancePercentage)
			}

			// Should have some rewards
			if len(result.EstimatedRewards) == 0 {
				t.Errorf("Should have estimated rewards")
			}

			// Total reward value should be positive
			if result.TotalRewardValue <= 0 {
				t.Errorf("Total reward value should be positive, got %d", result.TotalRewardValue)
			}

			// GP per hour should be reasonable (at least 100k)
			if result.GPPerHour < 100000 {
				t.Errorf("GP per hour seems too low: got %f", result.GPPerHour)
			}
		})
	}
}

func TestCalculateXPPerGame(t *testing.T) {
	tests := []struct {
		currentLevel int
		targetLevel  int
		minXP        float64
		maxXP        float64
	}{
		{77, 80, 30000, 35000}, // Optimal level range
		{50, 60, 20000, 28000}, // Pre-optimal
		{90, 99, 38000, 47000}, // High level
	}

	for _, tt := range tests {
		xpPerGame := calculateXPPerGame(tt.currentLevel, tt.targetLevel)

		if xpPerGame < tt.minXP || xpPerGame > tt.maxXP {
			t.Errorf("XP per game for levels %d-%d out of range: got %f, want between %f and %f",
				tt.currentLevel, tt.targetLevel, xpPerGame, tt.minXP, tt.maxXP)
		}
	}
}

func TestSimulateAverageRewards(t *testing.T) {
	// Test with a reasonable number of searches
	searches := 1000

	rewards, totalValue := SimulateAverageRewards(searches)

	if len(rewards) == 0 {
		t.Errorf("Should have rewards from %d searches", searches)
	}

	if totalValue <= 0 {
		t.Errorf("Total value should be positive, got %d", totalValue)
	}

	// Should have common items like guardian essence
	hasEssence := false
	for _, reward := range rewards {
		if reward.Name == "Guardian essence" {
			hasEssence = true
			break
		}
	}

	if !hasEssence {
		t.Errorf("Should have guardian essence in rewards")
	}
}
