## Context

Introduced in 1.7, the context package provides us a way to deal with context in our application. Those contexts can help when it comes to the cancellation of a task or defining timeout. It can also be useful to propagate request-scoped values through the context, but for this article, we will focus on the cancellation feature.

### The Context Interface

```go
type Context interface {
	// Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
	Deadline() (deadline time.Time, ok bool)

	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	//
	// Done is provided for use in select statements:
	//
	//  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
	// See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancelation.
	Done() <-chan struct{}

	// If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
	Err() error

	// Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	//
	// Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	//
	// A key identifies a specific value in a Context. Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
	//
	// Packages that define a Context key should provide type-safe accessors
	// for the values stored using that key:
	//
	// 	// Package user defines a User type that's stored in Contexts.
	// 	package user
	//
	// 	import "context"
	//
	// 	// User is the type of value stored in the Contexts.
	// 	type User struct {...}
	//
	// 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
	// 	type key int
	//
	// 	// userKey is the key for user.User values in Contexts. It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
	// 	var userKey key
	//
	// 	// NewContext returns a new Context that carries value u.
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	// 	// FromContext returns the User value stored in ctx, if any.
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key interface{}) interface{}
}
```

### Default contexts

The Go context package builds a context based on an existing one. Two default contexts exist in the package called TODO and Background:

```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return background
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
	return todo
}
```

both of them are empty contexts. This context is a simple context that is never cancelled and does not carry any value
You should use context.Background() in the main() function, init() functions, and tests.
You should use context.TODO() if you're not sure what context to use.

#### Example 1:

file example_context.go

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func contextDemo(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	name := ctx.Value("name")
	for {
		if ok {
			fmt.Println(name, "will expire at:", deadline)
		} else {
			fmt.Println(name, "has no deadline")
		}
		time.Sleep(time.Second)
	}
}

type contextKey string

const (
	contextKey1 = contextKey("contextKey1")
	contextKey2 = contextKey("contextKey2")
	contextKey3 = contextKey("contextKey3")
	// ...
)

func main() {
	timeout := 3 * time.Second
	deadline := time.Now().Add(4 * time.Second)

	// create context with timeout
	timeOutContext, _ := context.WithTimeout(context.Background(), timeout)
	cancelContext, cancelFunc := context.WithCancel(context.Background())
	deadlineContext, _ := context.WithDeadline(cancelContext, deadline)

	go contextDemo(context.WithValue(timeOutContext, contextKey1, "[timeoutContext]"))
	go contextDemo(context.WithValue(cancelContext, contextKey2, "[cancelContext]"))
	go contextDemo(context.WithValue(deadlineContext, contextKey3, "[deadlineContext]"))

	// Wait for the timeout to expire
	<-timeOutContext.Done()

	// This will cancel the deadline context as well as its
	// child - the cancelContext
	fmt.Println("Cancelling the cancel context...")
	cancelFunc()

	<-cancelContext.Done()
	fmt.Println("The cancel context has been cancelled...")

	// Wait for both contexts to be cancelled
	<-deadlineContext.Done()
	fmt.Println("The deadline context has been cancelled...")
}
```

source

- https://golang.org/pkg/context/
- https://blog.golang.org/context
- https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
- http://p.agnihotry.com/post/understanding_the_context_package_in_golang/
- https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/
