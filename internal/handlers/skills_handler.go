package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"osrs-xp-kits/internal/calculators/skills"
	"path/filepath"
	"strings"
)

func SkillDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 || pathParts[0] != "api" || pathParts[1] != "skill-data" {
		http.Error(w, "Invalid API path structure. Expected /api/skill-data/{skillName}", http.StatusBadRequest)
		return
	}
	skillName := strings.ToLower(pathParts[len(pathParts)-1])

	if skillName == "" {
		http.Error(w, "Skill name not provided in path", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("internal", "calculators", "skills", "json", fmt.Sprintf("%s.json", skillName))

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, fmt.Sprintf("Skill data not found for: %s", skillName), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error reading skill data file: %v", err), http.StatusInternalServerError)
		return
	}

	var skillData skills.SkillData
	if err := json.Unmarshal(jsonData, &skillData); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing skill data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(skillData); err != nil {
		fmt.Printf("Error encoding skill data to response: %v\n", err)
	}
}
