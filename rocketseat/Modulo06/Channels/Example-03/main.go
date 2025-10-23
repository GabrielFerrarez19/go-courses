package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	C chan struct{}
}

func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{make(chan struct{}, maxConcurrent)}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	select {
	case <-s.C:

	default:
		fmt.Println("Nada para liberar")
	}
}

func worker(id int, sema *Semaphore, work func()) {
	sema.Acquire()
	go func() {
		defer sema.Release()
		work()
	}()
}

func main() {
	sema := NewSemaphore(3)

	for i := 0; i < 30; i++ {
		id := i

		worker(id, sema, func() {
			fmt.Println("Go routine iniciada", id)
			time.Sleep(2 * time.Second)
			fmt.Println("Go routine stop", id)
		})
	}

	time.Sleep(60 * time.Second)
}
