package main

import "fmt"

func main() {
	//var productRanks map[string]int
	/*
		var productRanks map[string]int = make(map[string]int)
		productRanks["Pen"] = 2
	*/
	/*
		var productRanks map[string]int = map[string]int{}
		productRanks["Pen"] = 2
	*/

	//var productRanks map[string]int = map[string]int{"Pen": 2, "Pencil": 1, "Marker": 3}

	var productRanks map[string]int = map[string]int{
		"Pen":      2,
		"Pencil":   1,
		"Marker":   3,
		"Notebook": 4,
	}
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println()
	fmt.Println("Get the value of a key")
	fmt.Printf("Rank of Pencil is %d\n", productRanks["Pencil"])

	fmt.Println()
	fmt.Println("Check for the presence of a key")
	//keyToCheck := "Marker"
	keyToCheck := "Scribble-pad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	fmt.Println()
	keyToDelete := "Scribble-pad"
	fmt.Printf("Removing a key (%q)\n", keyToDelete)
	delete(productRanks, keyToDelete)
	fmt.Println(productRanks)

	productRanks["Marker"] = 7
	fmt.Println(productRanks)
}
