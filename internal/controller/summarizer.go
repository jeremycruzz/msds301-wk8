package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jeremycruzz/msds301-wk8/pkg/summarizer"
)

type SummarizerController struct {
	SummarizerService *summarizer.Service
}

type SummarizeResponse struct {
	Summary string `json:"summary"`
	Eli5    string `json:"eli5"`
	Error   string `json:"error,omitempty"`
}

func NewSummarizerController(service *summarizer.Service) *SummarizerController {
	return &SummarizerController{SummarizerService: service}
}

// Get summary gets a summary and eli5 for a given topic
func (sc *SummarizerController) GetSummary(w http.ResponseWriter, r *http.Request) {
	topic := strings.ToLower(chi.URLParam(r, "topic"))
	if topic == "" {
		http.Error(w, "Topic is required", http.StatusBadRequest)
		return
	}

	summary, eli5, err := sc.SummarizerService.Summarize(topic)

	if err != nil {
		respondJSON(w, SummarizeResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	respondJSON(w, SummarizeResponse{Summary: summary, Eli5: eli5}, http.StatusOK)
}

// Get summary gets new summary and new eli5 for a given topic that already exists
func (sc *SummarizerController) UpdateSummary(w http.ResponseWriter, r *http.Request) {
	topic := strings.ToLower(chi.URLParam(r, "topic"))
	if topic == "" {
		http.Error(w, "Topic is required", http.StatusBadRequest)
		return
	}

	summary, eli5, err := sc.SummarizerService.UpdateSummary(topic)

	if err != nil {
		respondJSON(w, SummarizeResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	respondJSON(w, SummarizeResponse{Summary: summary, Eli5: eli5}, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
