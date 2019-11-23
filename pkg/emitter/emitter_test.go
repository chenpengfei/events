package event

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestEmit(t *testing.T) {
	assert := assert.New(t)

	expected := struct {
		NameA  string
		ValueA string
		NameB  string
		ValueB string
	}{
		NameA:  "nameA",
		ValueA: "valueA",
		NameB:  "nameB",
		ValueB: "valueB",
	}

	t.Run("one Emit, one On", func(t *testing.T) {
		event := NewEmitter()
		event.On(expected.NameA, func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		})
		event.Emit(expected.NameA, expected.ValueA)
	})

	t.Run("one Emit, two On", func(t *testing.T) {
		event := NewEmitter()
		event.On(expected.NameA, func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		})
		event.On(expected.NameA, func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		})
		event.Emit(expected.NameA, expected.ValueA)
	})

	t.Run("two Emit, two On", func(t *testing.T) {
		event := NewEmitter()
		event.On(expected.NameA, func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		})
		event.On(expected.NameB, func(data interface{}) {
			assert.Equal(expected.ValueB, data.(string))
		})
		event.Emit(expected.NameA, expected.ValueA)
		event.Emit(expected.NameB, expected.ValueB)
	})

	t.Run("RemoveListener", func(t *testing.T) {
		counter := 0
		cb := func(data interface{}) {
			counter++
			assert.Equal(expected.ValueA, data.(string))
		}
		event := NewEmitter()
		event.On(expected.NameA, cb)
		event.Emit(expected.NameA, expected.ValueA)
		assert.Equal(1, counter)
		event.RemoveListener(expected.NameA, cb)
		event.Emit(expected.NameA, expected.ValueA)
		assert.Equal(1, counter)
	})

	t.Run("RemoveListener, On again", func(t *testing.T) {
		counter := 0
		cb := func(data interface{}) {
			counter++
			assert.Equal(expected.ValueA, data.(string))
		}
		event := NewEmitter()
		event.On(expected.NameA, cb)
		event.Emit(expected.NameA, expected.ValueA)
		assert.Equal(1, counter)
		event.RemoveListener(expected.NameA, cb)
		event.On(expected.NameA, cb)
		event.Emit(expected.NameA, expected.ValueA)
		assert.Equal(2, counter)
	})

	t.Run("ListenerCount", func(t *testing.T) {
		cbA := func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		}
		event := NewEmitter()
		event.On(expected.NameA, cbA)
		assert.Equal(1, event.ListenerCount(expected.NameA))
		event.RemoveListener(expected.NameA, cbA)
		assert.Equal(0, event.ListenerCount(expected.NameA))
	})

	t.Run("data race", func(t *testing.T) {
		cbA := func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		}
		cbB := func(data interface{}) {
			assert.Equal(expected.ValueB, data.(string))
		}
		n := 1000
		event := NewEmitter()
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			for i := 0; i < n; i++ {
				event.Emit(expected.NameA, expected.ValueA)
				event.Emit(expected.NameB, expected.ValueB)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < n; i++ {
				event.On(expected.NameA, cbA)
				event.On(expected.NameB, cbB)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < n; i++ {
				event.RemoveListener(expected.NameA, cbA)
				event.RemoveListener(expected.NameB, cbB)
			}
			wg.Done()
		}()

		wg.Wait()
	})
}
