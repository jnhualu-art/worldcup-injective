package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"worldcup-injective/internal/matches"
	"worldcup-injective/internal/predictions"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/api/health", healthHandler)
	r.Get("/api/matches", matchesHandler)
	r.Get("/api/matches/{id}", matchHandler)
	r.Get("/api/predict/{id}", freePredictHandler)       // free tier
	r.Get("/api/premium-predict/{id}", premiumPredictHandler) // gated by x402 gateway
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "service": "worldcup-injective-backend"})
}

func matchesHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, matches.All())
}

func matchHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, ok := matches.ByID(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "match not found"})
		return
	}
	writeJSON(w, http.StatusOK, m)
}

func freePredictHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, ok := matches.ByID(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "match not found"})
		return
	}
	pred := predictions.Free(m)
	writeJSON(w, http.StatusOK, pred)
}

func premiumPredictHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, ok := matches.ByID(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "match not found"})
		return
	}
	pred := predictions.Premium(m)
	writeJSON(w, http.StatusOK, pred)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
