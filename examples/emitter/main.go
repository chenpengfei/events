package main

import (
	"fmt"
	event "github.com/chenpengfei/events/pkg/emitter"
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
