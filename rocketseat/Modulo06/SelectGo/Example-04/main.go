package main

import (
	"fmt"
	"runtime"
	"time"
)

func takesTooLong(ch chan<- int) {
	time.Sleep(10 * time.Second)
	ch <- 20
}

func takesNotSoLong(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 999
}

func main() {
	stop := time.After(5 * time.Second)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go takesTooLong(ch1)
	go takesNotSoLong(ch2)
	defer fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	for {
		select {
		case <-ch1:
			fmt.Println("Too long finished")
		case <-ch2:
			fmt.Println("not so long finished")
		case <-stop:
			fmt.Println("This job is taking too long finish... aborting")
			return
		}
	}
}
