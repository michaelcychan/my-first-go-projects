package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times, returns 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounterValue(t, 3, counter.Value())
	})
	t.Run("runs safely concurrently", func(t *testing.T) {
		expectedCounter := 2000
		counter := Counter{}

		// sync.WaitGroup wold wait a number of goroutine to finish
		// .Add tells WaitTroup how many goroutines there are
		var wg sync.WaitGroup
		wg.Add(expectedCounter)

		for i := 0; i < expectedCounter; i++ {
			go func() {
				counter.Inc()
				// .Done() ensures WaitGroup knows this is one goroutine to perform
				wg.Done()
			}()
		}
		// .Wait() makes sure the goroutines are finished before moving on
		wg.Wait()

		assertCounterValue(t, expectedCounter, counter.Value())
	})
}

func assertCounterValue(t testing.TB, expected, got int) {
	if expected != got {
		t.Errorf("expected %d, but got %d", expected, got)
	}
}
