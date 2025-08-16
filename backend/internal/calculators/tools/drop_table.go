package tools

import (
	"math/rand"
)

// DropItem represents an item that can be dropped, with a name, its probability (0.0-1.0), and its price.
type DropItem struct {
	Name        string
	Probability float64 // Probability of this item dropping (0.0 to 1.0)
	Price       int     // Price/value of the item
}

// DropTable is a collection of DropItems.
type DropTable []DropItem

// SimulateSingleDrop simulates a single drop from the table based on item probabilities.
// It takes a *rand.Rand instance to ensure proper random number generation in loops.
// It returns a pointer to the dropped DropItem, or nil if no item is dropped (shouldn't happen if probabilities sum to 1).
func (dt DropTable) SimulateSingleDrop(r *rand.Rand) *DropItem {
	roll := r.Float64() // Generate a random number between 0.0 and 1.0

	cumulativeProbability := 0.0
	for _, item := range dt {
		cumulativeProbability += item.Probability
		if roll < cumulativeProbability {
			return &item // Return a pointer to the dropped item
		}
	}
	return nil // Should not be reached if probabilities sum to 1.0
}
