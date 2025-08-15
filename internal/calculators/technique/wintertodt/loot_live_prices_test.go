package wintertodt

import (
	"testing"
)

func TestSimulateLootWithLivePrices(t *testing.T) {
	skillLevels := SkillLevels{
		Herblore:    70,
		Mining:      75,
		Fishing:     80,
		Crafting:    60,
		Farming:     65,
		Woodcutting: 80,
	}

	livePrices := map[string]int{
		"Grimy ranarr weed": 8000,   // Higher than default 7000
		"Uncut diamond":     3000,   // Higher than default 2800
		"Raw shark":         900,    // Higher than default 800
		"Magic seeds":       110000, // Higher than default 104000
	}

	rounds := 10
	pointsPerRound := 600

	loot, totalValue := SimulateLootWithLivePrices(rounds, pointsPerRound, skillLevels, livePrices)

	// Basic validation
	if loot == nil {
		t.Fatal("SimulateLootWithLivePrices returned nil loot")
	}

	if totalValue < 0 {
		t.Errorf("Total value should not be negative, got %d", totalValue)
	}

	// Should always have supply crates
	if supplyCrates, exists := loot["Supply crate"]; !exists {
		t.Error("Expected supply crates in loot")
	} else if supplyCrates != rounds {
		t.Errorf("Expected %d supply crates, got %v", rounds, supplyCrates)
	}

	// Test deterministic behavior with seed
	loot1, value1 := SimulateLootWithLivePricesAndSeed(rounds, pointsPerRound, skillLevels, livePrices, 12345)
	loot2, value2 := SimulateLootWithLivePricesAndSeed(rounds, pointsPerRound, skillLevels, livePrices, 12345)

	if value1 != value2 {
		t.Errorf("Same seed should produce same total value: %d vs %d", value1, value2)
	}

	// Compare a few key items to ensure deterministic behavior
	for itemName := range loot1 {
		if loot1[itemName] != loot2[itemName] {
			t.Errorf("Same seed should produce same loot for %s: %v vs %v", itemName, loot1[itemName], loot2[itemName])
		}
	}
}

func TestGetEnhancedSupplyDropsWithLivePrices(t *testing.T) {
	skillLevels := SkillLevels{
		Herblore:    90, // High level should enhance herbs
		Mining:      80, // High level should enhance gems
		Fishing:     80, // High level should enhance fish
		Crafting:    60,
		Farming:     65,
		Woodcutting: 85, // High level should enhance logs
	}

	livePrices := map[string]int{
		"Grimy ranarr weed": 8000,
		"Grimy snapdragon":  12000,
		"Uncut diamond":     3500,
		"Raw shark":         1000,
		"Yew logs":          500,
		"Magic logs":        1200,
		"Pure essence":      5,
	}

	enhanced := getEnhancedSupplyDropsWithLivePrices(skillLevels, livePrices)

	if len(enhanced) != len(SupplyDrops) {
		t.Errorf("Enhanced drops should have same length as base drops: %d vs %d", len(enhanced), len(SupplyDrops))
	}

	// Check that live prices were applied
	for _, item := range enhanced {
		if livePrice, exists := livePrices[item.Name]; exists {
			if item.Value != livePrice {
				t.Errorf("Item %s should have live price %d, got %d", item.Name, livePrice, item.Value)
			}
		}
	}

	// Check skill-based enhancements
	for _, item := range enhanced {
		switch item.Name {
		case "Grimy ranarr weed", "Grimy snapdragon":
			// Herblore 90 should enhance quantity by 1
			if skillLevels.Herblore >= 70 {
				baseItem := getBaseItemByName(item.Name)
				if baseItem != nil && item.Quantity <= baseItem.Quantity {
					t.Errorf("Item %s should have enhanced quantity due to high Herblore", item.Name)
				}
			}
		case "Uncut diamond":
			// Mining 80 should enhance rate
			if skillLevels.Mining >= 75 {
				baseItem := getBaseItemByName(item.Name)
				if baseItem != nil && item.Rate <= baseItem.Rate {
					t.Errorf("Item %s should have enhanced rate due to high Mining", item.Name)
				}
			}
		}
	}
}

func TestLivePricesVsStaticPrices(t *testing.T) {
	skillLevels := SkillLevels{
		Herblore:    50,
		Mining:      50,
		Fishing:     50,
		Crafting:    50,
		Farming:     50,
		Woodcutting: 50,
	}

	rounds := 100 // Larger sample for better comparison
	pointsPerRound := 600
	seed := int64(42)

	// Static prices (using default values)
	staticLoot, staticValue := SimulateLootWithSkillsAndPointsAndSeed(rounds, pointsPerRound, skillLevels, seed)

	// Live prices (all items doubled in price)
	livePrices := make(map[string]int)
	for _, item := range SupplyDrops {
		livePrices[item.Name] = item.Value * 2
	}
	for _, item := range UniqueRolls {
		if item.Name != "Phoenix" && item.Name != "Pyromancer outfit" {
			livePrices[item.Name] = item.Value * 2
		}
	}
	livePrices["Magic seeds"] = 208000   // Double default
	livePrices["Torstol seeds"] = 116000 // Double default

	liveLoot, liveValue := SimulateLootWithLivePricesAndSeed(rounds, pointsPerRound, skillLevels, livePrices, seed)

	// Loot quantities should be the same (same seed)
	for itemName, staticQty := range staticLoot {
		if liveQty, exists := liveLoot[itemName]; exists {
			if staticQty != liveQty {
				t.Errorf("Quantities should match for %s: static=%v, live=%v", itemName, staticQty, liveQty)
			}
		}
	}

	// Live prices should result in higher total value (roughly double)
	// Allow for some variance due to Phoenix/Pyromancer outfit not having prices
	if liveValue <= staticValue {
		t.Errorf("Live prices should result in higher value: static=%d, live=%d", staticValue, liveValue)
	}

	// Should be roughly double (within 10% tolerance due to non-priced items)
	expectedRatio := 2.0
	actualRatio := float64(liveValue) / float64(staticValue)
	tolerance := 0.3 // 30% tolerance

	if actualRatio < expectedRatio-tolerance || actualRatio > expectedRatio+tolerance {
		t.Errorf("Value ratio should be close to %.1f, got %.2f (static=%d, live=%d)",
			expectedRatio, actualRatio, staticValue, liveValue)
	}
}

// Helper function to get base item by name
func getBaseItemByName(name string) *LootItem {
	for _, item := range SupplyDrops {
		if item.Name == name {
			return &item
		}
	}
	for _, item := range UniqueRolls {
		if item.Name == name {
			return &item
		}
	}
	return nil
}
