package calculators

import (
	"testing"
)

func TestLevelForXP(t *testing.T) {
	tests := []struct {
		name          string
		xp            float64
		expectedLevel int
	}{
		{"Zero XP", 0, 1},
		{"Just below level 2", 82, 1},
		{"Exactly level 2", 83, 2},
		{"Just above level 2", 84, 2},
		{"Level 50", 101333, 50},
		{"Level 70", 737627, 70},
		{"Level 77", 1475581, 77},
		{"Level 90", 5346332, 90},
		{"Exactly level 99", 13034431, 99},
		{"Above level 99", 20000000, 99}, // Should cap at 99
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LevelForXP(tt.xp)
			if result != tt.expectedLevel {
				t.Errorf("LevelForXP(%f) = %d, want %d", tt.xp, result, tt.expectedLevel)
			}
		})
	}
}

func TestXPRequired(t *testing.T) {
	tests := []struct {
		name         string
		currentLevel int
		targetLevel  int
		expectedXP   float64
		expectError  bool
	}{
		{
			name:         "Level 1 to 2",
			currentLevel: 1,
			targetLevel:  2,
			expectedXP:   83, // Level 2 requires 83 XP total, level 1 is 0, so 83-0=83
			expectError:  false,
		},
		{
			name:         "Level 50 to 70",
			currentLevel: 50,
			targetLevel:  70,
			expectedXP:   737627 - 101333, // Level 70 XP - Level 50 XP
			expectError:  false,
		},
		{
			name:         "Level 77 to 99",
			currentLevel: 77,
			targetLevel:  99,
			expectedXP:   13034431 - 1475581, // Level 99 XP - Level 77 XP
			expectError:  false,
		},
		{
			name:         "Level 98 to 99",
			currentLevel: 98,
			targetLevel:  99,
			expectedXP:   13034431 - 11805606, // Should be exactly the XP difference
			expectError:  false,
		},
		{
			name:         "Invalid: current level 0",
			currentLevel: 0,
			targetLevel:  50,
			expectError:  true,
		},
		{
			name:         "Invalid: target level 100",
			currentLevel: 50,
			targetLevel:  100,
			expectError:  true,
		},
		{
			name:         "Invalid: current >= target",
			currentLevel: 70,
			targetLevel:  70,
			expectError:  true,
		},
		{
			name:         "Invalid: current > target",
			currentLevel: 80,
			targetLevel:  70,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := XPRequired(tt.currentLevel, tt.targetLevel)

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

			if result != tt.expectedXP {
				t.Errorf("XPRequired(%d, %d) = %f, want %f",
					tt.currentLevel, tt.targetLevel, result, tt.expectedXP)
			}

			// Additional validation: result should always be positive for valid inputs
			if result <= 0 {
				t.Errorf("XP required should be positive, got %f", result)
			}
		})
	}
}

func TestXPTableConsistency(t *testing.T) {
	// Test that the XP table is consistently increasing
	for level := 2; level < len(XPTable); level++ {
		if XPTable[level] <= XPTable[level-1] {
			t.Errorf("XP table should be increasing: level %d (%f) should be > level %d (%f)",
				level+1, XPTable[level], level, XPTable[level-1])
		}
	}

	// Test that level 1 starts at 0 XP
	if XPTable[0] != 0 {
		t.Errorf("Level 1 should start at 0 XP, got %f", XPTable[0])
	}

	// Test known XP values
	knownValues := map[int]float64{
		2:  83,
		10: 1154,
		50: 101333,
		99: 13034431,
	}

	for level, expectedXP := range knownValues {
		actualXP := XPTable[level-1] // Array is 0-indexed, levels are 1-indexed
		if actualXP != expectedXP {
			t.Errorf("Level %d should require %f XP, got %f", level, expectedXP, actualXP)
		}
	}
}

func TestLevelForXPAndXPRequiredConsistency(t *testing.T) {
	// Test that LevelForXP and XPRequired are consistent with each other
	testCases := []int{1, 25, 50, 77, 90, 99}

	for _, level := range testCases {
		if level >= len(XPTable) {
			continue
		}

		xpForLevel := XPTable[level-1] // Convert 1-indexed level to 0-indexed array
		calculatedLevel := LevelForXP(xpForLevel)

		if calculatedLevel != level {
			t.Errorf("Consistency check failed for level %d: LevelForXP(%f) = %d, want %d",
				level, xpForLevel, calculatedLevel, level)
		}

		// Test XP just below the threshold
		if xpForLevel > 1 {
			calculatedLevelBelow := LevelForXP(xpForLevel - 1)
			if calculatedLevelBelow >= level {
				t.Errorf("Level calculation incorrect: LevelForXP(%f) = %d should be < %d",
					xpForLevel-1, calculatedLevelBelow, level)
			}
		}
	}
}

// Test boundary conditions
func TestXPTableBoundaryConditions(t *testing.T) {
	// Test very small XP values
	smallValues := []float64{0, 0.1, 0.9, 1.0}
	for _, xp := range smallValues {
		level := LevelForXP(xp)
		if level != 1 {
			t.Errorf("Very small XP values should give level 1: LevelForXP(%f) = %d", xp, level)
		}
	}

	// Test very large XP values (should cap at level 99)
	largeValues := []float64{20000000, 100000000, 1000000000}
	for _, xp := range largeValues {
		level := LevelForXP(xp)
		if level != 99 {
			t.Errorf("Very large XP values should cap at level 99: LevelForXP(%f) = %d", xp, level)
		}
	}

	// Test exact boundary values
	for level := 2; level <= 99; level++ {
		xpThreshold := XPTable[level-1]

		// Test exactly at threshold
		calculatedLevel := LevelForXP(xpThreshold)
		if calculatedLevel != level {
			t.Errorf("Exact threshold test failed: LevelForXP(%f) = %d, want %d",
				xpThreshold, calculatedLevel, level)
		}

		// Test just below threshold
		if xpThreshold > 1 {
			calculatedLevelBelow := LevelForXP(xpThreshold - 0.1)
			if calculatedLevelBelow >= level {
				t.Errorf("Below threshold test failed: LevelForXP(%f) = %d should be < %d",
					xpThreshold-0.1, calculatedLevelBelow, level)
			}
		}
	}
}

// Benchmark tests
func BenchmarkLevelForXP(b *testing.B) {
	xpValues := []float64{100000, 500000, 1000000, 5000000, 10000000}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, xp := range xpValues {
			LevelForXP(xp)
		}
	}
}

func BenchmarkXPRequired(b *testing.B) {
	testCases := [][2]int{{1, 50}, {50, 77}, {77, 99}, {90, 99}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			XPRequired(tc[0], tc[1])
		}
	}
}
