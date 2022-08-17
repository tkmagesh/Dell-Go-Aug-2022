package main

import (
	"errors"
	"fmt"
)

var DivisibleByZeroError = errors.New("divisor cannot be 0")

func main() {
	q, r, err := divide(100, 0)
	if err == DivisibleByZeroError {
		fmt.Println("Please dont attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q, r)

}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	return x / y, x % y, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivisibleByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
