package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueDequeue(t *testing.T) {
	q := Queue[int]{3, 2, 1}

	for i := 3; i >= 0; i-- {
		v, hasVal := q.Dequeue()

		assert.Equal(t, i, v)
		if i == 0 {
			assert.False(t, hasVal)
		} else {
			assert.True(t, hasVal)
		}
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := Queue[string]{}

	q.Enqueue("Hello", "World")

	assert.Equal(t, 2, q.Length())
	v, _ := q.Dequeue()
	assert.Equal(t, "Hello", v)
	v, _ = q.Dequeue()
	assert.Equal(t, "World", v)
}
