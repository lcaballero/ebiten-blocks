package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Mark_In(t *testing.T) {
	cases := []struct {
		name     string
		expected bool
		mark     *mark
		setup    func() grid
		call     func(*mark) *mark
	}{
		{
			name:     "empty mark not in empty grid",
			expected: false,
			mark:     &mark{},
			setup: func() grid {
				return grid{}
			},
			call: func(m *mark) *mark {
				return m
			},
		},
		{
			name:     "mark in grid",
			expected: true,
			mark: &mark{
				rc: [2]int{11, 18},
			},
			setup: func() grid {
				g := grid{}
				g[[2]int{11, 17}] = &mark{}
				return g
			},
			call: func(m *mark) *mark {
				return m.up()
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grid := c.setup()
			m := c.call(c.mark)
			assert.Equal(t, c.expected, m.in(grid))
		})
	}
}
