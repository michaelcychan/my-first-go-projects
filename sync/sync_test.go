package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times, returns 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("expected %d, but got %d", 3, counter.Value())
		}
	})
}
