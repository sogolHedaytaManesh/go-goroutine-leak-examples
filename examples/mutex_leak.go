package examples

import (
	"fmt"
	"sync"
	"time"
)

// MutexLeak shows a goroutine leak where a mutex is locked but never unlocked,
// causing other goroutines to block indefinitely.
func MutexLeak() {
	var mu sync.Mutex
	mu.Lock() // locked here and never unlocked

	go func() {
		fmt.Println("Goroutine trying to lock mutex (will block forever)")
		mu.Lock()
		fmt.Println("This line never reached")
		mu.Unlock()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("MutexLeak finished (goroutine leaked)")
}

// FixedMutexLeak unlocks the mutex properly so no goroutine blocks.
func FixedMutexLeak() {
	var mu sync.Mutex
	mu.Lock()

	go func() {
		fmt.Println("Goroutine waiting to lock mutex")
		mu.Lock()
		fmt.Println("Goroutine acquired mutex")
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)
	mu.Unlock()
	time.Sleep(1 * time.Second)
	fmt.Println("FixedMutexLeak finished")
}
