package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Url = "https://jsonplaceholder.typicode.com/todos"

func PreformGetRequest() {
	res, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	var todos []Todo

	err = json.NewDecoder(res.Body).Decode(&todos)
	if err != nil {
		fmt.Println("error :-", err)
	}

	for _, todo := range todos {
		fmt.Println(todo.UserID, todo.Title)
	}
}

func PreformPostRequest() {
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "Go Lang HTTP Request Example",
		Completed: false,
	}
	payload, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("ERROR :-", err)
		return
	}
	jsonString := string(payload)
	payloadReader := strings.NewReader(jsonString)
	res, err := http.Post(Url, "application/json", payloadReader)
	if err != nil {
		fmt.Println("ERROR :-", err)
		return
	}
	fmt.Println("result", string(res.Status))
	defer res.Body.Close()
}

func PerformPutRequest() {
	// Define the URL
	putUrl := "https://jsonplaceholder.typicode.com/todos/1"
	// Create a Todo object with updated data
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "Updated Go Lang HTTP Request Example",
		Completed: true,
	}

	// Marshal the Todo object into JSON
	payload, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a payload reader
	payloadReader := bytes.NewReader(payload)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPut, putUrl, payloadReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Response Status:", res.Status)

	// Close the response body
	defer res.Body.Close()
}
func PerformDeleteRequest() {
	// Define the URL for the resource to delete
	deleteUrl := "https://jsonplaceholder.typicode.com/todos/1"
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request with the DELETE method
	req, err := http.NewRequest(http.MethodDelete, deleteUrl, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Perform the HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Response Status:", res.Status)

	// Close the response body
	defer res.Body.Close()
}

func main() {
	// calling the CRUD Functions
	PreformGetRequest()
	PreformPostRequest()
	PerformPutRequest()
	PerformDeleteRequest()
}
