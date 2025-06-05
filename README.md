# go-goroutine-leak-examples
Real-world examples of goroutine leaks in Go, with detection strategies and fixes.

## Example Scenarios

This repo contains example programs illustrating common goroutine leaks:

- `examples/blocked_channel.go` — goroutine blocked waiting on a channel forever.
- `examples/context_leak.go` — goroutine ignoring context cancellation.
- `examples/timer_leak.go` — leaking timers due to missing Stop().
- `examples/waitgroup_leak.go` — WaitGroup wait blocked due to missing Done().
- `examples/mutex_leak.go` — goroutine blocked on unreleased mutex.
- `examples/http_leak.go` — goroutine leaking by not closing HTTP response body.

## Running Examples

### Running an individual example

Run a single scenario like this:

```bash
go run examples/blocked_channel.go
```
Run all scenarios with:

```bash
go run main.go
```

## Related Article

This repository accompanies my Medium article on goroutine leaks in Go, which explains the root causes, real-world examples, and detection strategies in detail:

[https://medium.com/@sogol.hedayatmanesh/goroutine-leaks-in-go-root-causes-real-world-examples-and-ironclad-detection-strategies-435c938d66ed](https://medium.com/@sogol.hedayatmanesh/goroutine-leaks-in-go-root-causes-real-world-examples-and-ironclad-detection-strategies-435c938d66ed)