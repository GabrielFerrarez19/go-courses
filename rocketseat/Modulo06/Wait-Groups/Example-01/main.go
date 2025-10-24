package main

import (
	"fmt"
	"time"
)

func main() {
	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Println("Terminou!!!")
	// }()

	done := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Terminou!!!")
		done <- 0
	}()

	<-done
}
