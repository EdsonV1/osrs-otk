package handlers

import (
	"context"
	"net/http"
	"strings"

	"osrs-xp-kits/internal/domain/skill"
	"osrs-xp-kits/pkg/response"
)

// SkillHandler handles skill-related HTTP requests
type SkillHandler struct {
	skillService *skill.Service
}

// NewSkillHandler creates a new skill handler
func NewSkillHandler(skillService *skill.Service) http.HandlerFunc {
	h := &SkillHandler{
		skillService: skillService,
	}
	return h.handleSkillData
}

// handleSkillData handles requests for skill data
func (h *SkillHandler) handleSkillData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, 
			http.ErrNotSupported)
		return
	}

	// Extract skill name from path
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 || pathParts[0] != "api" || pathParts[1] != "skill-data" {
		response.Error(w, http.StatusBadRequest, 
			http.ErrMissingFile)
		return
	}
	
	skillName := pathParts[len(pathParts)-1]
	if skillName == "" {
		response.Error(w, http.StatusBadRequest, 
			http.ErrMissingFile)
		return
	}

	// Get skill data
	ctx := context.Background()
	skillData, err := h.skillService.GetSkillData(ctx, skillName)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Return skill data
	response.Success(w, skillData)
}