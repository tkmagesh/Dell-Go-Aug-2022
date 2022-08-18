package calc

import "fmt"

func init() {
	fmt.Println("Calculator package initialized - add.go")
}

func Add(x, y int) int {
	opCount++
	return x + y
}
