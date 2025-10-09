package timeservice

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTimeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/time", nil)
	rr := httptest.NewRecorder()

	TimeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var resp TimeResponse
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if resp.CurrentTime == "" {
		t.Error("expected current_time in response, got empty string")
	}
}
