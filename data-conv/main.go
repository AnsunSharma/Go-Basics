package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	num := 10
	fmt.Printf("Original Number:%T \n", num)

	data := float64(num)
	fmt.Printf("\nAfter Conversion to Float64: %T \n", data)

	str := strconv.Itoa(int(data)) //converting float64 to string using Itoa function of str
	fmt.Printf("\nAfter Conversion to String: %T \n", str)
	// time and date conversion

	currentTime := time.Now()
	fmt.Println("the time is", currentTime)
	fmt.Printf("type of time is %T \n", currentTime) // get the year from the time object
	formated := currentTime.Format("2008/08/8 , Monday")
	fmt.Println("Formatted Time : ", formated)

}
