package tools

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// SimulateMultipleDrops simulates a given number of drops and aggregates the results.
// It returns a map of item types to their quantity and total value, and the overall total value.
// https://www.khanacademy.org/math/cc-seventh-grade-math/cc-7th-probability-statistics/cc-7th-basic-prob/v/basic-probability
func SimulateMultipleDrops(dt DropTable, numDrops int) (map[string]map[string]int, int, error) {
	return SimulateMultipleDropsWithSeed(dt, numDrops, time.Now().UnixNano())
}

// SimulateMultipleDropsWithSeed simulates drops with a specific seed for deterministic testing
func SimulateMultipleDropsWithSeed(dt DropTable, numDrops int, seed int64) (map[string]map[string]int, int, error) {
	r := rand.New(rand.NewSource(seed))

	results := map[string]map[string]int{}
	totalValue := 0

	for i := 0; i < numDrops; i++ {

		droppedItem := dt.SimulateSingleDrop(r)

		if droppedItem != nil {
			// Use the first word of the item name as the key for aggregation
			key := strings.ToLower(strings.Split(droppedItem.Name, " ")[0])

			if _, exists := results[key]; !exists {
				results[key] = map[string]int{"quantity": 0, "value": 0}
			}
			results[key]["quantity"]++
			results[key]["value"] += droppedItem.Price
			totalValue += droppedItem.Price
		} else {
			fmt.Println("Warning: SimulateSingleDrop returned nil. Check drop table probabilities.")
		}
	}

	return results, totalValue, nil
}
