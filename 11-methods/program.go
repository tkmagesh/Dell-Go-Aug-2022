package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("id=%d name=%q, cost=%v, units=%d, category=%q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discountPercentage int) {
	p.Cost = p.Cost * ((100 - float32(discountPercentage)) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

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

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, expiry=%q", pp.Product.Format(), pp.Expiry)
}

/* func (pp *PerishableProduct) ApplyDiscount(discountPercentage float32) {
	//pp.Product.Cost = pp.Product.Cost * ((100-discountPercentage)/100)
	pp.Cost = pp.Cost * ((100 - discountPercentage) / 100)
} */

func main() {

	p := Product{Id: 200, Name: "Pen", Cost: 5, Units: 10, Category: "Stationary"}
	fmt.Println(p.Format())
	//ApplyDiscount(&p, 10)
	//(&p).ApplyDiscount(10)
	p.ApplyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(p.Format())

	fmt.Println()
	grapes := NewPerishableProduct(100, "Grapes", 50, 2, "Fruits", "2 Days")
	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10) // invoking the "ApplyDiscount" method of the base "Product" type
	fmt.Println("After applying 10% discount")
	fmt.Println(grapes.Format())
}
