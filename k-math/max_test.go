package kmath

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 3), 3)
	assert.Equal(t, Max[int64](5, 7), int64(7))
	assert.Equal(t, Max("M$", "MS"), "MS")
	assert.Equal(t, Max(float32(1.1), float32(1.2)), float32(1.2))
	assert.Equal(t, Max(1.1, 1.2), 1.2)

	assert.NotEqual(t, Max[int64](5, 7), 7)
	assert.NotEqual(t, Max(float32(1.1), float32(1.2)), 1.2)

	a := time.Now()
	b := a.Add(time.Hour)
	assert.Equal(t, MaxTimes(a, b), b)
}
