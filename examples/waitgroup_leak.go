package examples

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroupLeak demonstrates a goroutine leak where WaitGroup.Wait()
// blocks forever because Done() is never called.
func WaitGroupLeak() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Goroutine started but forgot to call Done()")
		time.Sleep(1 * time.Second)
		// Missing wg.Done() here causes Wait() to block forever
	}()

	fmt.Println("Waiting on WaitGroup (will block forever)")
	wg.Wait()
	fmt.Println("WaitGroupLeak finished (this line never reached)")
}

// FixedWaitGroupLeak calls Done() properly to avoid blocking.
func FixedWaitGroupLeak() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Goroutine started and will call Done()")
		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	fmt.Println("Waiting on WaitGroup")
	wg.Wait()
	fmt.Println("FixedWaitGroupLeak finished")
}
