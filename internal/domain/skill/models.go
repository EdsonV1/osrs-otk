package skill

// TrainingMethod represents a method for training a skill
type TrainingMethod struct {
	ID                string          `yaml:"id" json:"id"`
	Name              string          `yaml:"name" json:"name"`
	LevelRequired     int             `yaml:"level_required" json:"levelReq"`
	XPRate            int             `yaml:"xp_rate" json:"xpRate"`
	MarksPerHour      *int            `yaml:"marks_per_hour,omitempty" json:"marksPerHour"`
	XPPerAction       *float64        `yaml:"xp_per_action,omitempty" json:"xpPerAction"`
	ActionName        *string         `yaml:"action_name,omitempty" json:"actionName"`
	AlternativeXPRate []AlternativeXP `yaml:"alternative_xp_rate,omitempty" json:"alternativeXpRate"`
	Location          *string         `yaml:"location,omitempty" json:"location"`
	ItemsRequired     []string        `yaml:"items_required,omitempty" json:"itemsRequired"`
	QuestsRequired    []string        `yaml:"quests_required,omitempty" json:"questsRequired"`
	Notes             *string         `yaml:"notes,omitempty" json:"notes"`
	Tags              []string        `yaml:"tags,omitempty" json:"tags"`
	Type              *string         `yaml:"type,omitempty" json:"type"`
}

// AlternativeXP represents alternative XP gained from a method
type AlternativeXP struct {
	Type string `yaml:"type" json:"type"`
	Rate int    `yaml:"rate" json:"rate"`
}

// SkillData represents complete skill information
type SkillData struct {
	SkillNameCanonical string           `yaml:"name_canonical" json:"skillNameCanonical"`
	SkillNameDisplay   string           `yaml:"name_display" json:"skillNameDisplay"`
	Description        *string          `yaml:"description,omitempty" json:"description,omitempty"`
	TrainingMethods    []TrainingMethod `yaml:"training_methods" json:"trainingMethods"`
}
