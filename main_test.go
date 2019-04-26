package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/json"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestGetRoot(t *testing.T) {
	w := performRequest(GetMainEngine(), "GET", "/")

	if w.Code != http.StatusOK {
		t.Errorf("Non-expected status code%v:\n\tbody: %v", http.StatusOK, w.Code)
	}
}

func TestGetPing(t *testing.T) {
	w := performRequest(GetMainEngine(), "GET", "/ping")

	if w.Code != http.StatusOK {
		t.Errorf("Non-expected status code%v:\n\tbody: %v", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]
	
	if err != nil {
		t.Errorf("Transfer format error")
	}
	if exists == false {
		t.Errorf("Non-expected empty key")
	}
	if value != "pong" {
		t.Errorf("Non-expected wrong value")
	}
}