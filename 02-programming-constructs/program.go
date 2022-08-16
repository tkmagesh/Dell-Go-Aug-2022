package main

import "fmt"

func main() {

	fmt.Println("if else construct")
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { // the scope of 'no' variable is limited to the 'if else' block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	fmt.Printf("\nswitch case construct \n")
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	/* score := 11
	switch score {
	case 0:
		fmt.Println("Terrible")
	case 1:
		fmt.Println("Terrible")
	case 2:
		fmt.Println("Terrible")
	case 3:
		fmt.Println("Terrible")
	case 4:
		fmt.Println("Not Bad")
	case 5:
		fmt.Println("Not Bad")
	case 6:
		fmt.Println("Not Bad")
	case 7:
		fmt.Println("Not Bad")
	case 8:
		fmt.Println("Very Good")
	case 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	} */

	/*
		score := 6
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not Bad")
		case 8, 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/

	/*
		x := 21
		switch {
		case x%2 == 0:
			fmt.Printf("%d is an even number\n", x)
		case x%2 == 1:
			fmt.Printf("%d is an odd number\n", x)
		}
	*/

	switch x := 21; {
	case x%2 == 0:
		fmt.Printf("%d is an even number\n", x)
	case x%2 == 1:
		fmt.Printf("%d is an odd number\n", x)
	}

	fmt.Println("Switch case with fallthrough")
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	plan := "pro"
	switch plan {
	case "premium":
		fmt.Println("All premium features")
		fallthrough
	case "super":
		fmt.Println("All super features")
		fallthrough
	case "pro":
		fmt.Println("All pro features")
		fallthrough
	case "free":
		fmt.Println("All free features")
		fallthrough
	default:
		fmt.Println("Trial for 30 days")
	}

	fmt.Printf("\nfor construct\n")
	fmt.Println("for v1.0")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\nfor v2.0 (like a while loop)\n")
	no := 1
	for no < 100 {
		no += no
	}
	fmt.Printf("no = %d\n", no)

	fmt.Printf("\nfor v3.0 (infinite loop)\n")
	numSum := 1
	for {
		numSum += numSum
		if numSum > 100 {
			break
		}
	}
	fmt.Printf("numSum = %d\n", numSum)

	fmt.Printf("\nusing labels\n")
OUTERLOOP:
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			fmt.Println(i, j)
			if i == j {
				continue OUTERLOOP
			}
		}
	}

}
