package examples

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPLLeak demonstrates a goroutine leak caused by not closing
// the HTTP response body, which blocks the goroutine.
func HTTPLLeak(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	go func() {
		fmt.Println("Goroutine started: reading HTTP response without closing body")
		// Missing resp.Body.Close() causes leak
		_, _ = io.ReadAll(resp.Body)
		fmt.Println("This line may never be reached")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("HTTPLLeak finished (goroutine leaked)")
}

// FixedHTTPLLeak closes the response body properly to avoid leak.
func FixedHTTPLLeak(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	go func() {
		fmt.Println("Goroutine started: reading HTTP response with body close")
		defer resp.Body.Close()
		_, _ = io.ReadAll(resp.Body)
		fmt.Println("Goroutine finished reading and closed body")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("FixedHTTPLLeak finished")
}
