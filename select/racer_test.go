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

func TestRacerCh(t *testing.T) {
	t.Run("basic test using select, returns the url of the faster server", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		// defer makes the comman run at the end of function
		// but we can put them at the top to improve readability
		defer fastServer.Close()
		defer slowServer.Close()

		slowerURL := slowServer.URL
		fasterURL := fastServer.URL

		expected := fasterURL
		actual, err := RacerCh(slowerURL, fasterURL)

		if err != nil {
			t.Fatalf("Did not expect an error, but got one: %v", err)
		}

		if expected != actual {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
	})
	t.Run("returns an error if a server does not respond within a specific timeout duration", func(t *testing.T) {
		serverA := makeDelayServer(1 * time.Second)

		defer serverA.Close()

		_, err := RacerConfigurable(serverA.URL, serverA.URL, 50*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error, but got none")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
