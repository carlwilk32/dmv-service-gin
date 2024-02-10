package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestByDistanceEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/distance", nil)
	rec := httptest.NewRecorder()

	ByDistance(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

// todo add valid scenario tests and mock dmv client
