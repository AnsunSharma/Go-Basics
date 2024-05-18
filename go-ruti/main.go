package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(0 * time.Millisecond)
		fmt.Println(s)
	}
}

func say2() {
	fmt.Println("started say2")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("ended say2")
}

func main() {
	go say("world")
	go say("hello")
	go say2()
	time.Sleep(800 * time.Millisecond)

}
