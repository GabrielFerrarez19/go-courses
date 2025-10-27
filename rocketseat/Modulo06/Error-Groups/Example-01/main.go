package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func TaskWithContext(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if id != 2 {
			time.Sleep(time.Millisecond * 200)

			fmt.Printf("Task %d completed successfully\n", id)
			return nil
		}
		return errors.New("task 2 failed")
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := 1; i <= 5; i++ {
		i := i

		g.Go(func() error {
			return TaskWithContext(ctx, i)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %s \n", err.Error())
	}

	fmt.Printf("Everything executed with success")
}
