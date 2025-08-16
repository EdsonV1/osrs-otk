package birdhouses

import (
	"testing"
)

func TestCalculateBirdhouseData(t *testing.T) {
	tests := []struct {
		name            string
		typ             string
		quantity        int
		expectError     bool
		expectedMinXP   int
		expectedMaxXP   int
		expectedMinLoot int
		expectedMaxLoot int
	}{
		{
			name:            "Regular birdhouses - small quantity",
			typ:             "regular",
			quantity:        10,
			expectError:     false,
			expectedMinXP:   2500, // Hunter XP
			expectedMaxXP:   3000,
			expectedMinLoot: 30000,
			expectedMaxLoot: 100000,
		},
		{
			name:            "Oak birdhouses - medium quantity",
			typ:             "oak",
			quantity:        50,
			expectError:     false,
			expectedMinXP:   20000, // Hunter XP
			expectedMaxXP:   25000,
			expectedMinLoot: 250000,
			expectedMaxLoot: 1000000,
		},
		{
			name:            "Yew birdhouses - large quantity",
			typ:             "yew",
			quantity:        100,
			expectError:     false,
			expectedMinXP:   100000, // Hunter XP
			expectedMaxXP:   110000,
			expectedMinLoot: 1500000,
			expectedMaxLoot: 3000000,
		},
		{
			name:            "Magic birdhouses - high tier",
			typ:             "magic",
			quantity:        200,
			expectError:     false,
			expectedMinXP:   220000, // Hunter XP
			expectedMaxXP:   240000,
			expectedMinLoot: 3000000,
			expectedMaxLoot: 6000000,
		},
		{
			name:            "Redwood birdhouses - highest tier",
			typ:             "redwood",
			quantity:        100,
			expectError:     false,
			expectedMinXP:   115000, // Hunter XP
			expectedMaxXP:   125000,
			expectedMinLoot: 2000000,
			expectedMaxLoot: 4000000,
		},
		{
			name:        "Invalid birdhouse type",
			typ:         "invalid",
			quantity:    10,
			expectError: true,
		},
		{
			name:        "Zero quantity",
			typ:         "oak",
			quantity:    0,
			expectError: true,
		},
		{
			name:        "Negative quantity",
			typ:         "oak",
			quantity:    -10,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateBirdhouseData(tt.typ, tt.quantity)

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

			// Test Hunter XP range
			if result.HunterXP < tt.expectedMinXP || result.HunterXP > tt.expectedMaxXP {
				t.Errorf("Hunter XP out of range: got %d, want between %d and %d",
					result.HunterXP, tt.expectedMinXP, tt.expectedMaxXP)
			}

			// Test total loot value
			if result.TotalLoot < tt.expectedMinLoot || result.TotalLoot > tt.expectedMaxLoot {
				t.Errorf("Total loot out of range: got %d, want between %d and %d",
					result.TotalLoot, tt.expectedMinLoot, tt.expectedMaxLoot)
			}

			// Test that higher tier birdhouses give more XP per unit
			expectedXPPerHouse := hunterXPPerBirdhouse[tt.typ]
			if expectedXPPerHouse > 0 {
				calculatedXPPerHouse := result.HunterXP / tt.quantity
				if abs(calculatedXPPerHouse-expectedXPPerHouse) > 50 {
					t.Errorf("XP per birdhouse incorrect: got %d, expected around %d",
						calculatedXPPerHouse, expectedXPPerHouse)
				}
			}

			// Test crafting XP
			if result.CraftingXP <= 0 {
				t.Errorf("Crafting XP should be positive, got %d", result.CraftingXP)
			}

			// Test estimated nests
			if result.EstimatedNests <= 0 {
				t.Errorf("Estimated nests should be positive, got %f", result.EstimatedNests)
			}

			// Test efficiency calculations
			if result.DaysLowEff <= 0 || result.DaysMedEff <= 0 || result.DaysHighEff <= 0 {
				t.Errorf("All efficiency calculations should be positive: low=%d, med=%d, high=%d",
					result.DaysLowEff, result.DaysMedEff, result.DaysHighEff)
			}

			// Test efficiency ordering (high eff should take fewer or equal days)
			if result.DaysHighEff > result.DaysMedEff || result.DaysMedEff > result.DaysLowEff {
				t.Errorf("Efficiency ordering wrong: high=%d, med=%d, low=%d",
					result.DaysHighEff, result.DaysMedEff, result.DaysLowEff)
			}

			// Test seed drops structure
			if len(result.SeedDrops) == 0 {
				t.Errorf("Should have seed drops")
			}

			// Test that seed drops have proper structure
			for seedName, seedData := range result.SeedDrops {
				if len(seedName) == 0 {
					t.Errorf("Seed name should not be empty")
				}
				if quantity, ok := seedData["quantity"]; !ok || quantity <= 0 {
					t.Errorf("Seed %s should have positive quantity", seedName)
				}
				if value, ok := seedData["value"]; !ok {
					t.Errorf("Seed %s should have value field", seedName)
				} else if value <= 0 {
					// Some seeds like spirit seeds legitimately have 0 value - just log it
					t.Logf("Seed %s has zero value (this is expected for untradeable seeds)", seedName)
				}
			}
		})
	}
}

