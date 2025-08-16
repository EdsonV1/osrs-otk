package tools

import (
	"testing"
)

// TestSimulateMultipleDropsWithSeed tests deterministic behavior with fixed seed
func TestSimulateMultipleDropsWithSeed(t *testing.T) {
	// Create a simple test drop table
	testDropTable := DropTable{
		{Name: "Test Item 1", Probability: 0.5, Price: 100},
		{Name: "Test Item 2", Probability: 0.3, Price: 200},
		{Name: "Test Item 3", Probability: 0.2, Price: 300},
	}

	const seed = int64(12345)
	const numDrops = 1000

	// Run simulation twice with same seed
	results1, totalValue1, err1 := SimulateMultipleDropsWithSeed(testDropTable, numDrops, seed)
	if err1 != nil {
		t.Fatalf("First simulation failed: %v", err1)
	}

	results2, totalValue2, err2 := SimulateMultipleDropsWithSeed(testDropTable, numDrops, seed)
	if err2 != nil {
		t.Fatalf("Second simulation failed: %v", err2)
	}

	// Results should be identical with same seed
	if totalValue1 != totalValue2 {
		t.Errorf("Total values differ: %d vs %d", totalValue1, totalValue2)
	}

	// Check that all items have same quantities
	for itemName, data1 := range results1 {
		data2, exists := results2[itemName]
		if !exists {
			t.Errorf("Item %s missing from second simulation", itemName)
			continue
		}

		if data1["quantity"] != data2["quantity"] {
			t.Errorf("Quantity differs for %s: %d vs %d", itemName, data1["quantity"], data2["quantity"])
		}

		if data1["value"] != data2["value"] {
			t.Errorf("Value differs for %s: %d vs %d", itemName, data1["value"], data2["value"])
		}
	}

	// Run with different seed to ensure it actually changes results
	_, totalValue3, err3 := SimulateMultipleDropsWithSeed(testDropTable, numDrops, seed+1)
	if err3 != nil {
		t.Fatalf("Third simulation failed: %v", err3)
	}

	// Results should be different with different seed
	if totalValue1 == totalValue3 {
		t.Errorf("Total values should differ with different seeds, but both are %d", totalValue1)
	}
}
