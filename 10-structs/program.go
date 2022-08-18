package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	City   string
	Salary float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", "Bengaluru", 10000}
	//var emp Employee = Employee{Id: 100, Name: "Magesh", City: "Bengaluru"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		}
	*/
	/*
		emp := Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		}
	*/
	/*
		emp := &Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		}
	*/
	//emp := new(Employee)
	//emp := &Employee{}

	emp := &Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
	}
	fmt.Printf("%#v\n", emp)

	fmt.Println("employee name :", emp.Name)
	//fmt.Println("employee name :", (*emp).Name)
}
