package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("fileZip.txt")
	if err != nil {
		fmt.Println("file is not created", err)
	}

	fileOne, error := os.Create("one.txt")
	if error != nil {
		fmt.Println("file is not created", error)
	}
	defer file.Close()

	_, err = file.WriteString("Hello World! This is a test for zip files.")
	if err != nil {
		fmt.Printf("Error writing to the file: %v\n", err)
	}

	byte, errorone := io.WriteString(fileOne, "file is first file in zip.\n")
	if errorone != nil {
		fmt.Printf("Error appending to the file: %v\n", errorone)
		return
	}
	fmt.Println(byte)

	// Close the second file before reading from it
	defer fileOne.Close()

	content, err := os.ReadFile("one.txt")
	if err != nil {
		fmt.Printf("Error reading from the file: %v\n", err)
		return
	}
	fmt.Printf("%s", content)

}
