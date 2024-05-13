package main

import "fmt"

func main() {

	var ptr *int

	var num1 int = 10
	ptr = &num1

	fmt.Println(*ptr) // Output: 10
	var num int = 10
	modifyValue(&num)
	fmt.Println(num) // Output: 20

	s := 2
	y := &s
	fmt.Println("The value of y is ", *y*4, "\n")
}

func modifyValue(ptr *int) {
	*ptr = 20
}
