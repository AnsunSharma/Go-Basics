package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func main() {
	defer fmt.Printf("first line")
	add := sum(10, 20)
	fmt.Println("\nsecond line") // will not be printed because of defer
	defer fmt.Println(add)       // this statement is executed last and prints the result of add function
}
