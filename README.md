# events
[![Build Status](https://travis-ci.com/chenpengfei/events.svg)](https://travis-ci.com/chenpengfei/events)
[![Coverage Status](https://coveralls.io/repos/github/chenpengfei/events/badge.svg)](https://coveralls.io/github/chenpengfei/events)

> An event emitter, safe for concurrent use.

## Asynchronous vs. Synchronous
The `Emitter` calls all listeners synchronously in the order in which they were registered.
When appropriate, listener functions can switch to an asynchronous mode of operation using the `go` keyword.
```
emitter := event.NewEmitter()
  emitter.On("_update", func(data interface{}) {
    go func() {
      fmt.Println(data)
    }()
  })
  emitter.Emit("_update", "hello")
```

## Reference
[![Node JS EventEmitter]](https://nodejs.org/api/events.html)