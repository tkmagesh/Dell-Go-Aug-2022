/*
	Span 100 goroutines with ids (1 - 100) which does the following
		print "goroutine [id] started"
		wait for random (1 - 10) seconds
		print "goroutine [id] completed"

	Once all the goroutines finished their task the app should print "Done" and exit

*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	if noOfGoroutines, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Invalid arguments. Usage 'demo-23 {count}")
		return
	} else {
		for i := 1; i <= noOfGoroutines; i++ {
			wg.Add(1)
			go fn(i, wg)
		}
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("goroutine [%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	fmt.Printf("goroutine [%d] completed\n", id)
}
