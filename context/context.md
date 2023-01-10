# Context in Go
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context  

## From Go Blog
In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. **When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.**  

At Google, we developed **a context package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.** The package is publicly available as context. This article describes how to use the package and provides a complete working example.

## type of Go Context interface
```
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```