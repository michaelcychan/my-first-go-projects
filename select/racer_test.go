package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowerURL := slowServer.URL
	fasterURL := fastServer.URL

	expected := fasterURL
	actual := Racer(slowerURL, fasterURL)

	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}

	fastServer.Close()
	slowServer.Close()
}
