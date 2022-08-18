package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	City   string
	Salary float32
	Format func() string
}

func NewEmployee(id int, name string, city string, salary float32) Employee {
	return Employee{
		Id:     id,
		Name:   name,
		City:   city,
		Salary: salary,
		Format: func() string {
			return fmt.Sprintf("Id = %d, Name = %q, city = %q, salary = %v", id, name, city, salary)
		},
	}
}

func main() {
	emp := NewEmployee(100, "Magesh", "Bengaluru", 1000)
	fmt.Println(emp.Format())
	emp.City = "Pune"
	fmt.Println(emp.Format())
}
