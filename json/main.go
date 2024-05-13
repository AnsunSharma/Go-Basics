package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Send HTTP GET request to the specified URL
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// fmt.Println(string(data))

	// Convert response data to slice of Todo structs
	var todos []Todo
	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// fmt.Println(todos)
	// Print the todos
	for _, todo := range todos {
		fmt.Println(todo.UserID, todo.Title)
	}
}