func TestBirdhouseTypeValidation(t *testing.T) {
	validTypes := []string{"regular", "oak", "willow", "teak", "maple", "mahogany", "yew", "magic", "redwood"}

	for _, typ := range validTypes {
		t.Run("Valid_"+typ, func(t *testing.T) {
			result, err := CalculateBirdhouseData(typ, 10)
			if err != nil {
				t.Errorf("Valid type %s should not error: %v", typ, err)
			}
			if result.HunterXP <= 0 {
				t.Errorf("Valid type %s should give positive Hunter XP", typ)
			}
		})
	}

	invalidTypes := []string{"pine", "birch", "invalid", "", "REGULAR"}
	for _, typ := range invalidTypes {
		t.Run("Invalid_"+typ, func(t *testing.T) {
			_, err := CalculateBirdhouseData(typ, 10)
			if err == nil {
				t.Errorf("Invalid type %s should error", typ)
			}
		})
	}
}

func TestBirdhouseScaling(t *testing.T) {
	// Test that higher tier birdhouses give proportionally more rewards
	types := []string{"regular", "oak", "willow", "yew", "magic", "redwood"}
	quantity := 50

	var prevHunterXP int
	var prevTotalLoot int

	for i, typ := range types {
		result, err := CalculateBirdhouseData(typ, quantity)
		if err != nil {
			t.Fatalf("Unexpected error for %s: %v", typ, err)
		}

		if i > 0 {
			if result.HunterXP <= prevHunterXP {
				t.Errorf("Higher tier %s should give more Hunter XP than previous (%d vs %d)",
					typ, result.HunterXP, prevHunterXP)
			}

			// Total loot might not always increase linearly, but should generally trend upward
			if i > 2 && result.TotalLoot < prevTotalLoot/2 {
				t.Errorf("Higher tier %s loot unexpectedly low compared to previous (%d vs %d)",
					typ, result.TotalLoot, prevTotalLoot)
			}
		}

		prevHunterXP = result.HunterXP
		prevTotalLoot = result.TotalLoot
	}
}

// Test consistency of calculations
func TestBirdhouseConsistency(t *testing.T) {
	typ := "yew"

	// Test that doubling quantity roughly doubles rewards
	result1, err1 := CalculateBirdhouseData(typ, 50)
	result2, err2 := CalculateBirdhouseData(typ, 100)

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected errors: %v, %v", err1, err2)
	}

	// Hunter XP should scale exactly
	if result2.HunterXP != result1.HunterXP*2 {
		t.Errorf("Hunter XP should scale exactly: %d * 2 != %d", result1.HunterXP, result2.HunterXP)
	}

	// Crafting XP should scale exactly
	if result2.CraftingXP != result1.CraftingXP*2 {
		t.Errorf("Crafting XP should scale exactly: %d * 2 != %d", result1.CraftingXP, result2.CraftingXP)
	}

	// Estimated nests should scale exactly
	if abs(int(result2.EstimatedNests*10)-int(result1.EstimatedNests*20)) > 1 {
		t.Errorf("Estimated nests should scale: %f * 2 ≈ %f", result1.EstimatedNests, result2.EstimatedNests)
	}
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Benchmark tests
func BenchmarkCalculateBirdhouseData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateBirdhouseData("yew", 100)
	}
}
