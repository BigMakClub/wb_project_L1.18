package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter int64
}

func NewCounter(startPoint int64) *Counter {
	return &Counter{counter: startPoint}
}

func Count(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&c.counter, 1)
}
func main() {

	startPoint := int64(0)

	c := NewCounter(startPoint)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Count(c, &wg)
	}
	wg.Wait()

	fmt.Println(c.counter)
}
