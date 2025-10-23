package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	ticker := time.NewTicker(250 * time.Millisecond).C
	for {
		select {
		case <-ticker:
			fmt.Println(name, "Tick...")

		case <-ctx.Done():
			fmt.Println("Finishing due to: ", ctx.Err().Error())
			fmt.Println(name, "Recieved a signal to finish... exiting...")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go worker(ctx, "Worker1")
	go worker(ctx, "Worker2")

	time.Sleep(12 * time.Second)
	fmt.Println("main exiting")
}
