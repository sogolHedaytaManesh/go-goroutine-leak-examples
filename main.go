package main

import (
	"context"
	"fmt"
	"github.com/sogolHedaytaManesh/go-goroutine-leak-examples/examples"
	"time"
)

func main() {
	fmt.Println("Running blocked channel leak example:")
	examples.BlockedChannelLeak()

	fmt.Println("\nRunning fixed version:")
	examples.FixedBlockedChannelLeak()

	fmt.Println("\nRunning context leak example:")
	ctx, cancel := context.WithCancel(context.Background())
	examples.ContextLeak(ctx)
	cancel()
	time.Sleep(1 * time.Second) // give goroutine time to run

	fmt.Println("\nRunning fixed context leak example:")
	ctx2, cancel2 := context.WithCancel(context.Background())
	examples.FixedContextLeak(ctx2)
	cancel2()
	time.Sleep(1 * time.Second)

	fmt.Println("\nRunning timer leak example:")
	examples.TimerLeak()

	fmt.Println("\nRunning fixed timer leak example:")
	examples.FixedTimerLeak()

	fmt.Println("\nRunning waitgroup leak example:")
	go examples.WaitGroupLeak()
	time.Sleep(2 * time.Second)

	fmt.Println("\nRunning fixed waitgroup leak example:")
	examples.FixedWaitGroupLeak()

	fmt.Println("\nRunning mutex leak example:")
	examples.MutexLeak()

	fmt.Println("\nRunning fixed mutex leak example:")
	examples.FixedMutexLeak()

	fmt.Println("\nRunning HTTP leak example:")
	examples.HTTPLLeak("https://www.google.com")

	fmt.Println("\nRunning fixed HTTP leak example:")
	examples.FixedHTTPLLeak("https://www.google.com")
}
