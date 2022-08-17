package main

import "fmt"

func main() {

	var x int = 100
	var xPtr *int = &x
	fmt.Println(xPtr)

	//dereferencing
	fmt.Println(*xPtr)

	//fmt.Println(x == *(&x))

	fmt.Println("Before incrementing, x = ", x)
	increment(&x)
	fmt.Println("After incrementing, x = ", x)

	n1, n2 := 10, 20
	fmt.Println("Before swapping ", n1, n2)
	swap(&n1, &n2)
	fmt.Println("After swapping ", n1, n2)
}

func increment(no *int) {
	*no++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
