package calculators

import (
	"math"
	"testing"
)

// TestHelper provides common testing utilities for calculators
type TestHelper struct {
	t *testing.T
}

// NewTestHelper creates a new test helper
func NewTestHelper(t *testing.T) *TestHelper {
	return &TestHelper{t: t}
}

// AssertInRange checks if a value is within a specified range
func (h *TestHelper) AssertInRange(value, min, max float64, message string) {
	if value < min || value > max {
		h.t.Errorf("%s: got %f, want between %f and %f", message, value, min, max)
	}
}

// AssertIntInRange checks if an integer value is within a specified range
func (h *TestHelper) AssertIntInRange(value, min, max int, message string) {
	if value < min || value > max {
		h.t.Errorf("%s: got %d, want between %d and %d", message, value, min, max)
	}
}

// AssertPositive checks if a value is positive
func (h *TestHelper) AssertPositive(value float64, message string) {
	if value <= 0 {
		h.t.Errorf("%s should be positive, got %f", message, value)
	}
}

// AssertPositiveInt checks if an integer value is positive
func (h *TestHelper) AssertPositiveInt(value int, message string) {
	if value <= 0 {
		h.t.Errorf("%s should be positive, got %d", message, value)
	}
}

// AssertApproximatelyEqual checks if two values are approximately equal within a tolerance
func (h *TestHelper) AssertApproximatelyEqual(actual, expected, tolerance float64, message string) {
	if math.Abs(actual-expected) > tolerance {
		h.t.Errorf("%s: got %f, want %f (±%f)", message, actual, expected, tolerance)
	}
}

// AssertError checks that an error occurred
func (h *TestHelper) AssertError(err error, message string) {
	if err == nil {
		h.t.Errorf("%s: expected error but got none", message)
	}
}

// AssertNoError checks that no error occurred
func (h *TestHelper) AssertNoError(err error, message string) {
	if err != nil {
		h.t.Errorf("%s: unexpected error: %v", message, err)
	}
}

// AssertNotEmpty checks that a string is not empty
func (h *TestHelper) AssertNotEmpty(value, message string) {
	if len(value) == 0 {
		h.t.Errorf("%s should not be empty", message)
	}
}

// AssertMapNotEmpty checks that a map is not empty
func (h *TestHelper) AssertMapNotEmpty(m map[string]interface{}, message string) {
	if len(m) == 0 {
		h.t.Errorf("%s should not be empty", message)
	}
}

// Test data for common OSRS levels and XP values
var CommonLevels = map[string]int{
	"Level1":  1,
	"Level50": 101333,
	"Level70": 737627,
	"Level77": 1475581,
	"Level90": 5346332,
	"Level99": 13034431,
}

// Common test scenarios for skill calculators
type SkillTestScenario struct {
	Name          string
	CurrentLevel  int
	TargetLevel   int
	ExpectError   bool
	MinEfficiency float64
	MaxEfficiency float64
}

// GetCommonSkillTestScenarios returns common test scenarios for skill calculators
func GetCommonSkillTestScenarios() []SkillTestScenario {
	return []SkillTestScenario{
		{
			Name:          "Early game (1-50)",
			CurrentLevel:  1,
			TargetLevel:   50,
			ExpectError:   false,
			MinEfficiency: 0.1,
			MaxEfficiency: 2.0,
		},
		{
			Name:          "Mid game (50-70)",
			CurrentLevel:  50,
			TargetLevel:   70,
			ExpectError:   false,
			MinEfficiency: 0.5,
			MaxEfficiency: 3.0,
		},
		{
			Name:          "High level (70-99)",
			CurrentLevel:  70,
			TargetLevel:   99,
			ExpectError:   false,
			MinEfficiency: 1.0,
			MaxEfficiency: 5.0,
		},
		{
			Name:        "Invalid: same level",
			CurrentLevel: 70,
			TargetLevel:  70,
			ExpectError: true,
		},
		{
			Name:        "Invalid: target lower",
			CurrentLevel: 80,
			TargetLevel:  70,
			ExpectError: true,
		},
	}
}