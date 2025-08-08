package skills

type TrainingMethod struct {
	ID                string   `json:"id" yaml:"id"`                       // Unique identifier for this method, e.g., "canifis_rooftop"
	Name              string   `json:"name" yaml:"name"`                   // Display name, e.g., "Canifis Rooftop Course"
	LevelReq          int      `json:"levelReq" yaml:"level_required"`     // Level required
	XPRate            float64  `json:"xpRate" yaml:"xp_rate"`              // Primary XP per hour
	MarksPerHour      *float64 `json:"marksPerHour" yaml:"marks_per_hour"` // Agility specific: Marks of Grace per hour (pointer for optional)
	XPPerAction       *float64 `json:"xpPerAction" yaml:"xp_per_action"`   // Optional: XP per single action (lap, log, ore)
	ActionName        *string  `json:"actionName" yaml:"action_name"`      // Optional: Name of the action, e.g., "lap", "log"
	AlternativeXPRate []struct {
		Type string  `json:"type" yaml:"type"`
		Rate float64 `json:"rate" yaml:"rate"`
	} `json:"alternativeXpRate,omitempty" yaml:"alternative_xp_rate,omitempty"` // For methods giving XP in other skills
	Location       *string  `json:"location,omitempty" yaml:"location,omitempty"`              // In-game location
	ItemsRequired  []string `json:"itemsRequired,omitempty" yaml:"items_required,omitempty"`   // List of item names
	QuestsRequired []string `json:"questsRequired,omitempty" yaml:"quests_required,omitempty"` // List of quest names
	Notes          *string  `json:"notes,omitempty" yaml:"notes,omitempty"`                    // Additional notes or tips
	Tags           []string `json:"tags,omitempty" yaml:"tags,omitempty"`                      // For filtering/categorization
	Type           *string  `json:"type,omitempty" yaml:"type,omitempty"`                      // General category, e.g., "Rooftop Course"
}

type SkillData struct {
	SkillNameCanonical string           `json:"skillNameCanonical" yaml:"name_canonical"` // e.g., "agility", "mining" (lowercase)
	SkillNameDisplay   string           `json:"skillNameDisplay" yaml:"name_display"`     // e.g., "Agility", "Mining"
	Description        string           `json:"description,omitempty" yaml:"description,omitempty"`
	TrainingMethods    []TrainingMethod `json:"trainingMethods" yaml:"training_methods"`
}
