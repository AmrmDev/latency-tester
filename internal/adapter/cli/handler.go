package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AmrmDev/latency-tester/internal/usecase"
)

type LatencyRequest struct {
	Host string `json:"host"`
}

type LatencyResponse struct {
	Host      string `json:"host"`
	Duration  string `json:"duration"`
	Timestamp string `json:"timestamp"`
}

func LatencyHandler(uc *usecase.LatencyTesterUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LatencyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Host == "" {
			http.Error(w, "Invalid request (HOST is required)", http.StatusBadRequest)
			return
		}
		result, err := uc.Execute(req.Host)
		if err != nil {
			http.Error(w, "Failed to run latency test: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := LatencyResponse{
			Host:	  result.Host,
			Duration:  result.Duration.String(),
			Timestamp: result.Timestamp.Format("2006-01-02 15:04:05"),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
