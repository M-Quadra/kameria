package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	que := NewQueue("AMD")
	assert.Equal(t, que.Size(), 1)

	que.Push("YES")
	assert.Equal(t, que.Size(), 2)
	assert.Equal(t, que.Front(), "AMD")
	assert.Equal(t, que.Back(), "YES")

	assert.False(t, que.Empty())
	assert.Equal(t, que.Pop(), "AMD")
	assert.Equal(t, que.Size(), 1)

	assert.False(t, que.Empty())
	assert.Equal(t, que.Pop(), "YES")
	assert.Equal(t, que.Size(), 0)

	assert.True(t, que.Empty())
}
