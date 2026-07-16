package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (m mockDB) Ping() error {
	return nil
}


func TestHealthHandler(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	recorder := httptest.NewRecorder()

	healthHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf(
			"expected status 200, got %d",
			recorder.Code,
		)
	}
}
