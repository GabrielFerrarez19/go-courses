package main

import (
	"fmt"
	"time"
)

func runner(ch chan<- int, id int) {
	fmt.Println("iniciado Go routine runner:", id, "\n")
	time.Sleep(5 * time.Second)
	for {
		fmt.Printf("[RUNNER %d] Estou tentando enviar\n", id)
		ch <- id
		fmt.Printf("[RUNNER %d] Consegui enviar\n", id)
	}
}

func main() {
	ch := make(chan int, 4)
	go runner(ch, 1)
	go runner(ch, 2)
	go runner(ch, 3)
	go runner(ch, 4)

	for {
		num := <-ch
		fmt.Printf("[MAIN] Recebi um novo valor %d\n", num)
		time.Sleep(3 * time.Second)
	}
}
