package counter

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu  sync.Mutex
	val int
}

func New(start int) SafeCounter {
	return SafeCounter{val: start}
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.val++
	c.mu.Unlock()
}

func (c *SafeCounter) IncN(n int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.val += n
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.val
}
