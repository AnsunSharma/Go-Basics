package main

import (
	"fmt"
	"my_learning/myutil"
	"my_learning/variables"
	"reflect"
)

func main() {
	age := 2
	fmt.Println("Hello World!", reflect.TypeOf(age))
	fmt.Printf("age is typeof  %T \n", reflect.TypeOf(age))
	myutil.PrintMassage("This is a massage from PrintMessage function.")
	variables.PrintAllVariables()

}
