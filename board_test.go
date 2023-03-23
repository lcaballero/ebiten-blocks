package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBoard(t *testing.T) {
	bg := NewBackground()
	b := NewBoard(bg.board)

	assert.NotNil(t, b.grid)
	assert.Equal(t, bg.board, b.box)
}
