package gotr

import (
	"testing"
)

// TestLootSimulatorDeterministic tests that loot simulation is deterministic with fixed seed
func TestLootSimulatorDeterministic(t *testing.T) {
	const seed = int64(12345)
	const searches = 1000

	// Create two simulators with same seed
	sim1 := NewLootSimulatorWithSeed(seed)
	sim2 := NewLootSimulatorWithSeed(seed)

	// Run simulations
	rewards1, totalValue1 := sim1.SimulateRewards(searches)
	rewards2, totalValue2 := sim2.SimulateRewards(searches)

	// Results should be identical
	if totalValue1 != totalValue2 {
		t.Errorf("Total values differ: %d vs %d", totalValue1, totalValue2)
	}

	if len(rewards1) != len(rewards2) {
		t.Errorf("Number of rewards differ: %d vs %d", len(rewards1), len(rewards2))
	}

	// Convert to map for easier comparison
	rewardMap1 := make(map[string]Reward)
	rewardMap2 := make(map[string]Reward)

	for _, r := range rewards1 {
		rewardMap1[r.Name] = r
	}
	for _, r := range rewards2 {
		rewardMap2[r.Name] = r
	}

	// Check each reward
	for name, r1 := range rewardMap1 {
		r2, exists := rewardMap2[name]
		if !exists {
			t.Errorf("Reward %s missing from second simulation", name)
			continue
		}

		if r1.Quantity != r2.Quantity {
			t.Errorf("Quantity differs for %s: %d vs %d", name, r1.Quantity, r2.Quantity)
		}
	}

	// Test with different seed to ensure results actually change
	sim3 := NewLootSimulatorWithSeed(seed + 1)
	_, totalValue3 := sim3.SimulateRewards(searches)

	if totalValue1 == totalValue3 {
		t.Errorf("Total values should differ with different seeds, but both are %d", totalValue1)
	}
}

// TestSimulateAverageRewardsDeterministic verifies the average rewards function is deterministic
func TestSimulateAverageRewardsDeterministic(t *testing.T) {
	const searches = 1000

	// Run multiple times - should always be identical since it's deterministic
	rewards1, totalValue1 := SimulateAverageRewards(searches)
	rewards2, totalValue2 := SimulateAverageRewards(searches)
	rewards3, totalValue3 := SimulateAverageRewards(searches)

	// All should be identical
	if totalValue1 != totalValue2 || totalValue1 != totalValue3 {
		t.Errorf("SimulateAverageRewards is not deterministic: %d, %d, %d", totalValue1, totalValue2, totalValue3)
	}

	if len(rewards1) != len(rewards2) || len(rewards1) != len(rewards3) {
		t.Errorf("Number of rewards not consistent: %d, %d, %d", len(rewards1), len(rewards2), len(rewards3))
	}
}
