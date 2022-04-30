package main

import (
	"testing"
	"time"

	kmath "github.com/M-Quadra/kameria/k-math"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, kmath.Min(0, 2), 0)
	assert.Equal(t, kmath.Min[int64](4, 6), int64(4))
	assert.Equal(t, kmath.Min("AMD", "Intel"), "AMD")
	assert.Equal(t, kmath.Min(float32(1.8), float32(2.0)), float32(1.8))
	assert.Equal(t, kmath.Min(2.2, 2.4), 2.2)

	assert.NotEqual(t, kmath.Min[int64](4, 6), 4)
	assert.NotEqual(t, kmath.Min(float32(1.8), float32(2.0)), 2.0)

	a := time.Now()
	b := a.Add(time.Hour)
	assert.Equal(t, kmath.MinTimes(a, b), a)
}
