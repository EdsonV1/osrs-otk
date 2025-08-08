package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"osrs-xp-kits/internal/domain/skill"

	"gopkg.in/yaml.v2"
)

// SkillRepository implements skill.Repository using file system
type SkillRepository struct {
	dataPath string
}

// NewSkillRepository creates a new file-based skill repository
func NewSkillRepository(dataPath string) *SkillRepository {
	return &SkillRepository{
		dataPath: dataPath,
	}
}

// GetSkillData retrieves skill data from YAML file
func (r *SkillRepository) GetSkillData(ctx context.Context, skillName string) (*skill.SkillData, error) {
	filePath := filepath.Join(r.dataPath, fmt.Sprintf("%s.yaml", skillName))
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("skill data not found for: %s", skillName)
		}
		return nil, fmt.Errorf("failed to read skill file: %w", err)
	}

	var skillData skill.SkillData
	if err := yaml.Unmarshal(data, &skillData); err != nil {
		return nil, fmt.Errorf("failed to parse skill data: %w", err)
	}

	// Ensure canonical name matches filename
	if skillData.SkillNameCanonical == "" {
		skillData.SkillNameCanonical = skillName
	}

	return &skillData, nil
}

// ListSkills returns all available skill names by scanning directory
func (r *SkillRepository) ListSkills(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(r.dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read skills directory: %w", err)
	}

	var skills []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		
		name := entry.Name()
		if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
			skillName := strings.TrimSuffix(strings.TrimSuffix(name, ".yaml"), ".yml")
			skills = append(skills, skillName)
		}
	}

	return skills, nil
}