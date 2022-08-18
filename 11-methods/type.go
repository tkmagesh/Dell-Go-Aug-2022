package main

import "fmt"

type MyStr string // alias for string type

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("This is a string")
	fmt.Println(str.Length())
}
