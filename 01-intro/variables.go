package main

import "fmt"

var myVar = 100

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

	x, y, str := 100, 200, "add result of "
	result := x + y

	//fmt.Println(str, result)
	//fmt.Printf("%s %d[of type %T] and %d[of type %T] is %d[of type %T]\n", str, x, x, y, y, result, result)

	s := fmt.Sprintf("%s %d[of type %T] and %d[of type %T] is %d[of type %T]\n", str, x, x, y, y, result, result)
	fmt.Println(s)

	/*
		var myVar = 100
		myVar = 200
		fmt.Println(myVar)
	*/

	//constants
	const pi = 3.14

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	const (
		red = iota * 2
		green
		_
		blue
	)

	fmt.Printf("red = %d, green = %d, blue = %d\n", red, green, blue)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)

}
