package main

import (
	"fmt"
	"my_learning/myutil"
	"my_learning/variables"
)

func main() {
	age := 2
	fmt.Println("Hello World!")
	fmt.Printf("age is typeof  %T \n", age)
	myutil.PrintMassage("This is a massage from PrintMessage function.")
	variables.PrintAllVariables()

}
