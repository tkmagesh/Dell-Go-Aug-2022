package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//for i := 0; i < 20; i++ {
	wg.Add(1)
	go f1() // handing over the f1() execution to the scheduler
	//}
	f2()
	wg.Wait() // wait for the counter to become 0
	fmt.Println("Exiting main")
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
