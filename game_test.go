package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type vals map[string]interface{}

func (v vals) IsSet(string) bool {
	return false
}
func (v vals) Bool(string) bool {
	return false
}
func (v vals) Int(string) int {
	return 1
}
func (v vals) Int64(string) int64 {
	return 2
}
func (v vals) String(string) int64 {
	return ""
}

func Test_NewGame(t *testing.T) {
	g := NewGame(NewGameOpts{vals: vals{}})
	assert.NotNil(t, g.pieces)
	assert.NotNil(t, g.board)
	assert.NotNil(t, g.background)
	assert.NotNil(t, g.current)
	assert.NotNil(t, g.next)
	assert.NotNil(t, g.rnd)
	assert.NotNil(t, g.keys)
	assert.False(t, g.paused)

	w, h := g.Layout(1, 1)
	assert.Equal(t, 320, w)
	assert.Equal(t, 240, h)

	assert.Equal(t, 0, g.frames)
	assert.Equal(t, time.Duration(0), g.elapsed)
	assert.Equal(t, time.Duration(0), g.accum)
	assert.Equal(t, time.Duration(0), g.seconds)
}
