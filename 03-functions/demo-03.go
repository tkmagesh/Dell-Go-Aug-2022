/* functions as values to variables */
package main

import "fmt"

var operation func(int, int) int

func main() {

	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s\n", userName)
	}
	greet("Magesh")
	greet("Suresh")

	if operation != nil {
		fmt.Println(operation(100, 200))
	} else {
		fmt.Println("undefined operation")
	}

	/* operation = func(x, y int) int {
		return x + y
	}


	operation = func(n1, n2 int) int {
		return n1 - n2
	}
	fmt.Println(operation(100, 200)) */
}
