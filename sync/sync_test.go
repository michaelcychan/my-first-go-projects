package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times, returns 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounterValue(t, 3, counter)
	})
	t.Run("runs safely concurrently", func(t *testing.T) {
		expectedCounter := 2000
		counter := NewCounter()

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

		assertCounterValue(t, expectedCounter, counter)
	})
}

func assertCounterValue(t testing.TB, expected int, counter *Counter) {
	if expected != counter.Value() {
		t.Errorf("expected %d, but got %d", expected, counter.Value())
	}
}

// using this function, the returned value will be a pointer of a Counter object
func NewCounter() *Counter {
	return &Counter{}
}
