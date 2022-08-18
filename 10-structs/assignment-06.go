/*
Create the following functions for the Product
	Format() => return a string representation for the given product
		ex: 'Id=100, Name="Grapes", Cost=50, Units=2, Category="Fruits"
	ApplyDiscount()
		to apply a discount to the given product

Create the above functions for PerishableProduct as well

*/

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

	p := Product{Id: 200, Name: "Pen", Cost: 5, Units: 10, Category: "Stationary"}
	fmt.Println(Format(p))
	ApplyDiscount(&p, 10)
	fmt.Println("After applying 10% discount")
	fmt.Println(Format(p))

	fmt.Println(FormatPP(grapes))
	ApplyDiscountPP(&grapes, 10)
	fmt.Println("After applying 10% discount")
	fmt.Println(FormatPP(grapes))
}

func Format(p Product) string {
	return fmt.Sprintf("id=%d name=%q, cost=%v, units=%d, category=%q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func FormatPP(pp PerishableProduct) string {
	return fmt.Sprintf("%s, expiry=%q", Format(pp.Product), pp.Expiry)
}

func ApplyDiscountPP(pp *PerishableProduct, discountPercentage float32) {
	//pp.Product.Cost = pp.Product.Cost * ((100-discountPercentage)/100)
	pp.Cost = pp.Cost * ((100 - discountPercentage) / 100)
}
