package main

import "fmt"

func main() {
	z := 10
	if z > 10 {
		fmt.Println(z, " is greater than 10")
	} else if z < 10 {
		fmt.Println(z, "is less than 10")
	} else {
		fmt.Println(z, "is equal to 10")
	}

	day := 5
	switch day {
	case 1:
		fmt.Println(day, " Monday")
		break
	case 2:
		fmt.Println(day, " Tuesday")
		break
	case 3:
		fallthrough // will also execute the next case statement
	case 4:
		fmt.Println(day, " Wednesday")
		break
	default:
		fmt.Printf("I don't know what %d is!", day)
	}

	// nested if statements
}
