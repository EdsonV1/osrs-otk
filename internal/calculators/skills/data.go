package skills

type TrainingMethod struct {
	ID                string   `json:"id"`           // Unique identifier for this method, e.g., "canifis_rooftop"
	Name              string   `json:"name"`         // Display name, e.g., "Canifis Rooftop Course"
	LevelReq          int      `json:"levelReq"`     // Level required
	XPRate            float64  `json:"xpRate"`       // Primary XP per hour
	MarksPerHour      *float64 `json:"marksPerHour"` // Agility specific: Marks of Grace per hour (pointer for optional)
	XPPerAction       *float64 `json:"xpPerAction"`  // Optional: XP per single action (lap, log, ore)
	ActionName        *string  `json:"actionName"`   // Optional: Name of the action, e.g., "lap", "log"
	AlternativeXPRate []struct {
		Type string  `json:"type"`
		Rate float64 `json:"rate"`
	} `json:"alternativeXpRate,omitempty"` // For methods giving XP in other skills
	Location       *string  `json:"location,omitempty"`       // In-game location
	ItemsRequired  []string `json:"itemsRequired,omitempty"`  // List of item names
	QuestsRequired []string `json:"questsRequired,omitempty"` // List of quest names
	Notes          *string  `json:"notes,omitempty"`          // Additional notes or tips
	Tags           []string `json:"tags,omitempty"`           // For filtering/categorization
	Type           *string  `json:"type,omitempty"`           // General category, e.g., "Rooftop Course"
}

type SkillData struct {
	SkillNameCanonical string           `json:"skillNameCanonical"` // e.g., "agility", "mining" (lowercase)
	SkillNameDisplay   string           `json:"skillNameDisplay"`   // e.g., "Agility", "Mining"
	Description        string           `json:"description,omitempty"`
	TrainingMethods    []TrainingMethod `json:"trainingMethods"`
}
