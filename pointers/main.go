package main

import "fmt"

func main() {
	num := 5
	modifyValue(num)
	modifyValueRef(&num)
}

func modifyValue(ptr int) {
	fmt.Println(ptr)
}

func modifyValueRef(ptr *int) {
	*ptr = 10
	nums := *ptr
	fmt.Println(*ptr, nums)
}
