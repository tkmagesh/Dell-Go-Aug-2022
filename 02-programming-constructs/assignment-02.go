/*
	Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

	accept the menu choice from the user
	if menu choice == 5 then exit the application
	if menu choice < 1 or > 5 then
		print "invalid choice"
		display the menu again
	if menu choice >= 1 and <= 4 then
		accept 2 numbers from the user
		perform the respective operation
		print the result
		display the menu again
*/

package main

import "fmt"

func main() {
	var menu_choice, n1, n2, result int
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
