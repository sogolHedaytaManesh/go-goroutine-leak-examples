package examples

import (
	"context"
	"fmt"
	"time"
)

// ContextLeak shows a goroutine that ignores context cancellation,
// causing it to run indefinitely and leak.
func ContextLeak(ctx context.Context) {
	go func() {
		fmt.Println("Goroutine started: ignoring context cancellation")
		for {
			// Simulated work
			time.Sleep(200 * time.Millisecond)
			// No ctx.Done() check here — leaks if context is cancelled
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("ContextLeak finished (goroutine leaked)")
}

// FixedContextLeak respects context cancellation and exits when done.
func FixedContextLeak(ctx context.Context) {
	go func() {
		fmt.Println("Goroutine started: listening to context cancellation")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled, goroutine exiting")
				return
			default:
				time.Sleep(200 * time.Millisecond) // simulate work
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("FixedContextLeak finished")
}
