package main

import "fmt"

func main() {
	fmt.Println("lets learn arrays")
	name := [5]string{"John", "Doe"}
	var age [5]int64 // declare an array of type string with a length
	age[0] = 30

	fmt.Println(name, "the are the name ", age, len(name)) // print the entire array
	var price [3]string
	fmt.Printf("%q", price) // %q is used to display strings in quotes
}
