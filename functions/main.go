package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your name: ")
	SimpleFunction()
}

func SimpleFunction() any {

	name := "John Doe"
	fmt.Scanf("\nHello, %s! How are you today?", name)

	return name
}
