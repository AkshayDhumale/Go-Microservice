package greetservice

import (
	"encoding/json"
	"net/http"
)

type GreetResponse struct {
	Message string `json:"message"`
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	resp := GreetResponse{Message: "Hello, " + name + "!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
