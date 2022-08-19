package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(2 * time.Second)
		result3 := <-ch3
		fmt.Println("Data from ch3 = ", result3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case result1 := <-ch1:
			fmt.Println(result1)
		case result2 := <-ch2:
			fmt.Println(result2)
		case ch3 <- 300:
			fmt.Println("sent the data to ch3")
		default:
			fmt.Println("no channel operations were successful")
		}
	}
}
