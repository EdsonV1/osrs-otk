package server

import (
	"net/http"
	"slices"
	"strings"

	"osrs-xp-kits/internal/config"
	"osrs-xp-kits/internal/domain/skill"
	"osrs-xp-kits/internal/handlers"
	"osrs-xp-kits/internal/repository/file"
	"osrs-xp-kits/internal/services"
)

// Server represents the HTTP server
type Server struct {
	config       *config.Config
	mux          *http.ServeMux
	cacheManager *services.CacheManager
}

// New creates a new server instance
func New(cfg *config.Config) *Server {
	// Initialize OSRS API service
	osrsAPI := services.NewOSRSAPIService()

	// Initialize cache manager
	cacheManager := services.NewCacheManager("./cache", osrsAPI)
	cacheManager.StartDailyRefresh()

	s := &Server{
		config:       cfg,
		mux:          http.NewServeMux(),
		cacheManager: cacheManager,
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

	// Create API handlers for external services
	osrsAPI := services.NewOSRSAPIService()
	apiHandlers := handlers.NewAPIHandlers(osrsAPI, s.cacheManager)

	// Create enhanced handlers with live price support
	wintertodtLiveHandler := handlers.NewWintertodtLiveHandler(s.cacheManager)
	birdhouseLiveHandler := handlers.NewBirdhouseLiveHandler(s.cacheManager)

	// Health check endpoint
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"osrs-otk"}`))
	})

	// Legacy calculator handlers (keep existing functionality)
	s.mux.HandleFunc("/api/birdhouse", handlers.BirdhouseCalcHandler)
	s.mux.HandleFunc("/api/birdhouse/live", birdhouseLiveHandler.Calculate)
	s.mux.HandleFunc("/api/ardyknights", handlers.ArdyKnightCalcHandler)
	s.mux.HandleFunc("/api/wintertodt", handlers.WintertodtCalcHandler)
	s.mux.HandleFunc("/api/wintertodt/live", wintertodtLiveHandler.Calculate)
	s.mux.HandleFunc("/api/tools/gotr", handlers.GOTRCalcHandler)
	s.mux.HandleFunc("/api/tools/gotr/strategy", handlers.GOTRStrategyHandler)
	s.mux.HandleFunc("/api/tools/gotr/tips", handlers.GOTRProTipsHandler)

	// Tips endpoints for other calculators
	s.mux.HandleFunc("/api/tools/wintertodt/tips", handlers.WintertodtProTipsHandler)
	s.mux.HandleFunc("/api/tools/ardyknights/tips", handlers.ArdyKnightProTipsHandler)
	s.mux.HandleFunc("/api/tools/birdhouse/tips", handlers.BirdhouseProTipsHandler)

	// New skill data handler
	s.mux.HandleFunc("/api/skill-data/", handlers.NewSkillHandler(skillService))

	// External API endpoints
	s.mux.HandleFunc("/api/player-stats/", apiHandlers.GetPlayerStats)
	s.mux.HandleFunc("/api/prices", apiHandlers.GetCurrentPrices)
	s.mux.HandleFunc("/api/prices/refresh", apiHandlers.RefreshPrices)
	s.mux.HandleFunc("/api/cache-status", apiHandlers.GetCacheStatus)
}

// corsMiddleware adds CORS headers
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		origin := r.Header.Get("Origin")
		if slices.Contains(s.config.CORS.AllowedOrigins, "*") || slices.Contains(s.config.CORS.AllowedOrigins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
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
