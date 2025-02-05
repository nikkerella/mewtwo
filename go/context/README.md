# Context

## Definition

``` go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```

- Deadline():
  - Returns the time when work should be canceled.
  - Returns false when no deadline is set.
- Done():
  - Returns a channel when work done (cancelled). It can be nil if the context can never be canceled.
- Err()
  - Returns non-nil error if Done is closed.
- Value()
  - Returns the value of the key with the context, if no value is associated with the key, it will return nil.

### emptyCtx

``` go
type emptyCtx struct{}

func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (emptyCtx) Done() <-chan struct{} {
	return nil
}

func (emptyCtx) Err() error {
	return nil
}

func (emptyCtx) Value(key any) any {
	return nil
}
```

### backgroundCtx

backgroundCtx = emptyCtx + String()

``` go
type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

func Background() Context {
  return backgroundCtx{}
}
```

### todoCtx

todoCtx = emptyCtx + String()

``` go
type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}

func TODO() Context {
  return todoCtx{}
}
```

### cancelCtx

``` go
// A cancelCtx can be canceled. When canceled, it also cancels any children that implement canceler.
type cancelCtx struct {
	Context

	mu       sync.Mutex
	done     atomic.Value
	children map[canceler]struct{}
	err      error
	cause    error
}
```

- Context: Embedded context, the cancelCtx struct embeds the Context interface. This means it inherits all the methods of the Context interface, such as Deadline(), Done(), Err(), and Value().
- mu: protects the fields
- done: ???
- children:
  - This is a map that keeps track of child contexts that implement the canceler interface. When the parent context is canceled, it cancels all its children.
  - The map is set to nil after the first cancellation to release references and allow garbage collection.
- err: stores the error that caused the context to be cancelled.
- cause: stores  custom errors or additional information that explains the cancellation? (Unlike err is standard.)


``` go
func (c *cancelCtx) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}
```



### canceler

``` go
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}


```

## Features

### Control Flow

Context is mainly used to control the lifecycle of concurrent operations. It allows you to propagate cancellation signals, timeouts, and deadlines across multiple Goroutines, enabling you to manage their execution.

- Cancellation: A context created with context.WithCancel can propagate cancellation signals. When the cancel function is called, all Goroutines listening to this context will receive the cancellation signal and stop execution.
- Timeout and Deadline: A context created with context.WithTimeout or context.WithDeadline will automatically cancel the operation after a specified duration or deadline. This is particularly useful for limiting the maximum execution time of an operation.

### Value Propagation

Context can also safely propagate request-scoped values across multiple Goroutines. Using context.WithValue to store key-value pairs in a context, which can then be accessed throughout the request chain.

Request-Scoped Values: For example, store authentication information in the context, and these values can be accessed and used at different stages of request processing.

### Concurrency Safety

Context is concurrency-safe, meaning it can be read or passed across multiple Goroutines without requiring additional synchronization mechanisms. This is because context is immutable—each time a new context is created, it returns a new instance without affecting the original one.