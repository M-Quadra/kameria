package kmath

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(0, 2), 0)
	assert.Equal(t, Min[int64](4, 6), int64(4))
	assert.Equal(t, Min("AMD", "Intel"), "AMD")
	assert.Equal(t, Min(float32(1.8), float32(2.0)), float32(1.8))
	assert.Equal(t, Min(2.2, 2.4), 2.2)

	assert.NotEqual(t, Min[int64](4, 6), 4)
	assert.NotEqual(t, Min(float32(1.8), float32(2.0)), 2.0)

	a := time.Now()
	b := a.Add(time.Hour)
	assert.Equal(t, MinTimes(a, b), a)
}
