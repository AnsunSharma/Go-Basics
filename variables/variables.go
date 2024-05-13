package variables

import (
	"fmt"
	"reflect"
)

func PrintAllVariables() {
	fmt.Println("print all type of variables")
	// declare and initialize a variable with the
	// var keyword, without specifying its data type

	var name string = "John Doe"
	fmt.Println(name)
	var age = 25
	fmt.Println(age)

	// you can also use the := syntax to declare and assign values in one line

	country := "USA"
	fmt.Println(reflect.TypeOf(country), country)

	// if you don't specify
	// the data type for a variable, Go will infer it from the value on the right side of the assignment operator the data type for a variable

}
