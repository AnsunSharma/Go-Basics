package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("This loop will run indefinitely unless terminated explicitly")
	// }

	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
