package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // handing over the f1() execution to the scheduler
	f2()
	time.Sleep(1 * time.Second)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
