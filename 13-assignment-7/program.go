package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following

	IndexOf => return the index of the given product
		ex:  returns the index of the given product

	Includes => return true if the given product is present in the collection else return false
		ex:  returns true if the given product is present in the collection

	Filter => return a new collection of products that satisfy the given condition
		use cases:
			1. filter all costly products (cost > 1000)
				OR
			2. filter all stationary products (category = "Stationary")
				OR
			3. filter all costly stationary products
			etc

	Any => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			OR
			2. are there any stationary product (category = "Stationary")?
			OR
			3. are there any costly stationary product?
			etc

	All => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			OR
			2. are all the products stationary products (category = "Stationary")?
			OR
			3. are all the products costly stationary products?
			etc
	Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc
*/
type Products []Product

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(byName{products})
	case "Cost":
		sort.Sort(byCost{products})
	case "Units":
		sort.Sort(byUnits{products})
	}
}

//Sort implmentation
//Default implementation (by Id)
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//Sort by name
type byName struct {
	Products
}

func (productsByName byName) Less(i, j int) bool {
	return productsByName.Products[i].Name < productsByName.Products[j].Name
}

//Sort by cost
type byCost struct {
	Products
}

func (productsByCost byCost) Less(i, j int) bool {
	return productsByCost.Products[i].Cost < productsByCost.Products[j].Cost
}

//Sort by units
type byUnits struct {
	Products
}

func (productsByUnits byUnits) Less(i, j int) bool {
	return productsByUnits.Products[i].Units < productsByUnits.Products[j].Units
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Stationary"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Products List")
	fmt.Print(products)
	marker := Product{103, "Marker", 50, 20, "Stationary"}
	fmt.Println("Index of Marker = ", products.IndexOf(marker))

	fmt.Println("Costly Products")
	costlyProductPredicate := func(p Product) bool {
		return p.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Are all the products costly products ?", products.All(costlyProductPredicate))
	fmt.Println("Are there any costly products ?", products.Any(costlyProductPredicate))

	fmt.Println("Sort")
	fmt.Println("Default sort (by id)")
	//sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting by name")
	//sort.Sort(byName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting by cost")
	//sort.Sort(byCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sorting by units")
	//sort.Sort(byUnits{products})
	products.Sort("Units")
	fmt.Println(products)
}
