package main

import (
	"fmt"
	"time"
)

func main() {
	prime_nos := generate_primes()
	for prime_no := range prime_nos {
		fmt.Println(prime_no)
	}
}

func generate_primes() <-chan int {
	ch := make(chan int)
	timeoutCh := time.After(5 * time.Second)
	go func() {
		no := 3
	LOOP:
		for {
			if !is_prime(no) {
				no++
				select {
				case <-timeoutCh:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case <-timeoutCh:
				break LOOP
			case ch <- no:
				no++
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */

func is_prime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
