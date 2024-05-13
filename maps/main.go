package main

import "fmt"

func main() {
	studentGrade := make(map[string]int64, 5)
	studentGrade["Alice"] = 95
	studentGrade["Bob"] = 87
	studentGrade["Charlie"] = 92
	studentGrade["Anusn"] = 100
	studentGrade["Tanbir"] = 99
	studentGrade["Diya"] = 80

	fmt.Println("map allthings", studentGrade)
	fmt.Println("map one key value: ", studentGrade["Alice"])
	delete(studentGrade, "Alice") // delete a pair by its key
	fmt.Println("after deleting Alice's grade from the map: ", studentGrade)

	for key, value := range studentGrade {
		fmt.Println(key, value)
	}

}
