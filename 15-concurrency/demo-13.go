package main

import (
	"fmt"
	"time"
)

func main() {

	ch := add(100, 200)
	//ch <- 100
	result := <-ch
	fmt.Println(result)
}

func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
