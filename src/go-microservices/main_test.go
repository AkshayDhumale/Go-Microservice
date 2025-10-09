package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	health(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var resp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if resp["status"] != "Ok" {
		t.Errorf("expected status 'Ok', got '%s'", resp["status"])
	}

	if _, ok := resp["timestamp"]; !ok {
		t.Error("expected timestamp in response")
	}
}
