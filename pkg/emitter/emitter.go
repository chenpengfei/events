package event

import (
	"fmt"
)

type Callback func(data interface{})

type listener struct {
	callback Callback
	once     bool
}

type Emitter struct {
	store map[string][]listener
}

func NewEmitter() *Emitter {
	return &Emitter{
		store: make(map[string][]listener),
	}
}

func (e *Emitter) Emit(name string, data interface{}) {
	if _, ok := e.store[name]; ok {
		for _, ln := range e.store[name] {
			if ln.once {
				e.RemoveListener(name, ln.callback)
			}
			ln.callback(data)
		}
	}
}

func (e *Emitter) On(name string, cb Callback) {
	e.register(name, cb, false)
}

// listener is removed and then invoked
func (e *Emitter) Once(name string, cb Callback) {
	e.register(name, cb, true)
}

func (e *Emitter) register(name string, cb Callback, once bool) {
	if _, ok := e.store[name]; !ok {
		e.store[name] = make([]listener, 0)
	}

	e.store[name] = append(e.store[name], listener{
		callback: cb,
		once:     once,
	})
}

func (e *Emitter) RemoveListener(name string, cb Callback) {
	if cbs, ok := e.store[name]; ok {
		i := 0
		found := false
		for _, v := range cbs {
			if e.equal(cb, v.callback) && !found {
				found = true
				continue
			}

			cbs[i] = v
			i++
		}
		e.store[name] = cbs[:i]
	}

}

func (e *Emitter) equal(a Callback, b Callback) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func (e *Emitter) ListenerCount(name string) int {
	return len(e.store[name])
}
