package gotr

// GOTR Constants based on current game mechanics
const (
	// XP rates per reward search - updated for realistic rates
	BaseExpPerSearch = 2000.0 // Base XP at level 77 RC per search
	ExpMultiplier    = 1.8    // Multiplier for level scaling

	// Pet chance (Abyssal Protector)
	PetRatePerSearch = 1.0 / 4000.0 // 1/4000 per reward search

	// Game timing - updated for realistic game lengths
	AverageGameDuration = 10.0 // minutes per game (experienced players)
	GamesPerHour        = 6.0  // 60 / 10 = 6 games per hour

	// Reward search rates - updated for better performance
	AverageSearchesPerGame = 20.0 // searches per game (good performance)
	MinSearchesPerGame     = 15.0
	MaxSearchesPerGame     = 25.0

	// Points and reward scaling
	PointsPerSearch = 12.0 // approximate points per search
)

// Reward represents a GOTR reward item
type Reward struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Value    int     `json:"value"`
	DropRate string  `json:"drop_rate,omitempty"`
}

// RewardTable represents possible rewards from GOTR
var RewardTable = map[string]RewardItem{
	// Rune rewards (common)
	"nature_rune": {
		Name:         "Nature rune",
		BaseQuantity: 20,
		VarianceMin:  15,
		VarianceMax:  25,
		Value:        250, // current GE price
		Weight:       100,
		Category:     "runes",
	},
	"death_rune": {
		Name:         "Death rune",
		BaseQuantity: 15,
		VarianceMin:  10,
		VarianceMax:  20,
		Value:        210,
		Weight:       80,
		Category:     "runes",
	},
	"blood_rune": {
		Name:         "Blood rune",
		BaseQuantity: 10,
		VarianceMin:  8,
		VarianceMax:  15,
		Value:        385,
		Weight:       60,
		Category:     "runes",
	},
	"soul_rune": {
		Name:         "Soul rune",
		BaseQuantity: 8,
		VarianceMin:  5,
		VarianceMax:  12,
		Value:        220,
		Weight:       50,
		Category:     "runes",
	},

	// Catalysts and fragments
	"catalytic_guardian_stone": {
		Name:         "Catalytic guardian stone",
		BaseQuantity: 30,
		VarianceMin:  20,
		VarianceMax:  40,
		Value:        35,
		Weight:       120,
		Category:     "catalysts",
	},
	"elemental_guardian_stone": {
		Name:         "Elemental guardian stone",
		BaseQuantity: 25,
		VarianceMin:  20,
		VarianceMax:  35,
		Value:        28,
		Weight:       130,
		Category:     "catalysts",
	},
	"guardian_essence": {
		Name:         "Guardian essence",
		BaseQuantity: 150,
		VarianceMin:  100,
		VarianceMax:  200,
		Value:        15,
		Weight:       200,
		Category:     "essence",
	},

	// Rare rewards
	"raiments_of_the_eye_top": {
		Name:         "Raiments of the eye (top)",
		BaseQuantity: 1,
		VarianceMin:  1,
		VarianceMax:  1,
		Value:        750000,
		Weight:       2,
		Category:     "outfit",
	},
	"raiments_of_the_eye_bottom": {
		Name:         "Raiments of the eye (bottom)",
		BaseQuantity: 1,
		VarianceMin:  1,
		VarianceMax:  1,
		Value:        650000,
		Weight:       2,
		Category:     "outfit",
	},
	"hat_of_the_eye": {
		Name:         "Hat of the eye",
		BaseQuantity: 1,
		VarianceMin:  1,
		VarianceMax:  1,
		Value:        500000,
		Weight:       2,
		Category:     "outfit",
	},

	// Intrinsic rewards
	"intrinsic_catalyst": {
		Name:         "Intrinsic catalyst",
		BaseQuantity: 5,
		VarianceMin:  3,
		VarianceMax:  8,
		Value:        125,
		Weight:       40,
		Category:     "catalysts",
	},
	"lantern_lens": {
		Name:         "Lantern lens",
		BaseQuantity: 3,
		VarianceMin:  1,
		VarianceMax:  5,
		Value:        180,
		Weight:       30,
		Category:     "tools",
	},

	// Other valuable drops
	"abyssal_needle": {
		Name:         "Abyssal needle",
		BaseQuantity: 1,
		VarianceMin:  1,
		VarianceMax:  1,
		Value:        4500000, // ~4.5M gp
		Weight:       1,       // very rare
		Category:     "rare",
	},
	"abyssal_lantern": {
		Name:         "Abyssal lantern",
		BaseQuantity: 1,
		VarianceMin:  1,
		VarianceMax:  1,
		Value:        1800000, // ~1.8M gp
		Weight:       3,
		Category:     "rare",
	},
}

// RewardItem represents an item in the reward table
type RewardItem struct {
	Name         string
	BaseQuantity int
	VarianceMin  int
	VarianceMax  int
	Value        int // GP value
	Weight       int // Drop weight
	Category     string
}