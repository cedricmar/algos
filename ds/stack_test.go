package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPop(t *testing.T) {
	s := Stack[int]{1, 2, 3}

	for i := 3; i >= 0; i-- {
		v, hasVal := s.Pop()

		assert.Equal(t, i, v)
		if i == 0 {
			assert.False(t, hasVal)
		} else {
			assert.True(t, hasVal)
		}
	}
}

func TestStackPush(t *testing.T) {
	s := Stack[string]{}

	s.Push("Hello", "World")

	assert.Equal(t, 2, s.Length())
	v, _ := s.Pop()
	assert.Equal(t, "World", v)
	v, _ = s.Pop()
	assert.Equal(t, "Hello", v)
}
