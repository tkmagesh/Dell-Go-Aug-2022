package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	main - deferred")
	}()
	fmt.Println("main - started")
	f1()
	fmt.Println("main - completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	f1 - deferred[1]")
		}()
		defer func() {
			fmt.Println("	f1 - deferred[2]")
		}()
	*/
	defer fmt.Println("	f1 - deferred[1]")
	defer fmt.Println("	f1 - deferred[2]")
	defer f1Deferred()

	fmt.Println("f1 - started")
	f2Result := f2()
	fmt.Println("f2 result = ", f2Result)
	fmt.Println("f1 - completed")
}

func f1Deferred() {
	fmt.Println("	f1 - deferred[3]")
}

func f2() (result int) {
	defer func() {
		fmt.Println("	f2 - deferred")
		result = 200
	}()
	fmt.Println("f2 - started")
	fmt.Println("f2 - completed")
	result = 100
	return
}
