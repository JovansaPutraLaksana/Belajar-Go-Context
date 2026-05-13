package Belajar_Go_Context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
