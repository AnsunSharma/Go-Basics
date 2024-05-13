package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error is ", err)
		return
	}
	defer res.Body.Close()

	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		fmt.Printf("Could not read \n", er)
		return
	} else {
		fmt.Println(string(data))
	}

	// for index, value := range arr {
	// 	count++
	// 	fmt.Println(index, value)
	// }
}
