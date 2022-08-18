/*
	Refactor the below solution using functions
	Every function has only one responsibility (ONLY one reason to change - SRP )
*/

package main

import "fmt"

type Operation func(float32, float32) float32

var operations map[int]Operation = map[int]Operation{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	for {
		menu_choice := get_menu_choice()
		if menu_choice == 5 {
			break
		}
		if menu_choice < 1 || menu_choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		do_operation(menu_choice)
	}
}

func do_operation(menu_choice int) {
	var result float32
	n1, n2 := get_operands()
	operation := operations[menu_choice]
	result = operation(n1, n2)
	fmt.Println("Result = ", result)
}

func get_operands() (n1, n2 float32) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&n1, &n2)
	return
}

func get_menu_choice() (menu_choice int) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice :")
	fmt.Scanln(&menu_choice)
	return
}

func add(x, y float32) float32 {
	return x + y
}

func subtract(x, y float32) float32 {
	return x - y
}

func multiply(x, y float32) float32 {
	return x * y
}

func divide(x, y float32) float32 {
	return x / y
}
