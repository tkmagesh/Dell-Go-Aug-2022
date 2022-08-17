/*
	Write a program that will generate the prime numbers between start & end
	where start & end is accepted as input from the user
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	primes := generate_primes(start, end)
	fmt.Println(primes)
}

func generate_primes(start, end int) []int {
	primes := []int{}
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
