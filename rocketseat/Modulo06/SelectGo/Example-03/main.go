package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i, ch := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i+2) * time.Second)
				ch <- i + 1
			}
		}(i, ch)
	}

	for i := 0; i < 20; i++ {
		// v1 := <-chans[0]
		// fmt.Println("Got a value, on channel 1:", v1)
		// v2 := <-chans[1]
		// fmt.Println("Got a value, on channel 2:", v2)

		select {
		case v := <-chans[0]:
			log.Println("Got a value, on channel 1:", v)
		case v := <-chans[1]:
			log.Println("Got a value, on channel 2:", v)
		}
	}
}
