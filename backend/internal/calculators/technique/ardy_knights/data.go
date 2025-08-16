package ardyknights

// BaseXPPerPickpocket is the XP gained from a successful pickpocket of an Ardougne Knight.
const BaseXPPerPickpocket = 84

// MinCoinDrop and MaxCoinDrop define the range of coins dropped per successful pickpocket.
const MinCoinDrop = 1
const MaxCoinDrop = 50

// StunDamage is the damage taken when failing a pickpocket.
const StunDamage = 2

// The level mapping here is based on the Ardy Knight wiki table.
var BaseSuccessChance = map[int]float64{
	55: 0.65, // Base 65%
	60: 0.70,
	65: 0.75,
	70: 0.80,
	75: 0.85,
	80: 0.90,
	85: 0.92,
	90: 0.94,
	95: 0.96,
	99: 0.97,
}

// Boosts to success chance (additive)
const (
	ArdyHardBoost     = 0.10 // 10% from Ardougne Hard Diary
	ThievingCapeBoost = 0.10 // 10% from Thieving Cape
	ShadowVeilBoost   = 0.15 // 15% from Shadow Veil spell
)

// DefaultPickpocketsPerHour is a reasonable estimate for pickpocketing speed.
// Actual speed depends heavily on clicking efficiency and stalling.
const DefaultPickpocketsPerHour = 1300

// xpForLevel stores the total experience required to reach the start of each level.
// Data from OSRS Wiki: https://oldschool.runescape.wiki/w/Thieving#Experience_table
var xpForLevel = map[int]int{
	1: 0, 2: 83, 3: 174, 4: 276, 5: 388, 6: 512, 7: 650, 8: 801, 9: 969, 10: 1154,
	11: 1358, 12: 1584, 13: 1833, 14: 2107, 15: 2411, 16: 2746, 17: 3115, 18: 3523, 19: 3973, 20: 4470,
	21: 5018, 22: 5624, 23: 6291, 24: 7028, 25: 7842, 26: 8740, 27: 9730, 28: 10824, 29: 12031, 30: 13363,
	31: 14833, 32: 16456, 33: 18247, 34: 20224, 35: 22406, 36: 24815, 37: 27473, 38: 30408, 39: 33648, 40: 37224,
	41: 41171, 42: 45529, 43: 50339, 44: 55649, 45: 61512, 46: 67983, 47: 75127, 48: 83014, 49: 91721, 50: 101333,
	51: 111945, 52: 123660, 53: 136594, 54: 150872, 55: 166636, 56: 184040, 57: 203254, 58: 224466, 59: 247886, 60: 273742,
	61: 302288, 62: 333804, 63: 368599, 64: 407015, 65: 449428, 66: 496254, 67: 547953, 68: 605032, 69: 668051, 70: 737627,
	71: 814445, 72: 899257, 73: 992895, 74: 1096278, 75: 1210421, 76: 1336443, 77: 1475581, 78: 1629200, 79: 1798808, 80: 1986068,
	81: 2192818, 82: 2421087, 83: 2673114, 84: 2951373, 85: 3258594, 86: 3597792, 87: 3972294, 88: 4385776, 89: 4842295, 90: 5346332,
	91: 5902831, 92: 6517253, 93: 7195629, 94: 7944614, 95: 8771558, 96: 9684577, 97: 10692629, 98: 11805606, 99: 13034431,
}

// GetTotalXPForLevel returns the total XP required to reach the given level.
// Returns 0 for level < 1, and 13034431 for level 99 (max XP).
func GetTotalXPForLevel(level int) int {
	if level < 1 {
		return 0
	}
	if level > 99 { // Max level in OSRS
		return xpForLevel[99] // XP for level 99 is actually total XP to reach level 99
	}
	return xpForLevel[level]
}

// GetLevelForXP returns the Thieving level corresponding to the given total XP.
func GetLevelForXP(xp int) int {
	if xp < 0 {
		return 1
	}
	for level := 99; level >= 1; level-- {
		if xp >= xpForLevel[level] {
			return level
		}
	}
	return 1
}
