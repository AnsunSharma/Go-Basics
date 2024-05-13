package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

var emp = Employee{
	Name: "Alice",
	Age:  35,
	Address: Address{
		Street: "123 Main St",
		City:   "Anytown",
	},
}

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(emp, emp.Address) // prints "{Alice 35 {123 Main St Anytown}}"

	emp2 := new(Employee)
	*emp2 = emp       // make a copy of the value in 'emp'
	fmt.Println(emp2) // prints "{Alice

	v := Vertex{1, 2}
	p := &v
	p.X = 10
	fmt.Println(v)
}
