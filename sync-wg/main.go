package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d stated ", i)

	fmt.Printf("worker %d end ", i)
}

func main() {

	var wg sync.WaitGroup
	fmt.Println(" learning Sync wait group")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("working task complited")
}
