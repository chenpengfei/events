package event

import (
	"github.com/stretchr/testify/assert"
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

	t.Run("one listener", func(t *testing.T) {
		event := NewEmitter()
		event.On(expected.NameA, func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		})
		event.Emit(expected.NameA, expected.ValueA)
	})

	t.Run("two listener", func(t *testing.T) {
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

	t.Run("listen once", func(t *testing.T) {
		event := NewEmitter()
		times := 0
		event.Once(expected.NameA, func(data interface{}) {
			times++
			assert.Equal(expected.ValueA, data.(string))
		})
		event.Emit(expected.NameA, expected.ValueA)
		event.Emit(expected.NameA, expected.ValueA)
		assert.Equal(1, times)
		assert.Equal(0, event.ListenerCount(expected.NameA))
	})

	t.Run("two listener's callback is the same", func(t *testing.T) {
		event := NewEmitter()
		times := 0
		callback := func(data interface{}) {
			times++
			assert.Equal(expected.ValueA, data.(string))
		}
		event.On(expected.NameA, callback)
		event.On(expected.NameA, callback)
		event.Emit(expected.NameA, expected.ValueA)

		assert.Equal(2, times)
	})

	t.Run("remove listener", func(t *testing.T) {
		cbA := func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		}
		cbB := func(data interface{}) {
			assert.Equal(expected.ValueB, data.(string))
		}

		event := NewEmitter()
		event.On(expected.NameA, cbA)
		event.On(expected.NameA, func(data interface{}) {
		})
		event.On(expected.NameB, cbB)
		event.RemoveListener(expected.NameA, cbA)

		assert.Equal(1, event.ListenerCount(expected.NameA))
		assert.Equal(1, event.ListenerCount(expected.NameB))
	})

	t.Run("listener count", func(t *testing.T) {
		cbA := func(data interface{}) {
			assert.Equal(expected.ValueA, data.(string))
		}
		event := NewEmitter()
		event.On(expected.NameA, cbA)
		assert.Equal(1, event.ListenerCount(expected.NameA))
		event.On(expected.NameA, cbA)
		assert.Equal(2, event.ListenerCount(expected.NameA))
		event.RemoveListener(expected.NameA, cbA)
		assert.Equal(1, event.ListenerCount(expected.NameA))
		assert.Equal(0, event.ListenerCount(expected.NameB))
	})
}
