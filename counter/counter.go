package counter

import "sync"

// Counter implements a count
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a new pointer to a counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the counter value
func (c *Counter) Value() int {
	return c.value
}
