package wintertodt

type LootItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Value    int     `json:"value"`
	Rate     float64 `json:"rate"`
}

// Unique items rolled in order (1 roll per round regardless of points)
var UniqueRolls = []LootItem{
	{"Phoenix", 1, 0, 1.0 / 5000.0},           // 1/5,000
	{"Dragon axe", 1, 8500000, 1.0 / 10000.0}, // 1/10,000
	{"Tome of fire", 1, 500000, 1.0 / 1000.0}, // 1/1,000
	{"Warm gloves", 1, 150000, 1.0 / 150.0},   // 1/150
	{"Bruma torch", 1, 40000, 1.0 / 150.0},    // 1/150
	{"Pyromancer outfit", 1, 0, 1.0 / 150.0},  // 1/150 (random piece)
	{"Burnt page", 1, 750, 1.0 / 45.0},        // 1/45
}

// Regular supply drops (based on points and rolls)
var SupplyDrops = []LootItem{
	{"Grimy ranarr weed", 2, 7000, 0.1},
	{"Grimy snapdragon", 2, 11000, 0.08},
	{"Grimy torstol", 1, 25000, 0.05},
	{"Uncut diamond", 1, 2800, 0.12},
	{"Pure essence", 50, 4, 0.2},
	{"Raw shark", 3, 800, 0.15},
	{"Yew logs", 10, 400, 0.08},
	{"Magic logs", 5, 1000, 0.05},
}

// Experience calculation constants based on OSRS Wiki
const (
	PetRatePerCrate = 0.0002 // 1/5000

	// Firemaking XP multipliers
	LightingBrazierXP    = 6.0   // 6x Firemaking level
	FeedingRootXP        = 3.0   // 3x Firemaking level
	FeedingKindlingXP    = 3.8   // 3.8x Firemaking level
	SubduingWintertodtXP = 100.0 // 100x Firemaking level (500+ points)

	// Other skill XP multipliers
	CuttingRootXP   = 0.3 // 0.3x Woodcutting level
	FletchingRootXP = 0.6 // 0.6x Fletching level
	FixingBrazierXP = 4.0 // 4x Construction level
)

// Strategy types
type Strategy string

const (
	StrategyLargeGroup Strategy = "large_group"
	StrategySolo       Strategy = "solo"
	StrategyEfficient  Strategy = "efficient"
)

// Strategy data for points and time calculations
var StrategyData = map[Strategy]struct {
	PointsPerRound      int
	MinutesPerRound     float64
	Description         string
	Requirements        []string
	XPPerRoundBase      float64 // Base XP calculation multiplier
	IncludesFletching   bool
	IncludesWoodcutting bool
}{
	StrategyLargeGroup: {
		PointsPerRound:      600,
		MinutesPerRound:     4.0, // Base time, buffer added automatically
		Description:         "Standard mass world strategy with consistent teams",
		Requirements:        []string{"Level 50 Firemaking", "Axe", "Tinderbox"},
		XPPerRoundBase:      0, // Not used anymore, calculated dynamically
		IncludesFletching:   false,
		IncludesWoodcutting: true,
	},
	StrategySolo: {
		PointsPerRound:      1000,
		MinutesPerRound:     12.0, // Solo rounds take longer
		Description:         "Solo strategy for maximum rewards and Construction XP",
		Requirements:        []string{"Level 50 Firemaking", "Warm clothing", "Food"},
		XPPerRoundBase:      0, // Not used anymore, calculated dynamically
		IncludesFletching:   true,
		IncludesWoodcutting: true,
	},
	StrategyEfficient: {
		PointsPerRound:      500,
		MinutesPerRound:     3.5, // Very fast rounds
		Description:         "Efficient team strategy focusing on speed",
		Requirements:        []string{"Level 50 Firemaking", "Good team coordination"},
		XPPerRoundBase:      0, // Not used anymore, calculated dynamically
		IncludesFletching:   false,
		IncludesWoodcutting: true,
	},
}

// Skills that affect loot quality
type SkillLevels struct {
	Herblore    int `json:"herblore"`
	Mining      int `json:"mining"`
	Fishing     int `json:"fishing"`
	Crafting    int `json:"crafting"`
	Farming     int `json:"farming"`
	Woodcutting int `json:"woodcutting"`
}
