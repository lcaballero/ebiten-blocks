package main

import (
	"testing"

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

func Test_NewGame(t *testing.T) {
	g := NewGame(NewGameOpts{vals: vals{}})
	assert.NotNil(t, g.pieces)
	assert.NotNil(t, g.board)
	assert.NotNil(t, g.background)
	assert.NotNil(t, g.current)
	assert.NotNil(t, g.next)
}
