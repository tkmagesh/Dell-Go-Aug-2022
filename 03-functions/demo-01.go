package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Print(greetUser("Suresh"))
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/
	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
}

//no arguments & no return values
func fn() {
	fmt.Println("fn invoked")
}

//with 1 argument & no return values
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

//with 1 argument & 1 return value
func greetUser(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

//with 2 arguments and 1 return value
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

//with more then 1 return values
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
