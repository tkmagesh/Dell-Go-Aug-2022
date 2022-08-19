package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mutex sync.Mutex

var counter int

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn()
	}
	wg.Wait()
	fmt.Println("counter = ", counter)
}

func fn() {
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
	wg.Done()
}
