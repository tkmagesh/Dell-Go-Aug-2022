package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		err := recover()
		if err == DivideByZeroError {
			fmt.Println("please do not attempt to divide by 0")
			return
		}
		if err != nil {
			fmt.Println("something went wrong!")
			fmt.Println(err)
			return
		}
		fmt.Println("Thank you!")
	}()
	quotient, remainder := divide(100, 7)
	fmt.Println(quotient, remainder)
}

func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient, remainder = x/y, x%y
	return
}
