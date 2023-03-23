package main

import (
	"testing"

	"github.com/lcaballero/ebiten-01/rand"
	"github.com/stretchr/testify/assert"
)

func Test_Tetro_String(t *testing.T) {
	assert.Equal(t, "I", I.String())
	assert.Equal(t, "O", O.String())
	assert.Equal(t, "T", T.String())
	assert.Equal(t, "S", S.String())
	assert.Equal(t, "Z", Z.String())
	assert.Equal(t, "J", J.String())
	assert.Equal(t, "L", L.String())
	assert.Equal(t, "unknown", Tetro(42).String())
}

func Test_RandTetro(t *testing.T) {
	rnd := rand.NewRnd(33442)
	for i := 0; i < 1000; i++ {
		assert.NotEqual(t, RandTetro(rnd), "unknown")
	}
}

func Test_Rotation(t *testing.T) {
	assert.Equal(t, R1.Inc(J), R2)
	assert.Equal(t, R2.Inc(J), R3)
	assert.Equal(t, R3.Inc(J), R4)
	assert.Equal(t, R4.Inc(J), R1)
}
