package main

import "fmt"

func main() {
	// ch := make(chan int, 1)
	ch := make(chan int)

	ch <- 0
	n, ok := <-ch
	fmt.Println("Got value:", n, "| Is this channel closed?", !ok)

	close(ch)

	n, ok = <-ch
	fmt.Println("Got value:", n, "| Is this channel closed?", !ok)
}
