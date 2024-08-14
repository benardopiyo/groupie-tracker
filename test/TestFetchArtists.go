package test

import (
	"groupie-Tracker/pkg/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchArtists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": 1, "name": "Artist 1"}]`))
	}))
	defer server.Close()

	artists, err := api.FetchArtists()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(artists) != 1 {
		t.Fatalf("expected 1 artist, got %d", len(artists))
	}
}
