package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var c Counter

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn()
	}
	wg.Wait()
	fmt.Println("counter = ", c.count)
}

func fn() {
	c.Increment()
	wg.Done()
}
