package main

import (
	"log"
	"net/http"
	"osrs-xp-kits/internal/handlers"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := "http://localhost:5173"

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/birdhouse", handlers.BirdhouseCalcHandler)
	mux.HandleFunc("/api/ardyknights", handlers.ArdyKnightCalcHandler)
	mux.HandleFunc("/api/wintertodt", handlers.WintertodtCalcHandler)

	mux.HandleFunc("/api/skill-data/", handlers.SkillDataHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", CORSMiddleware(mux)); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
