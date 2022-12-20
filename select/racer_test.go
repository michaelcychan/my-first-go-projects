package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)

	// defer makes the comman run at the end of function
	// but we can put them at the top to improve readability
	defer fastServer.Close()
	defer slowServer.Close()

	slowerURL := slowServer.URL
	fasterURL := fastServer.URL

	expected := fasterURL
	actual := Racer(slowerURL, fasterURL)

	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
