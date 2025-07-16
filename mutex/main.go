package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	counter int
}

func NewCounter(startPoint int) *Counter {
	return &Counter{counter: startPoint}
}

func Count(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}
func main() {

	startPoint := 0
	counter := NewCounter(startPoint)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Count(counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.counter)
}
