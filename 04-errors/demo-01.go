/* handling errors */

package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter the number :")
	_, err := fmt.Scanf("%d\n", &no)
	if err == nil {
		fmt.Println("number = ", no)
	} else {
		fmt.Println("Please enter a number")
	}

}
