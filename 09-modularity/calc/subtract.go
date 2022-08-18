package calc

import "fmt"

func init() {
	fmt.Println("Calculator package initialized - subtract.go")
}

func Subtract(x, y int) int {
	opCount++
	return x - y
}
