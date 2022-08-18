package calc

import "fmt"

var opCount int

func init() {
	fmt.Println("Calculator package initialized - calculator.go")
}

func GetOpCount() int {
	return opCount
}
