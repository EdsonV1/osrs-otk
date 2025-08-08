package wintertodt

type LootItem struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Value    int     `json:"value"`
	Rate     float64 `json:"rate"`
}

var CommonLoot = []LootItem{
	{"Burnt page", 1, 750, 1.0},
	{"Supply crate", 1, 0, 1.0},
}

var RareLoot = []LootItem{
	{"Torstol seeds", 1, 58000, 0.02},
	{"Magic seeds", 1, 104000, 0.015},
	{"Palm tree seeds", 1, 37000, 0.025},
	{"Yew seeds", 1, 67000, 0.02},
	{"Dragon axe", 1, 8500000, 0.0001},
	{"Phoenix", 1, 0, 0.0002},
}

var SupplyLoot = []LootItem{
	{"Grimy ranarr weed", 2, 7000, 0.1},
	{"Grimy snapdragon", 2, 11000, 0.08},
	{"Grimy torstol", 1, 25000, 0.05},
	{"Uncut diamond", 1, 2800, 0.12},
	{"Pure essence", 50, 4, 0.2},
	{"Raw shark", 3, 800, 0.15},
}

// Experience calculation constants
const (
	BaseExpPerRound = 740
	ExpMultiplier   = 13.6
	PetRatePerCrate = 0.0002 // 1/5000
)