package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	//Dummy
	Expiry string
}

/*
type Dummy struct {
	Cost float32
}
*/

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:       100,
				Name:     "Grapes",
				Cost:     50,
				Units:    2,
				Category: "Fruits",
			},
			Expiry: "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, 2, "Fruits", "2 Days")
	//fmt.Println(grapes.Product.Name, grapes.Product.Cost)
	fmt.Println(grapes.Name, grapes.Cost)
}
