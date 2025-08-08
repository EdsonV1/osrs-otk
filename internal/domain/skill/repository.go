package skill

import "context"

// Repository defines the interface for skill data access
type Repository interface {
	GetSkillData(ctx context.Context, skillName string) (*SkillData, error)
	ListSkills(ctx context.Context) ([]string, error)
}
