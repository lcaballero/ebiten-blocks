package shapes

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rectAtCenter(vs ...Vec) Rects {
	rects := Rects{}
	for _, v := range vs {
		pos := v.Add(Vec{-5, -5})
		dim := Vec{10, 10}
		rects = append(rects, Rect{Pos: pos, Dim: dim})
	}
	return rects
}

func Test_(t *testing.T) {
	cases := []struct {
		name  string
		rects Rects
		exp   Rects
	}{
		{
			name:  "empty list of vecs works fine",
			rects: rectAtCenter(),
			exp:   Rects{},
		},
		{
			name:  "sort single vec",
			rects: rectAtCenter(Vec{3, 3}),
			exp:   rectAtCenter(Vec{3, 3}),
		},
		{
			name:  "sort two vecs",
			rects: rectAtCenter(Vec{3, 3}, Vec{23, 23}),
			exp:   rectAtCenter(Vec{3, 3}, Vec{23, 23}),
		},
		{
			name: "sort three vecs on y",
			rects: rectAtCenter(
				Vec{13, 3}, Vec{3, 3}, Vec{23, 3}),
			exp: rectAtCenter(
				Vec{3, 3}, Vec{13, 3}, Vec{23, 3}),
		},
		{
			name: "sort vecs on 2 rows",
			rects: rectAtCenter(
				Vec{13, 3}, Vec{3, 3}, Vec{33, 3}, Vec{23, 3},
				Vec{13, 13}, Vec{23, 13}, Vec{33, 13}, Vec{3, 13},
			),
			exp: rectAtCenter(
				Vec{3, 3}, Vec{13, 3}, Vec{23, 3}, Vec{33, 3},
				Vec{3, 13}, Vec{13, 13}, Vec{23, 13}, Vec{33, 13},
			),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(SortByCenter(c.rects))
			assert.Equal(t, len(c.exp), len(c.rects))
			for i, expected := range c.exp {
				actual := c.rects[i]
				a := expected.Center()
				b := actual.Center()
				assert.Equal(t, a.X(), b.X())
				assert.Equal(t, a.Y(), b.Y())
			}
		})
	}
}
