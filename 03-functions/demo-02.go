/* anonymous functions */
package main

import "fmt"

func main() {
	/* anonymous functions */
	func() {
		fmt.Println("fn invoked")
	}()

	func() {
		/* with arguments */
		func(userName string) {
			fmt.Printf("Hi %s, Have a nice day!\n", userName)
		}("Magesh")

		/* with return values */
		addResult := func(x, y int) int {
			return x + y
		}(100, 200)
		fmt.Println("addResult = ", addResult)

		subResult := func(x, y int) int {
			return x - y
		}(100, 200)
		fmt.Println("subResult = ", subResult)

		/* with multiple return values */
		q, r := func(x, y int) (quotient, remainder int) {
			quotient, remainder = x/y, x%y
			return
		}(100, 7)
		fmt.Println(q, r)
	}()
}
