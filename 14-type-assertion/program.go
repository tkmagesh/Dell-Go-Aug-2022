package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	//var x interface{}
	var x any
	x = 100
	x = "This is a string"
	x = 12.098
	x = true
	fmt.Println(x)

	var no interface{}
	no = "Magesh"
	if val, ok := no.(int); ok {
		result := val + 200
		fmt.Println(result)
	} else {
		fmt.Println("no is not an int")
	}

	//no = 100
	//no = true
	//no = struct{}{}
	//no = Product{Cost: 100}
	no = "Magesh"
	switch val := no.(type) {
	case int:
		fmt.Println("no is an int, no + 200 = ", val+200)
	case string:
		fmt.Println("no is a string, len(no) = ", len(val))
	case bool:
		fmt.Println("no is a bool, !no = ", !val)
	case struct{}:
		fmt.Println("no is an empty struct")
	case Product:
		fmt.Println("no is a product, no.Cost = ", val.Cost)
	default:
		fmt.Println("no is not a known type")
	}

}
