package main

import (
	event "events/pkg/emitter"
	"fmt"
	"time"
)

func main() {
	emitter := event.NewEmitter()
	emitter.On("_update", func(data interface{}) {
		go func() {
			fmt.Println(data)
		}()
	})
	emitter.Emit("_update", "hello")

	time.Sleep(time.Second)
}
