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
	fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
