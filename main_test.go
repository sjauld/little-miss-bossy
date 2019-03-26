package main

import (
	"net/http/httptest"

	"testing"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("method", "/health", nil)

	health(w, r)

	if code := w.Result().StatusCode; code != 200 {
		t.Errorf("Expected 200, got %d", code)
	}
}

func TestTeapot(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("method", "/coffee", nil)

	teapot(w, r)

	if code := w.Result().StatusCode; code != 418 {
		t.Errorf("Expected 418, got %d", code)
	}
}
