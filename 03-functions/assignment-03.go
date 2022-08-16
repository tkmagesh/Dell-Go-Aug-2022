/*
	Refactor the below solution using functions
	Every function has only one responsibility (ONLY one reason to change - SRP )
*/

package main

import "fmt"

func main() {
	var (
		menu_choice    int
		n1, n2, result float32
	)
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice :")
		fmt.Scanln(&menu_choice)
		if menu_choice == 5 {
			break
		}
		if menu_choice < 1 || menu_choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch menu_choice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result = ", result)
	}
}
