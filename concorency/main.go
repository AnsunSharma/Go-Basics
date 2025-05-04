package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	firstchan := make(chan int)

	wg.Add(1)
	go DoSomeMagic(firstchan, &wg)

	// Send values to the channel in a separate goroutine
	go func() {
		for i := 0; i < 10; i++ {
			firstchan <- 2
		}
		close(firstchan)
	}()

	// Wait for the goroutine to finish
	wg.Wait()

	// Sleep for a second to ensure all output is printed
	time.Sleep(1 * time.Second)

	// Example to show behavior of closed channels
	ch := make(chan int)
	close(ch)
	elem, ok := <-ch
	fmt.Println(elem, ok)
}

func DoSomeMagic(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range ch {
		result := value * 2
		fmt.Println("hi concurrency", result)
	}
}
