package main

import "fmt"

func main() {
	fmt.Println("welcome to error handling")
	anwer, _ := devide(10, 0)
	fmt.Println(anwer)
}

func devide(a, b float64) (float64, error) {
	var ans float64 = a / b
	if b == 0 {
		return 0, fmt.Errorf("denomiator must  not be zero") // return an error with message
	}
	return ans, nil
}
