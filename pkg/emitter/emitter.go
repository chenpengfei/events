// event lib, safe for concurrent
package event

import (
	"fmt"
	"sync"
)

//todo.需要改成异步？

type OnCallback func(data interface{})

var rw sync.RWMutex

type Emitter struct {
	store map[string][]OnCallback
}

func NewEmitter() *Emitter {
	return &Emitter{
		store: make(map[string][]OnCallback),
	}
}

func (e *Emitter) Emit(name string, data interface{}) {
	rw.RLock()
	defer rw.RUnlock()

	if _, ok := e.store[name]; ok {
		for _, cb := range e.store[name] {
			if cb != nil {
				cb(data)
			}
		}
	}
}

func (e *Emitter) On(name string, cb OnCallback) {
	rw.Lock()
	defer rw.Unlock()

	if _, ok := e.store[name]; !ok {
		e.store[name] = make([]OnCallback, 0)
	}

	for i := 0; i < len(e.store[name]); i++ {
		if e.store[name][i] == nil {
			e.store[name][i] = cb
			return
		}
	}
	e.store[name] = append(e.store[name], cb)
}

func (e *Emitter) RemoveListener(name string, cb OnCallback) {
	rw.Lock()
	defer rw.Unlock()

	if cbs, ok := e.store[name]; ok {
		for i := 0; i < len(cbs); i++ {
			if fmt.Sprintf("%v", cb) == fmt.Sprintf("%v", cbs[i]) {
				cbs[i] = nil
			}
		}
	}
}
