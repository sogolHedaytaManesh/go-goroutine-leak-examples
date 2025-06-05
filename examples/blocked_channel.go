package examples

import (
	"fmt"
	"time"
)

// BlockedChannelLeak demonstrates a goroutine leaking by blocking forever
// waiting to receive from a channel that never sends data or closes.
func BlockedChannelLeak() {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine started: waiting to receive from channel")
		<-ch // blocks forever
		fmt.Println("This line will never be reached")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("BlockedChannelLeak finished (goroutine leaked)")
}

// FixedBlockedChannelLeak shows how to avoid the leak by using select
// with a timeout case to prevent blocking forever.
func FixedBlockedChannelLeak() {
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine started: waiting with timeout")
		select {
		case <-ch:
			fmt.Println("Received from channel")
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout: no data received, avoiding leak")
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("FixedBlockedChannelLeak finished")
}
