package birdhouses

import (
	"fmt"
	"osrs-xp-kits/internal/calculators/tools"
)

var NestTable = tools.DropTable{
	{Name: "Acorn", Probability: 0.211685, Price: 100},
	{Name: "Apple tree seed", Probability: 0.168152, Price: 29},
	{Name: "Willow seed", Probability: 0.133529, Price: 102},
	{Name: "Banana tree seed", Probability: 0.106826, Price: 34},
	{Name: "Orange tree seed", Probability: 0.084104, Price: 40},
	{Name: "Curry tree seed", Probability: 0.067225, Price: 53},
	{Name: "Maple seed", Probability: 0.053418, Price: 3027},
	{Name: "Pineapple seed", Probability: 0.041555, Price: 90},
	{Name: "Papaya tree seed", Probability: 0.033626, Price: 1355},
	{Name: "Yew seed", Probability: 0.026711, Price: 26217},
	{Name: "Palm tree seed", Probability: 0.021768, Price: 19772},
	{Name: "Calquat tree seed", Probability: 0.016815, Price: 130},
	{Name: "Spirit seed", Probability: 0.010882, Price: 0}, // Not sold
	{Name: "Dragonfruit tree seed", Probability: 0.005935, Price: 197931},
	{Name: "Magic seed", Probability: 0.004947, Price: 91008},
	{Name: "Teak seed", Probability: 0.003955, Price: 141},
	{Name: "Mahogany seed", Probability: 0.003955, Price: 544},
	{Name: "Celastrus seed", Probability: 0.002967, Price: 67044},
	{Name: "Redwood tree seed", Probability: 0.001979, Price: 23919},
}

func SimulateNestLoot(nests int) (map[string]map[string]int, int, error) {
	if nests < 0 {
		return nil, 0, fmt.Errorf("nests cannot be negative")
	}

	results, totalValue, err := tools.SimulateMultipleDrops(NestTable, nests)
	if err != nil {
		fmt.Printf("Error during simulation: %v\n", err)
		return nil, 0, err
	}

	return results, totalValue, nil
}
