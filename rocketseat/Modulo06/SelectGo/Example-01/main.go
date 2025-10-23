package main

import (
	"fmt"
	"time"
)

func read(ch <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("Got something:", v)
			return
		default:
			fmt.Println("Nothing gere yet.")
		}
	}
}

func main() {
	ch := make(chan int)

	go read(ch)

	time.Sleep(5 * time.Second)

	ch <- 10
}
