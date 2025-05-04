package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your name: ")
	SimpleFunction()

	a := a()
	b := b()
	c := &a
	sum(a, b)
	fmt.Println(*c)
	fmt.Println(&a, &b)
	sumPointer(&a, &b)
	varedicFun(a)
	v := company()
	fmt.Println(v())
}

func SimpleFunction() any {

	name := "John Doe"
	fmt.Scanf("\nHello, %s! How are you today?", name)

	return name
}

func a() int {
	return 5
}

func b() int {
	return 10
}

func sum(a, b int) {
	fmt.Println("Sum of a and b is: ", a+b)

}

func sumPointer(a, b *int) {
	*a = *a + *b
	fmt.Println("Sum of a and b is: ", *a)

}

func varedicFun(v ...int) {
	fmt.Println(len(v))
}

func company() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

// resiver type function
