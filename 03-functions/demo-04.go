/* functions as arguments to other functiosn */

package main

import "fmt"

func main() {
	/*
		exec(fn)
		exec(func() {
			fmt.Println("anon fn invoked")
		})
	*/

	/*
		add(100,200)
		subtract(100,200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, operation func(int, int)) {
	fmt.Println("invocation started")
	operation(x, y)
	fmt.Println("invocation completed")
}

/*
func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
}
*/

/*
func exec(f func()) {
	f()
}

func fn() {
	fmt.Println("fn invoked")
}
*/

/* 3rd party library functions (cannot change them) */
func add(x, y int) {
	fmt.Println("add result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("subtract result = ", x-y)
}
