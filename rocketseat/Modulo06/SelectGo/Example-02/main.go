package main

import (
	"time"
)

func sendOrDrop(send chan<- []byte, log []byte) {
	select {
	case send <- log:
		// set with success... do nothing
	default:

	}
}

func main() {
	ch := make(chan int)

	time.Sleep(5 * time.Second)

	ch <- 10
}
