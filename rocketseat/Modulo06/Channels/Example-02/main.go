package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const numberOfRoutines = 30
	ch := make(chan int, numberOfRoutines)

	for i := 0; i <= numberOfRoutines; i++ {
		go func(ch chan<- int, i int) {
			ch <- i
			fmt.Println("I am:", i)
		}(ch, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Got value:", <-ch)
	time.Sleep(1 * time.Second)

	for {
		fmt.Println("Go routines running:", runtime.NumGoroutine()-1)
		time.Sleep(250 * time.Millisecond)
	}
}
