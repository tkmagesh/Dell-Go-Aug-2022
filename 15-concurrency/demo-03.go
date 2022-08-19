package main

import (
	"fmt"
)

func main() {
	go f1() // handing over the f1() execution to the scheduler
	f2()
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
