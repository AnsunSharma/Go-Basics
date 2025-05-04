package variables

import (
	"fmt"
	"log"
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

	fmt.Println("go" + "lang")
	// Integers and floats.

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Booleans, with boolean operators as you’d expect.

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	const (
		localhost = 8080
	)
	log.Println(localhost)
}
