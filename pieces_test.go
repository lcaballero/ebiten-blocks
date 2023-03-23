package main

import (
	"testing"

	"github.com/lcaballero/ebiten-01/rand"
	"github.com/stretchr/testify/assert"
)

func Test_Peices_New(t *testing.T) {
	pieces := NewPieces()
	assert.NotNil(t, pieces)
	assert.Equal(t, 7, pieces.Len())
	rnd := rand.NewRnd(92219)
	for i := 0; i < pieces.Len()*2; i++ {
		assert.NotNil(t, pieces.blocks.At(i))
		assert.NotNil(t, pieces.blocks.Pick(rnd))
	}
}
