package main

import "fmt"

func main() {
	fmt.Println("slice in Golang")
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers, len((numbers)), cap(numbers)) 
	age := make([]int, 2, 10)                          //dynamic size
	// Accessing elements of slice

	age = append(age, 7, 5, 9)
	fmt.Println("slice", age)
	fmt.Println("size", len(age), "capacity", cap(age))
}
