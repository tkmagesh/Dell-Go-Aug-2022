package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn()
	}
	wg.Wait()
	fmt.Println("counter = ", counter)
}

func fn() {
	atomic.AddInt64(&counter, 1)
	wg.Done()
}
