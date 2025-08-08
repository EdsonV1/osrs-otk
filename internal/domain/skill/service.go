package skill

import (
	"context"
	"fmt"
	"strings"
)

// Service handles skill-related business logic
type Service struct {
	repo Repository
}

// NewService creates a new skill service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetSkillData retrieves skill data by name
func (s *Service) GetSkillData(ctx context.Context, skillName string) (*SkillData, error) {
	// Normalize skill name
	normalizedName := strings.ToLower(strings.TrimSpace(skillName))
	
	if normalizedName == "" {
		return nil, fmt.Errorf("skill name cannot be empty")
	}

	skillData, err := s.repo.GetSkillData(ctx, normalizedName)
	if err != nil {
		return nil, fmt.Errorf("failed to get skill data for '%s': %w", normalizedName, err)
	}

	return skillData, nil
}

// ListSkills returns all available skills
func (s *Service) ListSkills(ctx context.Context) ([]string, error) {
	skills, err := s.repo.ListSkills(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list skills: %w", err)
	}

	return skills, nil
}