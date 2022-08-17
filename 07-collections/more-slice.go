package main

import "fmt"

func main() {
	//var nos []int
	var nos []int = make([]int, 0, 10)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 10)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 20)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 30)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 40)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 50)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	/*
		slicing
		[lo : hi] => from lo to hi-1
		[:hi] => from 0 to hi-1
		[lo:] => from lo to end of the slice
	*/

	fmt.Println("nos[2:4] => ", nos[2:4])
	fmt.Println("nos[:4] => ", nos[:4])
	fmt.Println("nos[2:] => ", nos[2:])

	subset := nos[2:4]
	subset[0] = 100

	fmt.Println("subset = ", subset)
	fmt.Println("nos = ", nos)

}
