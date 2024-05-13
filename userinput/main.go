package main

import (
	"bufio"
	"fmt"
	"os"
)

var a string

func main() {
	fmt.Println("Enter your name: ")
	// fmt.Scan(&a)
	// fmt.Println("hello mr ", a)
	reader := bufio.NewReader(os.Stdin)
	a, _ = reader.ReadString('\n') // read until the newline character '\n
	fmt.Println("hello mr ", a, "nice to meet you!")
}
