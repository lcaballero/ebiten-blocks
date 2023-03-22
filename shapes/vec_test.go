package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_Dir(t *testing.T) {
	cases := []struct {
		name string
		line Line
		vec  Vec
		dir  Vec
	}{
		{
			name: "pos diag",
			line: Line{A: Vec{20, 20}, B: Vec{10, 10}},
			vec:  Vec{-10, -10},
			dir:  Vec{1, 1}.Normalize().Negate(),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			vec := c.line.ToVec()
			assert.Equal(t, c.vec, vec)
			dir := vec.Normalize()
			assert.Equal(t, c.dir, dir)
			//assert.True(t, false)
			t.Logf("vec: %v, dir: %v", vec, dir)
		})
	}
}

func TestScale_Invert(t *testing.T) {
	a := Scale{2.0, 2.0}
	b := a.Invert()
	assert.Equal(t, .5, b[0])
	assert.Equal(t, .5, b[1])

	c := b.Invert()
	assert.Equal(t, a[0], c[0])
	assert.Equal(t, a[1], c[1])
}

func TestVec_Octant(t *testing.T) {
	cases := []struct {
		name   string
		dir    Vec
		octant Octant
	}{
		// no direction
		{name: "zero vec", dir: Zero, octant: NoDirection},

		// diags
		{name: "ne", dir: Vec{1, 1}, octant: NE},
		{name: "se", dir: Vec{1, -1}, octant: SE},
		{name: "sw", dir: Vec{-1, -1}, octant: SW},
		{name: "nw", dir: Vec{-1, 1}, octant: NW},

		// axis
		{name: "east", dir: Vec{1, 0}, octant: East},
		{name: "west", dir: Vec{-1, 0}, octant: West},
		{name: "north", dir: Vec{0, 1}, octant: North},
		{name: "south", dir: Vec{0, -1}, octant: South},

		// north octants
		{name: "ene", dir: Vec{2, 1}, octant: ENE},
		{name: "nne", dir: Vec{1, 2}, octant: NNE},
		{name: "nnw", dir: Vec{-1, 2}, octant: NNW},
		{name: "wnw", dir: Vec{-2, 1}, octant: WNW},

		// south octants
		{name: "wsw", dir: Vec{-2, -1}, octant: WSW},
		{name: "ssw", dir: Vec{-1, -2}, octant: SSW},
		{name: "sse", dir: Vec{1, -2}, octant: SSE},
		{name: "ese", dir: Vec{2, -1}, octant: ESE},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.octant, c.dir.Octant())
		})
	}
}
