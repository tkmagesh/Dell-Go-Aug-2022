package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{3, 1, 4, 2, 5}
	//nos := []int{3, 1, 4, 2, 5}

	nos := make([]int, 5)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	fmt.Println(nos)

	//nos = append(nos, 10)
	//nos = append(nos, 10, 20, 30)
	tens := []int{10, 20, 30, 40}
	nos = append(nos, tens...)
	fmt.Println(nos)
}
