package main

import (
	"math/rand"
	"testing"

	"github.com/M-Quadra/kameria"
	"github.com/stretchr/testify/assert"
)

func TestSliceToInterfaces(t *testing.T) {
	aryA := []int{rand.Int(), rand.Int(), rand.Int()}
	opt := kameria.SliceToInterfaces(aryA)
	assert.ElementsMatch(t, aryA, opt)
	opt = kameria.SliceToInterfaces(&aryA)
	assert.ElementsMatch(t, aryA, opt)

	aryB := make([]*int, 0, len(aryA))
	for i := range aryA {
		aryB = append(aryB, &aryA[i])
	}
	opt = kameria.SliceToInterfaces(aryB)
	assert.ElementsMatch(t, aryB, opt)
	opt = kameria.SliceToInterfaces(&aryB)
	assert.ElementsMatch(t, aryB, opt)

	assert.Nil(t, kameria.SliceToInterfaces(nil))
	assert.Nil(t, kameria.SliceToInterfaces(1))
}
