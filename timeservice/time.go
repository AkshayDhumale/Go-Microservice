package timeservice

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	resp := TimeResponse{CurrentTime: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
