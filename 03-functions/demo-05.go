/* functions as return result */

package main

import (
	"fmt"
	"time"
)

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

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profiledAdd := profileOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := profileOperation(subtract)
		profiledSubtract(100, 200)
	*/

	log_profile_add := profileOperation(getLogOperation(add))
	log_profile_subtract := profileOperation(getLogOperation(subtract))

	log_profile_add(100, 200)
	log_profile_subtract(100, 200)
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		operation(x, y)
		fmt.Println("invocation completed")
	}
}

func profileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Println("Time taken : ", elapsed)
	}
}

/* func logOperation(x, y int, operation func(int, int)) {
	fmt.Println("invocation started")
	operation(x, y)
	fmt.Println("invocation completed")
} */

/* 3rd party library functions (cannot change them) */
func add(x, y int) {
	time.Sleep(3 * time.Second)
	fmt.Println("add result = ", x+y)
}

func subtract(x, y int) {
	time.Sleep(5 * time.Second)
	fmt.Println("subtract result = ", x-y)
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
