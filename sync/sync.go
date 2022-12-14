package sync

import "sync"

type Counter struct {
	// use sync.Mutex for managing state
	// use channels when passing ownership of data
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// lock Counter so other goroutine can't run
	// Unlock Counter after current goroutine is done
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value += 1

}

func (c *Counter) Value() int {
	return c.value
}
