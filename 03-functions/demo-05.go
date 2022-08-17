/* functions as return result */

package main

import "fmt"

func main() {
	/*
		fn := getFn()
		fn()
	*/

	/*
		add(100, 200)
		subtract(100, 200)

		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	logAdd := getLogOperation(add)
	logAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	logSubtract(100, 200)
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		operation(x, y)
		fmt.Println("invocation completed")
	}
}

/* func logOperation(x, y int, operation func(int, int)) {
	fmt.Println("invocation started")
	operation(x, y)
	fmt.Println("invocation completed")
} */

/* 3rd party library functions (cannot change them) */
func add(x, y int) {
	fmt.Println("add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("subtract result = ", x-y)
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
