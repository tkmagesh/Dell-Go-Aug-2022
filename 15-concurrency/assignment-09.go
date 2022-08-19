package main

import "fmt"

func main() {
	//receive the the prime numbers and print them
	prime_nos := generate_primes(3, 100)
	for prime_no := range prime_nos {
		fmt.Println(prime_no)
	}
}

func generate_primes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if is_prime(no) {
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func is_prime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
