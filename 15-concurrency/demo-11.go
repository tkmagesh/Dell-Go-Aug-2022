package main

import "fmt"

func main() {
	//non-buffered channel
	/*
		ch := make(chan int)

		go func() {
			ch <- 100
		}()

		val := <-ch

		fmt.Println(val)
	*/

	//buffered channel
	ch := make(chan int, 1)
	ch <- 100
	val := <-ch
	fmt.Println(val)
}
