package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	/*
		fmt.Println("	Received data from ch : ", <-ch)
		fmt.Println("	Received data from ch : ", <-ch)
		fmt.Println("	Received data from ch : ", <-ch)
		fmt.Println("	Received data from ch : ", <-ch)
		fmt.Println("	Received data from ch : ", <-ch)
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("	Received data from ch : ", <-ch)
	}
}

func fn(ch chan int) {
	fmt.Println("Attempting to send 10")
	ch <- 10
	fmt.Println("successfully sent 10")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Attempting to send 20")
	ch <- 20
	fmt.Println("successfully sent 20")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Attempting to send 30")
	ch <- 30
	fmt.Println("successfully sent 30")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Attempting to send 40")
	ch <- 40
	fmt.Println("successfully sent 40")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Attempting to send 50")
	ch <- 50
	fmt.Println("successfully sent 50")
	time.Sleep(500 * time.Millisecond)

}
