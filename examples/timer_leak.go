package examples

import (
	"fmt"
	"time"
)

// TimerLeak demonstrates leaking timers by not stopping them,
// which can cause internal goroutines to leak.
func TimerLeak() {
	_ = time.NewTimer(5 * time.Second)
	fmt.Println("Timer started but never stopped")
	// No timer.Stop(), timer will fire and stay in memory longer than needed
	time.Sleep(2 * time.Second)
	fmt.Println("TimerLeak finished (timer goroutine leaked)")
}

// FixedTimerLeak stops and drains the timer to avoid leaks.
func FixedTimerLeak() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer started and will be stopped early")
	if !timer.Stop() {
		<-timer.C // drain the channel if needed
	}
	fmt.Println("Timer stopped and drained")
	time.Sleep(1 * time.Second)
	fmt.Println("FixedTimerLeak finished")
}
