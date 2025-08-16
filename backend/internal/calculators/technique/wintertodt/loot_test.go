package wintertodt

import (
	"testing"
)

// TestSimulateLootWithSeedDeterministic tests that Wintertodt loot simulation is deterministic with fixed seed
func TestSimulateLootWithSeedDeterministic(t *testing.T) {
	const seed = int64(12345)
	const rounds = 100

	// Run simulation twice with same seed
	loot1, totalValue1 := SimulateLootWithSeed(rounds, seed)
	loot2, totalValue2 := SimulateLootWithSeed(rounds, seed)

	// Results should be identical
	if totalValue1 != totalValue2 {
		t.Errorf("Total values differ: %d vs %d", totalValue1, totalValue2)
	}

	if len(loot1) != len(loot2) {
		t.Errorf("Number of loot items differ: %d vs %d", len(loot1), len(loot2))
	}

	// Check each loot item
	for itemName, quantity1 := range loot1 {
		quantity2, exists := loot2[itemName]
		if !exists {
			t.Errorf("Item %s missing from second simulation", itemName)
			continue
		}

		if quantity1 != quantity2 {
			t.Errorf("Quantity differs for %s: %v vs %v", itemName, quantity1, quantity2)
		}
	}

	// Test with different seed to ensure results actually change
	_, totalValue3 := SimulateLootWithSeed(rounds, seed+1)

	if totalValue1 == totalValue3 {
		t.Errorf("Total values should differ with different seeds, but both are %d", totalValue1)
	}
}
