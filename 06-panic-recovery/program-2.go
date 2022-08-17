package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong")
			return
		}
		fmt.Println("Thank you!")
	}()

	quotient, remainder, err := divide(100, 0)
	if err != nil {
		fmt.Println("dont attempt to divide by zero")
	} else {
		fmt.Println(quotient, remainder)
	}
}

func divide(x, y int) (quotient, remainder int, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = x/y, x%y
	return
}
