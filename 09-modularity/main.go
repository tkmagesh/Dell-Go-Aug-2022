package main

import (
	"fmt"
	mathOps "my-app/calc"

	//_ "my-app/calc"
	"my-app/calc/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red("Modularity demo")
	fmt.Println(mathOps.Add(100, 200))
	fmt.Println(mathOps.Subtract(100, 200))
	fmt.Println("Operation Count :", mathOps.GetOpCount())

	fmt.Println("is 21 an even number ? :", utils.IsEven(21))
}
