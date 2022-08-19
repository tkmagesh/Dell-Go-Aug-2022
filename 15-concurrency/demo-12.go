package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	time.Sleep(5 * time.Second)
	result := x + y
	ch <- result
}
