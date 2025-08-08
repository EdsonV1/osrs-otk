package server

import (
	"net/http"
	"strings"

	"osrs-xp-kits/internal/config"
	"osrs-xp-kits/internal/domain/skill"
	"osrs-xp-kits/internal/handlers"
	"osrs-xp-kits/internal/repository/file"
)

// Server represents the HTTP server
type Server struct {
	config *config.Config
	mux    *http.ServeMux
}

// New creates a new server instance
func New(cfg *config.Config) *Server {
	s := &Server{
		config: cfg,
		mux:    http.NewServeMux(),
	}

	s.setupRoutes()
	return s
}

// Start starts the HTTP server
func (s *Server) Start() error {
	handler := s.corsMiddleware(s.mux)
	return http.ListenAndServe(":"+s.config.Server.Port, handler)
}

// GetHandler returns the HTTP handler for testing
func (s *Server) GetHandler() http.Handler {
	return s.corsMiddleware(s.mux)
}

// setupRoutes configures all application routes
func (s *Server) setupRoutes() {
	// Create repositories
	skillRepo := file.NewSkillRepository(s.config.Assets.SkillDataPath)
	
	// Create services
	skillService := skill.NewService(skillRepo)

	// Legacy calculator handlers (keep existing functionality)
	s.mux.HandleFunc("/api/birdhouse", handlers.BirdhouseCalcHandler)
	s.mux.HandleFunc("/api/ardyknights", handlers.ArdyKnightCalcHandler)
	s.mux.HandleFunc("/api/wintertodt", handlers.WintertodtCalcHandler)
	s.mux.HandleFunc("/api/tools/gotr", handlers.GOTRCalcHandler)
	s.mux.HandleFunc("/api/tools/gotr/strategy", handlers.GOTRStrategyHandler)
	
	// New skill data handler
	s.mux.HandleFunc("/api/skill-data/", handlers.NewSkillHandler(skillService))
}

// corsMiddleware adds CORS headers
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range s.config.CORS.AllowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(s.config.CORS.AllowedMethods, ", "))
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(s.config.CORS.AllowedHeaders, ", "))
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}