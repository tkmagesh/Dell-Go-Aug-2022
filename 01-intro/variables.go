package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := syntax is applicable only in function scope (not in package scope)
	fmt.Println(msg)

	/*
		var x int
		var y int
		var result int
		var str string
		str = "add result is "
		x = 100
		y = 200
		result = x + y
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "add result is "
		var result int = x + y
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "add result is"
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		result = x + y
		str = "add result is"
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "add result is"
			result int    = x + y
		)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "add result is"
			result = x + y
		)
	*/

	/*
		var (
			x, y, str = 100, 200, "add result is"
			result    = x + y
		)
	*/

	x, y, str := 100, 200, "add result is"
	result := x + y

	fmt.Println(str, result)

}
