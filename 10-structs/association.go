package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	City   string
	Salary float32
	Org    *Organization
}

func main() {
	dellOrg := Organization{
		Name: "DELL",
		City: "Bengaluru",
	}

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		City:   "Bengaluru",
		Salary: 1000,
		Org:    &dellOrg,
	}

	fmt.Println(emp)

	emp2 := Employee{
		Id:     102,
		Name:   "Suresh",
		City:   "Bengaluru",
		Salary: 20000,
		Org:    &dellOrg,
	}
	fmt.Println(emp2)

	fmt.Println("Changing the city of Org to 'Pune'")
	dellOrg.City = "Pune"
	fmt.Println(emp.Org)
	fmt.Println(emp2.Org)
}
